package main

import "fmt"

func main() {
	fmt.Println("Loops in Golang")

	// Loops
	
	/* Just like while loop, for keyword is used in Go */
		x := 0
		for x < 5 {
			fmt.Printf("X equals %v \n", x)
			x++
		}
	

	/* Just like typical for loop in C++ */
		for x:=0; x<5; x++ {
			fmt.Printf("X equals %v \n", x)
		}
	

	/* Looping through slice */
		names := []string{"abhi", "yashi", "xoayo", "abhbd", "xyzh"}
		for i:=0; i<len(names);i++ {
			fmt.Println(names[i])
		}
	

	/* Just like For in loop, use range */
		
		for index, value := range names {
			fmt.Printf("Value at index %v is %v \n", index, value)
		}
	

	/* If index is not required */
		for _,value := range names {
			fmt.Printf("Value is %v \n", value)
			// Value can't be updated here as it like a local copy of original value
		}
	

}