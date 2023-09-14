package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://lco.dev:3000/learn?coursename=react&paymentid=ghbj456ghb"

func main() {
	fmt.Println("Handling Urls")
	res, err := url.Parse(myurl)											// we need to parse the url first
	
	if err != nil{
		panic(err)
	}

	/* some information any url contains
		fmt.Println(res.Scheme)
		fmt.Println(res.Host)
		fmt.Println(res.Path)
		fmt.Println(res.Port())
		fmt.Println(res.RawQuery)
	*/

	qParams := res.Query()												// get the query parameters
	fmt.Printf("Type of query %T \n", qParams)							// map - key, val pair

	for _, val := range qParams{
		fmt.Println("Value is : ", val)
	}

	// contruct an url
	partsofUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/home",
	}

	newUrl := partsofUrl.String()

	fmt.Println("Another url created: ", newUrl)


}