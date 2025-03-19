package main

import "fmt"

func main() {
	const num = 50
	const denom = 10

	ratio := float64(num) / float64(denom)
	fmt.Println("The ratio is", ratio)
}
