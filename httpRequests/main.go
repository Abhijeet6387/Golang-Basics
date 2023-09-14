package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const geturl = "https://jsonplaceholder.typicode.com/posts/1"
const posturl = "https://jsonplaceholder.typicode.com/posts"

func getRequest(url string){
	res, err := http.Get(url)
	if err != nil{
		panic(err)
	}
	defer res.Body.Close() 													// make sure to close the connection

	fmt.Println("Status Code: ", res.StatusCode)
	fmt.Println("Content Length: ", res.ContentLength)
	
	content, err := io.ReadAll(res.Body)

	if err != nil{
		panic(err)
	}
	
	// Using strings.Builder Write method to create response string
	var respString strings.Builder
	byteCount, _ := respString.Write(content)								// returns a bytecount and error
	fmt.Println("Byte Count: ", byteCount)									// print the bytecount
	fmt.Println("Content: ", respString.String())							// convert the bytes to string using .String() method
	// fmt.Println("Content: ", string(content))
}


func postRequest(url string, reqBody *strings.Reader){
	res, err := http.Post(url, "application/json", reqBody)

	if err != nil{
		panic(err)
	}

	defer res.Body.Close()
	content, err := io.ReadAll(res.Body)

	if err != nil{
		panic(err)
	}
	var respString strings.Builder
	byteCount, _ := respString.Write(content)
	fmt.Println("Bytecount: ", byteCount)
	fmt.Println("Content: ", respString.String())
	// fmt.Println(string(content))
}


func postFormRequest(url string, formData url.Values){
	res, err := http.PostForm(url, formData)
	if err != nil{
		panic(err)
	}

	defer res.Body.Close()

	content, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	var respString strings.Builder
	byteCount, _ := respString.Write(content)
	fmt.Println("Bytecount: ", byteCount)
	fmt.Println("Content: ", respString.String())
	// fmt.Println(string(content))

}


func main() {
	fmt.Println("Http Requests in Golang")

	// get request
	getRequest(geturl)
	fmt.Println("--------------------------------")


	// fake json payload for post request
	reqBody := strings.NewReader(`
		{
			"title":"foo",
			"body":"bar",
			"userId":1
		}
	`)
	// post request
	postRequest(posturl, reqBody)
	fmt.Println("--------------------------------")

	
	// creating form data using url.Values{}
	formData := url.Values{}
	formData.Add("title", "foo2")
	formData.Add("body", "bar2")
	formData.Add("userId", "2")
	// post request for form data
	postFormRequest(posturl, formData)
	fmt.Println("--------------------------------")


}