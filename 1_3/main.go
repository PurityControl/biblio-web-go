package main

import (
	"fmt"
	"net/http"
	"html/template"
)

func main() {
	templates := template.Must(template.ParseFiles("templates/index.html"))
	
	// test code to see everything okay
	// TODO find out about unit testing
	fmt.Println("Hello, Go Web Development!")

	http.HandleFunc("/", func(w http.ResponseWriter, r  *http.Request) {
		if err:= templates.ExecuteTemplate(w, "index.html", nil);
		err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// localhost:8080 not needed as localhost implied if missing
	fmt.Println(http.ListenAndServe(":8080", nil))
}
