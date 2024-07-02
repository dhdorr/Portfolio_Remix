package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		templ := template.Must(template.ParseFiles("index.html"))

		templ.Execute(w, nil)
	})

	http.HandleFunc("GET /home", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("./templates/home.html")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("sending /home response")
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("GET /about", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("./templates/about.html")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("sending /about response")
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("GET /projects", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("./templates/projects.html")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("sending /projects response")
		tmpl.Execute(w, nil)
	})

	// Add this line to serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":8000", nil)
}
