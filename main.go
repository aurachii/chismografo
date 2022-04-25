package main 

import {
	"log"
	"net/http"
}

// Home Handler function (writes a byte slice containing a response body)

func home(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Eres bien chismoso"))
}

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/",home)
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000",mux)
	log.Fatal(err)
}
