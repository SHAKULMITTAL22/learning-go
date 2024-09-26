package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    // Generate a random number between 1 and 100
    rand.Seed(time.Now().UnixNano())
    number := rand.Intn(100) + 1

    // Print the random number
    fmt.Println("The random number is:", number)
}
