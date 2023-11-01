package c_closures

import "fmt"

func main() {
	greetingFunc := createGreeting()
	response := greetingFunc("Ana")
	fmt.Println(response)
}

func createGreeting() func(string2 string) string {
	s := "Hello "
	return func(name string) string {
		return s + name
	}
}
