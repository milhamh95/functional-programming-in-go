package a_pure_impure

import "math/rand"

// pure function
// return consistent result
func add(a, b int) int {
	return a + b
}

// impure function
// won't return consistent result
func rollDice() int {
	return rand.Intn(6)
}
