package main

import "math/rand"

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Hello, world!")
}
