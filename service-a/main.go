package main

import (
	"fmt"
	"net/http"

	h "github.com/sithuaung/two-services/helper"
)

func main() {
	h.Log()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World! Welcome to the simple Golang service.")
	})

	port := ":8000"

	fmt.Printf("Starting server on port %s...\n", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
