package main

import (
	"fmt"
	"os"
)

func readFile(filename string){
	dataByte, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("Text is the file is : ", string(dataByte))
}

func main() {
	fmt.Println("Files Folder")
	content := "Hi, My name is Abhijeet Shankar Mishra. I am an upcoming Software Development Engineer"

	file, err := os.Create("./mytextfile.txt")

	if err != nil{
		panic(err)
	}

	error := os.WriteFile("./mytextfile.txt", []byte(content),  0644)

	if error != nil{
		panic(err)
	}


	defer file.Close()

	readFile("./mytextfile.txt")

}
