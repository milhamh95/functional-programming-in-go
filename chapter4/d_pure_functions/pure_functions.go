package d_pure_functions

import "fmt"

/*
1. Avoid global state
- avoid the use of const and var blocks at the pckage level as much as possible
- when use these blocks, there's a good change that the program is relied
upon by some functions.
- generating either side effects or having non-deterministic program execution
*/

//var (
//	name = "Remi"
//)
//
//func sayHello() string {
//	return fmt.Sprintf("hello %s", name)
//}
//
//func main() {
//	sayHello()
//}

func sayHello(name string) string {
	return fmt.Sprintf("hello %s", name)
}

func main() {
	sayHello("Remi")
}

/*
2. Separate pure and impure functionality
*/

// each function does exactly one thing

// not a pure function
// because we write num value to standard output.
// variant: write to filesystem / db
func add(a, b int) int {
	sum := a + b
	fmt.Println(sum)
	return sum
}

/*
- not pure function
- try to do 2 things: create user struct and validate the password
*/
//func createUser(username, password string) {
//	u := User{username, password}
//	if u.validPassword() {
//		userDb.save(u)
//	} else {
//		panic("invalid password")
//	}
//}

/*
refactor version

func signup(username, password string) {
	user, err := createuser(username, password)
	if err == nil {
		panic("could not create account")
	}

	saveUser(user)
}

func createUser(username, password string) (User, error) {

	u := User{username, password}
	if u.validPassword() {
		return u, nil
	}

	return User{}, Errors.new("invalid password")
}

func saveUser(u User) {
	userDb.save(u)
}
*/
