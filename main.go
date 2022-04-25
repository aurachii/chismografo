package main

import (
	"log"
	"net/http"
)

// Handler Function for Home Page

func home(w http.ResponseWriter, r *http.Request) {

	// This checks whether the current URL path exactly matches "/". If it doesn't call
	// then it will send a 404 response to the client.

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	//This writes home page message

	w.Write([]byte("Bienvenido al chismografo, he de informarte que eres un chismoso"))

}

// Handler Function for Showing Snippets

func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Se supone que este es un chisme"))
}

// Handler Function for Creating Snippets

func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Agregar un nuevo chisme"))
}

func main() {

	//Initialize a new servemux

	mux := http.NewServeMux()

	//Registering the home function as the handler for the "/" URL pattern

	mux.HandleFunc("/", home)

	//Registering the showSnippet function as the handler for the "/snippet" URL pattern

	mux.HandleFunc("/snippet", showSnippet)

	//Registering the createSnippet function as the handler for the "/create" URL pattern

	mux.HandleFunc("/create", createSnippet)

	// Start a new web server
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)

	// Log error message and exit

	log.Fatal(err)
}
