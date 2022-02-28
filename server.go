package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("Starting server at port 8080\n")
	http.HandleFunc("/", GreetingPage)
	http.ListenAndServe(":8080", nil)
}

func GreetingPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the bare minimum.")
}