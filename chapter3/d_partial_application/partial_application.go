package d_partial_application

import "fmt"

func createGreeting(greeting string) func(string) string {
	return func(name string) string {
		return greeting + name
	}
}

func main() {
	firstGreeting := createGreeting("Well, hello there ")
	secondGreeting := createGreeting("Hola ")
	fmt.Println(firstGreeting("Remi"))
	fmt.Println(secondGreeting("Ana"))

	// Partial Application Example

	// Partially apply the addition operation with the first operand 5
	add5 := partialApply(add, 5)

	// Partially apply the subtraction operation with the first operand 10
	subtractFrom10 := partialApply(subtract, 10)

	// Use the partially applied functions
	result1 := add5(3)           // Equivalent to add(5, 3)
	result2 := subtractFrom10(2) // Equivalent to subtract(10, 2)

	fmt.Println("Partial Application Example:")
	fmt.Println("5 + 3 =", result1)
	fmt.Println("10 - 2 =", result2)
}

// Partial Application Example

// Define a binary operation function
type BinaryOperation func(int, int) int

// Partially apply the first operand to create a new function
/*
Partial application refers to the process of fixing a certain number of arguments
of a function and producing a new function with the remaining, unfixed arguments.
*/
func partialApply(operation BinaryOperation, a int) func(int) int {
	return func(b int) int {
		return operation(a, b)
	}
}

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}
