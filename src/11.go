package main

import "math/rand"

func GenerateMathProblem() string {
	// Generate two random numbers between 1 and 10
	num1 := rand.Intn(10) + 1
	num2 := rand.Intn(10) + 1

	// Generate a random operation (+, -, *, /)
	op := "+"
	switch rand.Intn(4) {
	case 0:
		op = "+"
	case 1:
		op = "-"
	case 2:
		op = "*"
	case 3:
		op = "/"
	}

	// Generate a random answer based on the operation
	answer := 0
	switch op {
	case "+":
		answer = num1 + num2
	case "-":
		answer = num1 - num2
	case "*":
		answer = num1 * num2
	case "/":
		if num2 == 0 {
			return ""
		} else {
			answer = num1 / num2
		}
	}

	// Format the problem and answer as a string
	problem := fmt.Sprintf("%d %s %d", num1, op, num2)
	return fmt.Sprintf("What is %s? %d", problem, answer)
}
