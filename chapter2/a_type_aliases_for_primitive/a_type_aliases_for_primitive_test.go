package a_type_aliases_for_primitive

import (
	"fmt"
	"testing"
)

func TestPerson(t *testing.T) {
	p := Person{
		name:        "John",
		phonenumber: "123",
	}
	fmt.Printf("%v\n", p)
}
