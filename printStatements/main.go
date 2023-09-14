package main

import ( "fmt")

func main(){
	fmt.Println("Hello, World!")
	
	// Print Statements
	age := 25
	name := "Abhijeet"
	score := 25.89897

	fmt.Print("Hello")  													// prints on the same line until \n is added
	fmt.Print("User \n")
	fmt.Println("Hello User!")												// prints on a new line

	fmt.Println("Age : ", age, "Name : ", name)
	fmt.Printf("Age is %v and name is %v \n", age, name)  					// Formatted string (order matters), %_ = format specifier
	fmt.Printf("Age is %v and name is %q \n", age, name) 					// %q puts quotes around strings
	fmt.Printf("age is of Type %T \n", age)  								// %T, prints the type of variable age
	fmt.Printf("you scored %f \n", score)  									// %f, prints the floats variable
	fmt.Printf("you scored %0.3f \n", score) 								// %0.xf, prints floats upto x decimal points

	
	var str = fmt.Sprintf("Age is %v and name is %v \n", age, name)         // Sprintf - save formatted strings
	fmt.Println("Saved string : ", str)
}