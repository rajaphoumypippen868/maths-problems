package main

import "fmt"

func main() {
    var num1 int = 5;
    var num2 int = 3;

    sum := num1 + num2;
    product := num1 * num2;

    fmt.Printf("The sum of %d and %d is: %d\n", num1, num2, sum)
    fmt.Printf("The product of %d and %d is: %d\n", num1, num2, product)
}
