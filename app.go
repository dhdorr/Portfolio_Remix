package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Project struct {
	Title       string
	Description string
}

type ProjectPageData struct {
	Projects []Project
}

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
		// time.Sleep(1 * time.Second)
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("GET /about", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("./templates/about.html")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("sending /about response")
		// time.Sleep(1 * time.Second)
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("GET /projects", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("./templates/projects.html")
		if err != nil {
			log.Fatal(err)
		}

		data := ProjectPageData{
			Projects: []Project{
				{
					Title:       "Github",
					Description: "See what I'm working on now!",
				},
				{
					Title:       "Testing",
					Description: "This is confusing...",
				},
				{
					Title:       "Testing 2",
					Description: "This is still confusing...",
				},
				{
					Title:       "A* Path Visualizer",
					Description: "Try out an interactive pathfinding algorithm visualizer!",
				},
				{
					Title:       "Published Games",
					Description: "Play some published game jam games!",
				},
				{
					Title:       "Word Masters",
					Description: "Can you guess the word of the day in this Wordle clone?",
				},
				{
					Title:       "Pac-Man Portal",
					Description: "Chomp pellets and portal your way to victory in this demo!",
				},
				{
					Title:       "Testing",
					Description: "This is confusing...",
				},
				{
					Title:       "Testing 2",
					Description: "This is still confusing...",
				},
				{
					Title:       "A* Path Visualizer",
					Description: "Try out an interactive pathfinding algorithm visualizer!",
				},
				{
					Title:       "Published Games",
					Description: "Play some published game jam games!",
				},
				{
					Title:       "Word Masters",
					Description: "Can you guess the word of the day in this Wordle clone?",
				},
				{
					Title:       "Pac-Man Portal",
					Description: "Chomp pellets and portal your way to victory in this demo!",
				},
			},
		}
		fmt.Println("sending /projects response")
		// time.Sleep(1 * time.Second)
		tmpl.Execute(w, data)
	})

	// Add this line to serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":8000", nil)
}
