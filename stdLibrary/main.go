package main

import (
	"fmt"
	"strings"
	"sort"
)

func main() {
	fmt.Println("Some Standard Library Function in Golang")

	// Strings package
	greeting := "Hello there friends!"
	flag := strings.Contains(greeting, "hello") 				// Return bool value : True or False
	fmt.Printf("Type of flag %T\n", flag)							
	fmt.Println(flag)

	str1 := strings.ToUpper(greeting)							// Returns a new string
	str2 := strings.ReplaceAll(greeting, "Hello", "Hi")         // Returns a new string
	index := strings.Index(greeting, " ")						// Return the index 
	splitArray := strings.Split(greeting, " ")					// Return an array with elements after splitting

	fmt.Println("Strings : ", greeting, str1, str2) 
	fmt.Println("Index : ", index) 
	fmt.Println("Split : ", splitArray, len(splitArray)) 
	
	// Sort Package
	vec1 := []int{4,5,1,3,2}
	vec2 := []string{"abhi", "yashi", "xoayo"}
	
	sort.Ints(vec1)									// sorts the slice of ints
	sort.Strings(vec2)								// sorts the slice of strings		
	idx1 := sort.SearchInts(vec1, 3)				// returns index of 3 in vec1
	idx2 := sort.SearchStrings(vec2, "yashi")		// return index of "yashi" in vec2
	
	fmt.Println("Vec1 : ", vec1)
	fmt.Println("Vec2 : ", vec2)					
	fmt.Println("Index in vec1 : ", idx1)
	fmt.Println("Index in vec2 : ", idx2)

}