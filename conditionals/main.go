package main

import "fmt"

func main() {
	fmt.Println("Conditionals in Golang")

	// Boolean and Conditionals
	y := 10

	fmt.Println(y <= 9)      	// returns false
	fmt.Println(y >= 10)     	// returns true
	fmt.Println(y == 10)		// returns true
	fmt.Println(y != 10)  		// returns false

	// If, Else if, Else statements
	if y < 10{
		fmt.Println("y is less than 10")
	} else if y == 10{
		fmt.Println("y is equal to 10")
	} else{
		fmt.Println("y is greater than 10")
	}

	// conditionals with loops
	names:= []string{"abhi", "yashi", "xoayo", "abhbd", "xyzh"}
	for idx, val := range names{
		if idx == 1 {
			fmt.Println("index: ", idx)
			continue
		} 
		if idx == 3{
			fmt.Println("index: ", idx)
			break
		}
		
		fmt.Printf("Value at index %v is %v \n", idx, val)
	}
}