package main

import (
	"math/rand"
	"time"
)

func GenerateMathProblem(level int) string {
	rand.Seed(time.Now().UnixNano())

	// Select the operation based on the level
	operation := rand.Intn(3)

	// Generate two random numbers
	num1 := rand.Intn(100) + 1
	num2 := rand.Intn(100) + 1

	// Set the correct answer based on the operation
	var answer int
	switch operation {
	case 0:
		answer = num1 + num2
	case 1:
		answer = num1 - num2
	case 2:
		answer = num1 * num2
	}

	// Generate the problem string
	problem := fmt.Sprintf("%d %s %d = ?", num1, getOperationString(operation), num2)

	return problem + " " + strconv.Itoa(answer)
}
