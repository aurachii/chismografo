package main

// Import Packages 

import (

	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Define handlers

// Home handler function

func home(w http.ResponseWriter, r *http.Request){
	
	// Restrict the "/" pattern

	if r.URL.Path != "/"{
		http.NotFound(w,r)
		return
	}
	
	w.Write([]byte("Bienvenido al Chismografo"))
}
// Show snippet handler function

func showSnippet(w http.ResponseWriter, r *http.Request){

	
	// Retrieve the id parameter from the URL query string and convert the string value into a integer

	id,err:=strconv.Atoi(r.URL.Query().Get("id"))

	//Check whether it contains a positive integer value

	if err !=nil || id<1{
		http.NotFound(w,r)
		return
	}
	
	
	fmt.Fprint(w,"Display a specific snipper with id %d...",id)

}

// Create snippet handler function

func createSnippet(w http.ResponseWriter, r *http.Request){
	
	
	// Restrict http request method to  POST 

	if r.Method != "POST"{

		w.Header().Set("Allow","POST")
		// Send non-200 status code 
		http.Error(w,"Method not Allowed",405)
		return
	}

	w.Write([]byte("Create a snippet"))
}
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






