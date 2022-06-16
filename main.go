package main

// Import Packages 

import (
	"log"
	"net/http"
)

// Define handlers

// Response body handler (home page)

func home(w http.ResponseWritter, r *http.Request){
	w.Write([]byte("Bienvenido al Chismografo"))
}

func main(){

	//Initialize a new servemux

	mux:=http.NewServeMux()

	//Register the home function as the handler for the "/" URL pattern

	mux.HandleFunc("/",home)

	// Start Server

	log.Println("Starting server on :4000")
	err:=http.ListenAndServe(":4000",mux)
	log.Fatal(err)

}






