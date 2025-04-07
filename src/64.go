package main

import "fmt"

func main() {
    // Define a function to print the numbers from 1 to n
    func printNumbers(n int) {
        for i := 1; i <= n; i++ {
            fmt.Println(i)
        }
    }

    // Call the function with user-defined number and display it
    println("Enter a number: ")
    num, _ := fmt.Scanln(&n)
    if n < 1 {
        fmt.Println("Please enter a positive integer")
    } else {
        printNumbers(n)
    }
}
