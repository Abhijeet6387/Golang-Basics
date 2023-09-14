package main

import "fmt"

func main() {
	fmt.Println("Data Types in Golang")

	// Strings
	var nameOne string = "Abhijeet"  	
	var nameTwo = "abhijeet"  		 	
	var nameThree string			 	// empty string
	nameFour := "Mishraa"				// can't be used outside function (shorthand)

	fmt.Println(nameOne, nameTwo, nameThree, nameFour)

	nameOne = "Updated"					// Update the value, can't change the type of variable
	nameThree = "FirstTime"				// Initialize value after declaring

	fmt.Println(nameOne, nameTwo, nameThree, nameFour)


	// Integers
	var age1 int = 20
	var age2 = 30
	age3 := 40

	fmt.Println(age1, age2, age3)

	var numone int8 = 25     			// 8 bits integer- 128 to 127  (possible 8, 16, 32, 64 bits type)
	var numtwo int16 = 32767 			// 16 bits integer- 32768 to 32767
	var numthree uint8 = 255			// unsigned interger- non-negative 0 to 255 (possible 8, 16, 32, 64 bits type)

	fmt.Println(numone, numtwo, numthree)

	
	// Floats
	var score1 float32 = -1.55			// 32 bit float type
	var score2 float64 = 278764.884		// 64 bit float type
	score3 := 3635865.276836725         // By default 64 bit type

	fmt.Println(score1, score2, score3)

	// Arrays
	var ages [3]int = [3]int{20,30,40}    		// way to declare arrays in Go, they can't have variable size
	names := [2]string{"abhi","ritik"}			// Can't add or delete items from these arrays
	names[1] = "mishra"							// Can update the values
	fmt.Println(ages, len(ages))				// len() gives the length of the array
	fmt.Println(names, len(names))

	// Slices (use arrays under the hood)
	var scores = []int{100,50,20}
	fmt.Println("Intial", scores)
	scores[1] = 90
	fmt.Println("Updated", scores)

	scores = append(scores, 85)									// append() does'nt actually append instead it returns a new slice 
	fmt.Println(scores, len(scores))

	i := 1														// Delete value at index 1, using append in Slices
	scores = append(scores[:i], scores[i+1:]... )
	fmt.Println("Line 61", scores)

	// range in slices
	rangeScore1 := scores[1:3] 									// include pos 1 & 2, not 3
	rangeScore2 := scores[1:]									// include from pos 1 till end
	rangeScore3 := scores[:3]									// include from start till 2, not 3
	fmt.Println(rangeScore1, rangeScore2, rangeScore3)

	// Bool
	var flag1 bool = true
	flag2 := false

	fmt.Printf("Flag1 is %v and Flag2 is %v \n", flag1, flag2)
	
}