package main

import (
	"fmt"
	"math/big"
	// "math/rand"
	"crypto/rand"
)

func main(){
	
	fmt.Println("Random Numbers in Golang")
	
	// fmt.Println(rand.Intn(1000)) 														// Using math/rand package
	myrandNo, err := rand.Int(rand.Reader, big.NewInt(1000)) 								// Using crypto/rand package

	if err != nil{
		panic(err)
	}
	
	fmt.Println("Random number generated : ", myrandNo)
}