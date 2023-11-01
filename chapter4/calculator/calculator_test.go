package calculator

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	res := Add(5, 10)
	fmt.Println(res)
}
