package c_passing_functions_to_functions

import "fmt"

type predicate func(int) bool

func largerThanTwo(i int) bool {
	return i > 2
}

func filter(is []int, p predicate) []int {
	out := []int{}
	for _, i := range is {
		if p(i) {
			out = append(out, i)
		}
	}
	return out
}

func check() {
	ints := []int{1, 2, 3}
	filter(ints, largerThanTwo)
}

// inline function
func checkInLineFunction() {
	ints := []int{1, 2, 3}
	inlineFunction := func(i int) bool {
		return i > 2
	}
	filter(ints, inlineFunction)
}

// anonymous function
func checkAnonymousFunction() {
	filter([]int{1, 2, 3}, func(i int) bool {
		return i > 2
	})
}

// returning functions from functions
func createLargerThanPredicate(threshold int) predicate {
	return func(i int) bool {
		return i > threshold
	}
}

func checkCreateLargerThanPredicate() {
	ints := []int{1, 2, 3}
	largerThanTwo1 := createLargerThanPredicate(2)
	filter(ints, largerThanTwo1)

	//largerThanFive := createLargerThanPredicate(5)
	//largerThanHundred := createLargerThanPredicate(100)
}

// functions in var
var (
	largerThanThree   = createLargerThanPredicate(3)
	largerThanFive    = createLargerThanPredicate(5)
	largerThanHundred = createLargerThanPredicate(100)
)

// functions inside data structures
func checkFunctionInsideDataStructures() {
	ints := []int{1, 2, 3, 6, 101}
	predicates := []predicate{largerThanThree, largerThanFive, largerThanHundred}

	for _, predicate := range predicates {
		fmt.Printf("%v\n", filter(ints, predicate))
	}
}

// functions indisde struct
type ConstraintChecker struct {
	largerThan  predicate
	smallerThan predicate
}

func (c ConstraintChecker) check(input int) bool {
	return c.largerThan(input) && c.smallerThan(input)
}

func checkFunctionInsideStruct() {
	checker := ConstraintChecker{
		largerThan: createLargerThanPredicate(2),
		smallerThan: func(i int) bool {
			return i < 10
		},
	}
	fmt.Printf("%v\n", checker.check(5))
}
