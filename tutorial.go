package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

//a simple function
func calculateSum(w http.ResponseWriter, r *http.Request) int {
	return 7
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})


	log.Fatal(http.ListenAndServe(":8081", nil))

}
