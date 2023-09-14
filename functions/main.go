package main

import (
	"fmt"
	"math"
	"strings"
)

// Creating functions outside main
func sayHi(name string){
	fmt.Printf("Hi %v \n", name)
}
func sayBye(name string){
	fmt.Printf("Bye %v \n", name)
}
func cycle(names []string, fn func(string)){
	for _,val := range names{
		fn(val)										// call function fn for each value of slice
	}
}

// function with return type float64
func circleArea(radius float64) float64{       
	return math.Pi * radius * radius
}

// function with multiple return value - It takes the name (string) as paramter and returns the Initials (Capitalized)
func getInitials(name string) (string, string) {
	
	s := strings.ToUpper(name)						// Convert the name string in UpperCasem - returns a string
	namesArr := strings.Split(s, " ")				// Split the string when a space is encountered - returns a slice

	var ans []string								// Initializing a slice to store the initials "ans"

	for _,val := range namesArr {					// For each val in namesArr slice, we append the val in ans
		ans = append(ans, val[:1])					// Include only the first Character
	}

	if len(ans) > 1{								// If ans consists of more than 1, return both
		return ans[0], ans[1]
	}

	return ans[0], ""								// Else return the first and an empty string in second
}


func main() {
	fmt.Println("Functions in Golang")

	sayHi("Adam")    												// Prints Hi Adam
	sayBye("Adam")													// Prints Bye Adam
	
	names := []string{"Adam", "Eve", "Matthew","Virat", "KLRahul"}
	
	// cycle(slice, function)  
	cycle(names, sayHi)												// Iterates through the slice and prints Hi for each
	cycle(names, sayBye)											// Iterates through the slice and prints Bye for each

	a1 := circleArea(15.786)										// returns the area of circle with r = 15.786
	a2 := circleArea(16.897)										// returns the area of circle with r = 16.897

	fmt.Printf("Area of circle 1 is %0.3f \n", a1)
	fmt.Printf("Area of circle 2 is %0.3f \n", a2)
	
	// function with multiple return value
	fn, sn := getInitials("abhijeet mishra")						// returns A M
	fn1, sn1 := getInitials("abhijeet")								// returns A [space]
	
	fmt.Println(fn, sn)
	fmt.Println(fn1, sn1)
}