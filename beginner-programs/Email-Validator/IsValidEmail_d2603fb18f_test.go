package main

import (
    "fmt"
    "math"
)

func main() {
    // Declare the variables
    var a, b, c float64
    var discriminant float64

    // Prompt the user to enter the values of a, b, and c
    fmt.Print("Enter the values of a, b, and c: ")
    fmt.Scanln(&a, &b, &c)

    // Calculate the discriminant
    discriminant = b*b - 4*a*c

    // Check if the discriminant is positive, negative, or zero
    if discriminant > 0 {
        // Calculate the two real and distinct roots
        x1 := (-b + math.Sqrt(discriminant)) / (2 * a)
        x2 := (-b - math.Sqrt(discriminant)) / (2 * a)
        fmt.Printf("The two real and distinct roots are: %.2f and %.2f\n", x1, x2)
    } else if discriminant == 0 {
        // Calculate the single real root
        x := -b / (2 * a)
        fmt.Printf("The single real root is: %.2f\n", x)
    } else {
        // Print a message indicating that there are no real roots
        fmt.Println("There are no real roots for the given values.")
    }
}
