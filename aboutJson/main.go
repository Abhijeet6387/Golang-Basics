package main

import (
	"encoding/json"
	"fmt"
)

// We can pass third parameter in structs, as how these key will reflect in JSON data
type course struct{
	Name string	`json:"coursename"`					// Name will reflect as coursename
	Price int `json:"price"`						// Price will reflect as price
	Platform string	`json:"website"`				// Platform will reflect as website
	Password string	`json:"-"`						// We can omit sensitive information using "-"
	Tags []string `json:"tags,omitempty"`			// We can omit empty entries using omitempty (without space)
}

// Encoding JSON Data
func convertToJson(){

	courses := []course{
		{"Reactjs", 500, "Youtube","abc123", []string{"web-dev","js","react"}},
		{"Nextjs", 600, "Youtube","bcd123", nil},
		{"Nodejs", 800, "Youtube","xys123", []string{"backend","js","dev"}},
	}

	// JSONData, err := json.Marshal(courses)					// Marshal() encodes and returns data in JSON format 
	JSONData, err := json.MarshalIndent(courses,"","\t")		// MarshalIndent() encode and return data in JSON format with proper Indentation
	
	if err != nil{
		panic(err)
	}

	fmt.Printf("%s\n", JSONData)

}

// Decoding JSON Data
func convertJsonData(){
	jsonDataFromWeb := []byte(
		`{
                "coursename": "Reactjs",
                "price": 500,
                "website": "Youtube",
                "tags": [
                        "web-dev",
                        "js",
                        "react"
                ]
        }`,
	)
	
	// To store decoded data
	var courses course 
	
	// To store data in key value pair
	var courseMap map[string]interface{} 								// use interface because we are unaware of the type of value
	
	// check if jsonDataFromWeb is valid or not
	check := json.Valid(jsonDataFromWeb)

	if check{
		fmt.Println("Valid JSON")
		err := json.Unmarshal(jsonDataFromWeb, &courses)				// Unmarshal() decodes the JSON data, & stores the data in pointer to Interface (structs), here courses
		error := json.Unmarshal(jsonDataFromWeb, &courseMap)			// for data in key, value pair format

		if err != nil || error != nil{
			panic(err)
		}

		fmt.Printf("%#v\n", courses) 				// To print interface, we use %#v as format specifier
		fmt.Println("------------------")
		fmt.Printf("%#v\n", courseMap)
		fmt.Println("------------------")

		for key, val := range courseMap{
			fmt.Printf("key : %v, val : %v, type : %T \n", key, val, val )			// we can iterate through it and print the key value pairs
		}

	} else{
		fmt.Println("Invalid JSON")
	}
}
func main() {
	fmt.Println("More about JSON")
	convertToJson() 							// encoding JSON
	fmt.Println("------------------")
	convertJsonData() 							// decoding JSON
}