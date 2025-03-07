package main

import (
	"fmt"
	"math/rand"
)

func generateRandomNumber() int {
	return rand.Intn(10) + 1
}

func main() {
	num := generateRandomNumber()
	fmt.Println("Your random number is:", num)
}
