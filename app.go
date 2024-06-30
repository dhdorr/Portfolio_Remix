package main

import (
	"errors"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		templ := template.Must(template.ParseFiles("index.html"))

		templ.Execute(w, nil)
	})

	http.HandleFunc("GET /models", func(w http.ResponseWriter, r *http.Request) {
		makeParam := r.URL.Query().Get("make")
		tmpl, err := handleTemplate(makeParam)
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(w, nil)
	})

	http.ListenAndServe(":8000", nil)
}

func handleTemplate(param string) (*template.Template, error) {
	switch param {
	case "audi":
		audi, err := template.ParseFiles("./models/audi.html")
		if err != nil {
			return nil, errors.New("error on options parsing")
		}
		return audi, nil
	case "toyota":
		toyota, err := template.ParseFiles("./models/toyota.html")
		if err != nil {
			return nil, errors.New("error on options parsing")
		}
		return toyota, nil
	case "bmw":
		bmw, err := template.ParseFiles("./models/bmw.html")
		if err != nil {
			return nil, errors.New("error on options parsing")
		}
		return bmw, nil
	default:
		audi, err := template.ParseFiles("/models/audi.html")
		if err != nil {
			return nil, errors.New("error on options parsing")
		}
		return audi, nil
	}
}
