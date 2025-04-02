package main

import "fmt"

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    fmt.Println("The sum of all numbers in the array is:", sumArray(numbers))
}

func sumArray(numbers []int) int {
    total := 0
    for _, number := range numbers {
        total += number
    }
    return total
}
