package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Handle Url is Golang")

	res, err := http.Get(url)							// net/http package for handling web requests

	if err != nil{
		panic(err)
	}

	// fmt.Printf("Type of response %v", res)
	
	defer res.Body.Close() 								// make sure to close the http connection, defer makes the line of code to execute at the last

	dataBytes, err := io.ReadAll(res.Body)				// read all the databytes from the body

	if err != nil{
		panic(err)
	}

	content := string(dataBytes)						// convert to string

	error := os.WriteFile("./webrequesttext", []byte(content) , 0644)    	// write into file

	if error != nil {
		panic(error)
	}

	fmt.Println(content)
}