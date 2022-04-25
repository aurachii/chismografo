package main

import (
	"log"
	"net/http"
)

// Handler Function for Home Page

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Eres bien chismoso"))
}

func main() {

	//Initialize a new servemux

	mux := http.NewServeMux()

	//Registering the home function as the handler for the "/" URL pattern

	mux.HandleFunc("/", home)

	log.Println("Starting server on :4000")

	// Start a new web server

	err := http.ListenAndServe(":4000", mux)

	// Log error message and exit

	log.Fatal(err)
}
