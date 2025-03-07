package main

import "math/rand"

func GenerateMathProblem() string {
	// Generate a random math problem using random numbers between 1 and 10
	a := rand.Intn(10) + 1
	b := rand.Intn(10) + 1
	operator := rand.Intn(2)
	if operator == 0 {
		return fmt.Sprintf("What is %d + %d?", a, b)
	} else {
		return fmt.Sprintf("What is %d - %d?", a, b)
	}
}
