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

	http.HandleFunc("GET /test", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("./templates/test.html")
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("GET /home", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("./templates/home.html")
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("GET /about", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("./templates/about.html")
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("GET /projects", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("./templates/projects.html")
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(w, nil)
	})

	// Add this line to serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":8000", nil)
}

func handleTemplate(param string) (*template.Template, error) {
	switch param {
	case "audi":
		audi, err := template.ParseFiles("./templates/models/audi.html")
		if err != nil {
			return nil, errors.New("error on options parsing")
		}
		return audi, nil
	case "toyota":
		toyota, err := template.ParseFiles("./templates/models/toyota.html")
		if err != nil {
			return nil, errors.New("error on options parsing")
		}
		return toyota, nil
	case "bmw":
		bmw, err := template.ParseFiles("./templates/models/bmw.html")
		if err != nil {
			return nil, errors.New("error on options parsing")
		}
		return bmw, nil
	default:
		audi, err := template.ParseFiles("./templates/models/audi.html")
		if err != nil {
			return nil, errors.New("error on options parsing")
		}
		return audi, nil
	}
}
