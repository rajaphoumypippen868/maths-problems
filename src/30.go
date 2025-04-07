package main

import (
    "fmt"
)

func main() {
    var s string
    fmt.Print("Enter a sentence: ")
    fmt.Scanln(&s)
    fmt.Println("You entered:", s)
}
