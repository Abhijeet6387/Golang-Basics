package main

import "fmt"

func main() {
	fmt.Println("Maps in Golang")

	

	// Maps - map[key type] value type
	menu := map[string]float64{
		"maggie": 10.99,
		"Pie": 3.99,
		"Tea": 3.99,
		"Coffee": 5.99,
	}

	fmt.Println(menu)									// Printing map
	fmt.Println(menu["Tea"])							// Printing value for key = "Tea"

	menu["Lassi"] = 4.99								// Add value in map with key = "Lassi", value = 4.99
	menu["Tea"] = 2.99									// Updating value in map with key = "Tea"
	delete(menu, "Pie")									// Deleting value with key = "Pie" from map

	price, isPresent := menu["Roti"]					// Returns Two values, If value exists for a key

	if isPresent {
		fmt.Printf("Price of Tea is : %v\n", price)
	} else {
		fmt.Println("Sorry, item is unavailable!")
	}

	for key, val := range menu{
		fmt.Println(key, "-", val)						// Prints the map key-value pair (updated)
	}
}