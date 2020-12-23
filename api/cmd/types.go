package main

// import (
// 	"encoding/json"
// 	"errors"
// 	"log"
// 	"net/url"
// 	"strings"
// 	"time"
// )

// // "title": "L'avantage d'innover \u00e0 l'\u00e9tat pur",
// // "uri": "http://www.hood.net/about.html",
// // "date": "March 30 1985"

// func (it *Item) UnmarshalJSON(j []byte) error {

// 	var timeFormats = []string{"2006-01-02T15", "January 02", "2006-01-02T15:04:05", "2006-01-02", "January 2 2006", "02 January 2006", "02 Jan 2006", "2006-01-02T15:04:05Z07:00"}

// 	//timeFormats := []
// 	var rawStr map[string]string

// 	if err := json.Unmarshal(j, &rawStr); err != nil {
// 		return err
// 	}

// 	for k, v := range rawStr {
// 		switch strings.ToLower(k) {
// 		case "title":
// 			it.Title = v
// 		case "url":
// 			u, err := url.Parse(v)
// 			if err != nil {
// 				return err
// 			}
// 			it.Uri = *u
// 		case "date":
// 			t, err := parseTime(v, timeFormats)
// 			if err != nil {
// 				return err
// 			}

// 			it.Date = t
// 		}
// 	}
// 	return nil

// }

// func parseTime(input string, formats []string) (time.Time, error) {
// 	for _, format := range formats {
// 		t, err := time.Parse(format, input)
// 		if err == nil {
// 			return t, nil
// 		}
// 	}
// 	log.Println("error parsing time: ", input)
// 	return time.Time{}, errors.New("Unrecognized time format")
// }
