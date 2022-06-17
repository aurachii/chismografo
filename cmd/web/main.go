package main

// Import Packages 

import (

	"log"
	"net/http"
)

func main(){

	//Initialize a new servemux

	mux:=http.NewServeMux()

	//Register the home function as the handler for the "/" URL pattern

	mux.HandleFunc("/",home)

	//Register the show snippet function as the handler for the "/snippet" pattern

	mux.HandleFunc("/snippet",showSnippet)

	//Register the create snippet function as the handler for the "/snippet/create" pattern

	mux.HandleFunc("/snippet/create",createSnippet)

	// Start Server

	log.Println("Starting server on :4000")
	err:=http.ListenAndServe(":4000",mux)
	log.Fatal(err)

}