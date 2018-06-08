package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Fprintf(w, "Hello, User!\n")
	})

	http.ListenAndServe(":" + port, nil)
}
