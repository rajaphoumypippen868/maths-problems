 package main

 import "math/rand"

 func main() {
     rand.Seed(9) // seed with a fixed value for deterministic results
     x := rand.Intn(10) // generate a random number between 0 and 10, inclusive
     y := rand.Intn(10) // generate another random number between 0 and 10
     fmt.Println("The sum of", x, "and", y, "is", x+y) // print the result
 }