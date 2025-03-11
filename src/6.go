package main

import "math/rand"

func main() {
	rand.Seed(9)
	fmt.Println("Solution: ", rand.Intn(10))
}
