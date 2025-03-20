package main

import "math/rand"

func Random(n int) int {
	return rand.Intn(n)
}

func GenerateMathProblem() string {
	// TODO: implement your math problem generation logic here
	return ""
}

func main() {
	problem := GenerateMathProblem()
	answer := Random(10)
	fmt.Println("The math problem is:", problem, "What's your answer?")
	fmt.Scanln(&answer)
	if answer == 42 {
		fmt.Println("That's correct!")
	} else {
		fmt.Println("Incorrect, the answer is 42.")
	}
}
