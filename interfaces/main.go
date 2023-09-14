package main

import "fmt"

func main() {
	fmt.Println("Interfaces in Golang..")
	shapes := []shape{							// slice of shapes
		square{length: 15.2},
		circle{radius: 4.56},
		square{length: 3.45},
		circle{radius: 12.3},
	}

	for _, val := range shapes{					// traversing through shapes and printing info for each
		printInfo(val)
		fmt.Println("----")
	}
}