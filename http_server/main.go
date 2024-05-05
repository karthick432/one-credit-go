package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	http.HandleFunc("/name", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "KARTHICK K")
	})

	http.HandleFunc("/rollno", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "7376212AL123")
	})

	http.HandleFunc("/dept", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "AIML")
	})

	http.HandleFunc("/colour", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "BLACK")
	})

	fmt.Printf("Server running (port=8080), route:http://localhost:8080/KARTHI\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
