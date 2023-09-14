package main

import "fmt"

// Function to update name
// At line 9, *string represents that name is a pointer to a string
// At line 10, *name is dereferencing the pointer and updating it's value

func updateName(name *string){
	*name = "Abhi"
}
func updateAge(age *int){
	*age = 23
}


func main() {

	/* 
		NOTE - 
		-> Strings, Ints, Floats, Booleans, Arrays and Structs are Non-Pointer Value  (Can't be updated by reference unless pointers are used)
		-> Slices, Maps, Function can be grouped as Pointer Wrapper Values	(Can be updated by reference automatically)
	*/

	fmt.Println("Pointers in Golang")


	Name := "Abhishek"
	Age := 10
	fmt.Printf("Intially Name is : %v, Age is : %v\n", Name, Age)
	
	
	ptr1 := &Name																// Pointer to a string
	ptr2 := &Age																// Pointer to an int
														
	// updateName(Name)           											    // This does'nt updates the name
	// fmt.Println(Name) 
	
	updateName(ptr1)															// Calling updateName function and passing the reference of firstName as parameter
	updateAge(ptr2)																// Calling updateAge function and passing the reference of Age as parameter

	fmt.Printf("Updated Name is : %v, Age is : %v\n", Name, Age)	

}