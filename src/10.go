// Generating a random math problem using Go
package main

import "math/rand"

func main() {
    // Generate two random numbers between 1 and 10
    num1 := rand.Intn(9) + 1
    num2 := rand.Intn(9) + 1

    // Print the math problem
    fmt.Printf("What is %d x %d?\n", num1, num2)
}
