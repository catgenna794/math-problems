
package main

import "math/rand"

func getRandomNumber(min int, max int) int {
  return min + rand.Intn(max-min)
}