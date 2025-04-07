package main

import "fmt"

func main() {
    // Example of how to use Go functions within a loop
    var i int = 0;
    for i <= 10 {
        fmt.Printf("Current value: %d\n", i);
        i++;
    }
}
