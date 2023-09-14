package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Struct
type bill struct {
	name string
	items map[string]float64
	tip float64

}


// make new bill 
func newBill(name string) bill {
	// create new bill object
	b := bill{
		name : name,
		items : map[string]float64{},		// create empty map
		tip: 0,
	}
	// return bill object
	return b
}

// Format the bill - Receiver Function - It limits the function to only Bill object
func (b *bill) format() string {
	fs := "Bill breakdown \n"
	var total float64 = 0;

	// list items
	for key, val := range (*b).items {
		fs += fmt.Sprintf("%-25v ...Rs%v \n", key+":", val )					// add each item to formatted string, -25 to align each item, it adds extra spaces
		total += val															// add the val to total
	}

	// add tip to bill
	fs += fmt.Sprintf("%-25v ...Rs%v \n", "tip:", b.tip)
	total += (*b).tip

	// total
	fs += fmt.Sprintf("%-25v ...Rs%0.2f", "total:", total)						// add total to formatted string

	return fs
}

// Update the Tip - Passing the bill pointer to struct else it will only update the copy of bill
func (b *bill) updateTip(tip float64){
	(*b).tip = tip																// dereferencing the pointer
}

// Add item to the bill
func (b *bill) addItem(name string, price float64){
	(*b).items[name] = price
}

// get Input function 
func getInput(prompt string, r *bufio.Reader) (string, error){					// receives a prompt (string) and reader (pointer to bufio.Reader) as params
	fmt.Print(prompt)															// prints the prompt
	input, err := r.ReadString('\n')											// reads the input on the same line
	return strings.TrimSpace(input), err										// returns trimmed input and error
}

// Create a bill
func createBill() bill{
	reader := bufio.NewReader(os.Stdin)											// create reader from bufio package, NewReader method, source is os.Stdin
	name, _:= getInput("Create a new bill name: ", reader)						// getInput function is call with a prompt to get name from user
	b := newBill(name)															// pass the name to newBill function
	fmt.Println("Created the bill - ", b.name)

	return b																	// return the bill object
}

// save a bill
func (b *bill) save() {
	data := []byte(b.format())													// convert the formatted bill to slice of bytes
	err := os.WriteFile("bills/"+(*b).name+".txt", data, 0644)					// WriteFile method from os package, three params, filename, data, address
	if err != nil{
		panic(err)																// if err, execution stops immediately panic() function
	}
	fmt.Println("Bill was saved in file")
}

// Prompt options
func promptOptions(b bill){
	reader := bufio.NewReader(os.Stdin)																			// create new reader
	opt, _:= getInput("Choose option (a - add item, s - save bill, t - add tip) : ", reader)					// get input from user (choice)

	switch opt {
	
		case "a": 																	
			itemName,_:= getInput("Item Name: ", reader)							// prompt user to enter item name
			price,_ := getInput("Price: ", reader)									// prompt user to enter price

			p, err := strconv.ParseFloat(price, 64)									// parse price as float64, stored in p
			if err != nil{
				fmt.Printf("Price must be a number")								// encountered when price is not a number
				promptOptions(b)
			}
			b.addItem(itemName, p)													// add item to bill
			promptOptions(b)														// promptOptions to user
	
		case "s": 
			b.save()
	
		case "t": 
			tipVal,_ := getInput("Tip amt: ", reader)								// prompt user to enter tip value
			t, err := strconv.ParseFloat(tipVal, 64)								// parse tip as float64, stored in t
			if err != nil {
				fmt.Println("Tip must be a number")									// encountered when tip is not a number
				promptOptions(b)
			}
			b.updateTip(t)															// update the tip
			promptOptions(b)														// promptOptions to user
	
		default : 
			fmt.Println("Invalid choice..")											// invalid choice
			promptOptions(b)
	}
}