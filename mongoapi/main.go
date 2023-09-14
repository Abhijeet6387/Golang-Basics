package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Abhijeet6387/mongoapi/router"
)


func main() {
    
	fmt.Println("MongoDB API")
	r := router.Router()									// Calling the routes in router.go file
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000")
	
}