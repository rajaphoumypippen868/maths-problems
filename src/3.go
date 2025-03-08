package main

import "math/rand"

func main() {
    rand.Seed(time.Now().UnixNano())
    x := rand.Intn(10)
    y := rand.Intn(10)
    z := x + y
    fmt.Println(z)
}
