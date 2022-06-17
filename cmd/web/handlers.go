package main

// Import Packages 

import (

	"fmt"
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
	
	
	fmt.Fprint(w,"Display a specific snippet with ID ",id, "...")

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