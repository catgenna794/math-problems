package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Generate a random number between 1 and 10
	randomNumber := rand.Intn(10) + 1

	// Print the random number
	fmt.Println("Random Number:", randomNumber)
}
