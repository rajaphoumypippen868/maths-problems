package main

import (
    "fmt"
)

func main() {
    // Define a function to check if a number is prime or not
    func isPrime(num int) bool {
        if num <= 1 {
            return false
        }
        for i := 2; i*i <= num; i++ {
            if num%i == 0 {
                return false
            }
        }
        return true
    }

    // Example usage of the prime checking function
    primes := make(map[int]bool)
    primes[3] = true
    primes[11] = true

    for _, prime := range primes {
        fmt.Printf("%d is prime: %t\n", prime, isPrime(prime))
    }
}
