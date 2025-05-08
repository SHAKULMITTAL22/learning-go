package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println("Enter the radius of the circle:")
    var radius float64
    fmt.Scanln(&radius)
    area := math.Pi * math.Pow(radius, 2)
    fmt.Println("Area of the circle:", area)
}
