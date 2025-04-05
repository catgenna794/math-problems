package main

import (
    "fmt"
    "testing"
)

func main() {
    // Test case 1: Add two numbers
    result := add(3, 4)
    fmt.Println("Result of adding 3 and 4:", result) // Output: Result of adding 3 and 4: 7

    // Test case 2: Subtract two numbers
    result = subtract(8, 4)
    fmt.Println("Result of subtracting 4 from 8:", result) // Output: Result of subtracting 4 from 8: 4

    // Test case 3: Multiply two numbers
    result = multiply(5, 6)
    fmt.Println("Result of multiplying 5 and 6:", result) // Output: Result of multiplying 5 and 6: 30

    // Test case 4: Divide two numbers
    result = divide(12, 3)
    fmt.Println("Result of dividing 12 by 3:", result) // Output: Result of dividing 12 by 3: 4.0

    // Test case 5: Calculate square root of a number
    result = sqrt(25)
    fmt.Println("Square root of 25:", result) // Output: Square root of 25: 5.0

    // Test case 6: Calculate factorial of a number (n!)
    result = factorial(3)
    fmt.Println("Factorial of 3:", result) // Output: Factorial of 3: 6
}
