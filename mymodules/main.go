package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello mod in golang")
	greeter()
	r := mux.NewRouter()									// create a router using gorilla/mux package
    r.HandleFunc("/", serveHome).Methods("GET")				// route to serveHome, use "GET" methods

	log.Fatal(http.ListenAndServe(":4000", r))				// getting the server up and running, log.Fatal() logs any error if is exists
}

func greeter(){
	fmt.Println("Hey there mod users")
}

func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcome to Golang Home Page</h1>"))
}