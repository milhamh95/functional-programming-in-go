package b_variable_scoping

import (
	"errors"
	"fmt"
)

// location 1
func scoping() {
	// location 2
	b := true
	if b {
		// location 3
		fmt.Println(b)
	}
}

// won't compile
//func scoping1() {
//	{
//		b := true
//	}
//	if b {
//		fmt.Println("b is true")
//	}
//}

func scoping2() {
	s := "hello"
	if true {
		s := "world"
		fmt.Println(s)
	}
	fmt.Println(s)
}

func scoping3() {
	s := "hello"
	if true {
		s = "world"
		fmt.Println(s)
	}
	fmt.Println(s)
}

// won't compile
//func scoping4() {
//	s := "hello"
//	s := "world"
//	fmt.Println(s)
//}

func scoping5() {
	str1, err := func1()
	if err != nil {
		panic(err)
	}

	str2, err := func2()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v %v\n", str1, str2)
}

func func1() (string, error) {
	return "", errors.New("error 1")
}

func func2() (string, error) {
	return "", errors.New("error 2")
}
