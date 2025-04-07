package main

import (
    "fmt"
)

func main() {
    // Example Go code to solve mathematical problems
    var sum float64 = 0.0
    for i := 1; i <= 5; i++ {
        if i%2 == 0 {
            sum += i
        }
    }

    fmt.Println("The sum of even numbers from 1 to 5 is:", sum)
}
