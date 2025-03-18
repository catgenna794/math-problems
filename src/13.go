
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(9)

	num1 := rand.Intn(10)
	num2 := rand.Intn(10)

	fmt.Println("What is", num1, "times", num2, "?")
}