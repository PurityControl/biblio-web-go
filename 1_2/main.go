package main

import (
	"fmt"
	"net/http"
)

func main() {
	// test code to see everything okay
	// TODO find out about unit testing
	fmt.Println("Hello, Go Web Development!")

	http.HandleFunc("/", func(w http.ResponseWriter, r  *http.Request) {
		fmt.Fprintf(w, "Hello, Go Web Development")
	})

	// localhost:8080 not needed as localhost implied if missing
	fmt.Println(http.ListenAndServe(":8080", nil))
}
