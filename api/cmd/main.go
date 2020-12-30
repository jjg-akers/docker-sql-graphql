package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/graphql-go/handler"
	"github.com/jjg-akers/docker-sql-graphql/cmd/schema"
	"github.com/jjg-akers/docker-sql-graphql/cmd/subscriptions"
)

func search(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		// must be get
		w.WriteHeader(http.StatusNotAcceptable)
	}

	// get query params
	if err := r.ParseForm(); err != nil {
		log.Println("could not parse form")
		w.WriteHeader(http.StatusBadRequest)
	}

	//check type
	var tp string

	if t, ok := r.Form["type"]; ok {
		switch t[0] {
		case "date":
			tp = "saerch by date"

		case "title":
			tp = "search by title"
		case "url":
			tp = "search by url"
		default:
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}

	//searchtype := r.Form.Get("type")

	// err := DB.Ping()
	// if err != nil {
	// 	fmt.Fprintf(w, "Error pining")

	// 	//fmt.Println("could not ping db")
	// 	return
	// }
	fmt.Fprintf(w, "result: %s", tp)
	// fmt.Fprintf(w, "search type: %s, %q", html.EscapeString(r.URL.Path))
}

func hi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi")
}

// func h(w http.ResponseWriter, r *http.Request) {

// }

// type Holder struct {
// 	Items []It `json:"items"`
// }
// type It struct {
// 	Title string `json:"title"`
// 	URI   string `json:"uri"`
// 	Date  string `json:"date"`
// }

var DB *sql.DB

func main() {

	// read in input json
	var t Thing

	input, err := ioutil.ReadFile("../input/input.json")
	if err != nil {
		log.Fatalln("error readFile: ", err)
	}
	//var data interface{}
	err = json.Unmarshal(input, &t)
	if err != nil {
		fmt.Println("error marshalling: ", err)
	}

	//fmt.Println("length of thing: ", len(t.Items))

	// time.Sleep(time.Second * 10)

	DB, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/exercise")
	if err != nil {
		fmt.Println("error opening db: ", err)
	}

	err = DB.Ping()
	if err != nil {
		fmt.Println("could not ping db, sleeping for 10s")
		time.Sleep(time.Second * 10)

	}

	err = DB.Ping()
	if err != nil {
		log.Panicln("could not ping db, sleeping for 10s")
		// time.Sleep(time.Second *10)
	}

	// insert all the data
	//query := "INSERT INT0 (title, uri, date) VALUES"

	sb := strings.Builder{}

	params := []interface{}{}

	sb.WriteString("INSERT INTO data (title, uri, date_added) VALUES")

	for i, record := range t.Items {
		params = append(params, record.Title, record.Uri.String(), record.Date.String())
		// fmt.Fprintf(&sb2, " (%s, %s, %s)", record.Title, record.Uri.String(), record.Date.String())
		sb.WriteString(" (?,?,?)")
		if i != len(t.Items)-1 {
			sb.WriteString(",")
		} else {
			sb.WriteString(";")
		}
	}

	ctx := context.Background()

	//fmt.Println("query: ", sb.String())

	//queryStr := strings.TrimSuffix(sb.String(), ","
	query, err := DB.PrepareContext(ctx, sb.String())
	if err != nil {
		log.Fatalln("error preparing: ", err)
	}

	res, err := query.ExecContext(ctx, params...)
	if err != nil {
		log.Fatalln("error exec: ", err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		fmt.Println("error rows: ", err)

	}

	fmt.Println("rows affected: ", rows)

	// fmt.Println("starting server")

	// http.HandleFunc("/", search)

	// http.HandleFunc("/hi", hi)

	// log.Fatal(http.ListenAndServe(":8080", nil))

	// UnmarshalJSON()

	h := handler.New(&handler.Config{
		Schema:     &schema.Schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})

	//graphql api server
	http.Handle("/", h)

	// set up handler for new subscriptions
	http.HandleFunc("/subscriptions", subscriptions.Handler)

	// serve Graphiql in-brower editor
	// fs := http.FileServer(http.Dir("../static"))
	// http.Handle("/", fs)

	fmt.Println("starting server of 8080")

	http.ListenAndServe(":8080", nil)

}

type Thing struct {
	Items []Item `json:"items"`
}

type Item struct {
	Title string    `json:"title,string"`
	Uri   url.URL   `json:"uri,string"`
	Date  time.Time `json:"date,string"`
}

func (it *Item) UnmarshalJSON(j []byte) error {

	var timeFormats = []string{"2006-01-02T15", "January 02", "2006-01-02T15:04:05", "2006-01-02", "January 2 2006", "02 January 2006", "02 Jan 2006", "2006-01-02T15:04:05Z07:00"}

	//timeFormats := []
	var rawStr map[string]string

	if err := json.Unmarshal(j, &rawStr); err != nil {
		return err
	}

	for k, v := range rawStr {
		switch strings.ToLower(k) {
		case "title":
			it.Title = v
		case "url":
			u, err := url.Parse(v)
			if err != nil {
				return err
			}
			it.Uri = *u
		case "date":
			t, err := parseTime(v, timeFormats)
			if err != nil {
				return err
			}

			it.Date = t
		}
	}
	return nil

}

func parseTime(input string, formats []string) (time.Time, error) {
	for _, format := range formats {
		t, err := time.Parse(format, input)
		if err == nil {
			return t, nil
		}
	}
	log.Println("error parsing time: ", input)
	return time.Time{}, errors.New("Unrecognized time format")
}
