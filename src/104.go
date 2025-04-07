package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    var num1, num2 int
    fmt.Print("Enter two numbers: ")
    fmt.Scanln(&num1, &num2)
    if num1 > num2 {
        num1, num2 = num2, num1
    }
    for i := 0; i < num1; i++ {
        a := rand.Intn(num2+1) // Generate a random number between 0 and the second largest number
        result := a * a
        fmt.Println(result)
    }
}
