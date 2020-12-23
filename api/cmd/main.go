package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/handler"
	"github.com/jjg-akers/go-docker-sql/schema"

	_ "github.com/go-sql-driver/mysql"
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

type Holder struct {
	Items []It `json:"items"`
}
type It struct {
	Title string `json:"title"`
	URI   string `json:"uri"`
	Date  string `json:"date"`
}

var DB *sql.DB

func main() {

	// // read in json
	// var t Thing

	// input, _ := ioutil.ReadFile("./input.json")
	// //var data interface{}
	// err := json.Unmarshal(input, &t)
	// if err != nil {
	// 	fmt.Println("error marshalling: ", err)
	// }

	//fmt.Println("length of thing: ", len(t.Items))

	// read in json

	// ---------- start test =-------
	// file, err := os.Open("./input.json")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()

	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }

	// if err := scanner.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	// var t Holder

	// input, _ := ioutil.ReadFile("./input.json")
	// //var data interface{}
	// err = json.Unmarshal(input, &t)
	// if err != nil {
	// 	fmt.Println("error marshalling: ", err)
	// }

	// fmt.Println("length of thing: ", len(t.Items))
	//--------------------- end test -------------

	// time.Sleep(time.Second * 10)

	// DB, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/exercise")
	// if err != nil {
	// 	fmt.Println("error opening db: ", err)
	// }

	// err = DB.Ping()
	// if err != nil {
	// 	fmt.Println("could not ping db: ", err)
	// }

	// fmt.Println("starting server")

	// http.HandleFunc("/", search)

	// http.HandleFunc("/hi", hi)

	// log.Fatal(http.ListenAndServe(":8080", nil))

	// UnmarshalJSON()

	h := handler.New(&handler.Config{
		Schema: &schema.Schema,
		Pretty: true,
	})

	// serve Graphiql in-brower editor
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	//graphql api server
	http.Handle("/graphql", h)

	http.ListenAndServe(":8080", nil)

}
