package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		templ := template.Must(template.ParseFiles("index.html"))

		templ.Execute(w, nil)
	})

	http.HandleFunc("GET /models", func(w http.ResponseWriter, r *http.Request) {
		makeParam := r.URL.Query().Get("make")
		switch makeParam {
		case "audi":
			audi, err := template.ParseFiles("./models/audi.html")
			if err != nil {
				fmt.Println("error on options parsing")
				return
			}
			audi.Execute(w, nil)
		case "toyota":
			toyota, err := template.ParseFiles("./models/toyota.html")
			if err != nil {
				fmt.Println("error on options parsing")
				return
			}
			toyota.Execute(w, nil)
		case "bmw":
			bmw, err := template.ParseFiles("./models/bmw.html")
			if err != nil {
				fmt.Println("error on options parsing")
				return
			}
			bmw.Execute(w, nil)
		default:
			audi, err := template.ParseFiles("/models/audi.html")
			if err != nil {
				fmt.Println("error on options parsing")
				return
			}
			audi.Execute(w, nil)
		}
		// 		optionsStr := fmt.Sprintf(`<option value='A1'>A1</option>
		// <option value='A4'>A4</option>
		// <option value='A6'>A6</option>`)

		// 		templ, err := template.New("t").Parse(optionsStr)
		// 		if err != nil {
		// 			fmt.Println("error on options parsing")
		// 			return
		// 		}
		// 		templ.Execute(w, nil)
	})

	http.ListenAndServe(":8000", nil)
}
