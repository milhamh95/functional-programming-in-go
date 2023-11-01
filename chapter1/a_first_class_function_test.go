package chapter1

import (
	"fmt"
	"testing"
)

func TestFilter(t *testing.T) {
	is := []int{1, 1, 2, 3, 5, 8, 13}
	larger := filter(is, largerThan5)
	fmt.Printf("%v", larger)
}
