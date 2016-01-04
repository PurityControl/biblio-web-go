package main

import (
	"fmt"
	"net/http"
	"html/template"
	
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Page struct {
	Name string
	DBStatus bool
}

func main() {
	templates := template.Must(template.ParseFiles("templates/index.html"))

	db, _ := sql.Open("sqlite3", "dev.db")
	
	// test code to see everything okay
	// TODO find out about unit testing
	fmt.Println("Hello, Go Web Development!")

	http.HandleFunc("/", func(w http.ResponseWriter, r  *http.Request) {
		// returns "Hello, Gopher" or greets you if a ?name="???"
		// is passed in with the web request.
		p := Page{Name: "Gopher"}
		if name := r.FormValue("name"); name != "" {
			p.Name = name
		}
		p.DBStatus = db.Ping() == nil
		
		if err:= templates.ExecuteTemplate(w, "index.html", p);
		err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// localhost:8080 not needed as localhost implied if missing
	fmt.Println(http.ListenAndServe(":8080", nil))
}
