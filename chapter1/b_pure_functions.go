package chapter1

type Person struct {
	Age  int
	Name string
}

// unpure function
// from A -> A1
func changeName(p *Person, newName string) {
	p.Name = newName
}

// pure function
// from A -> B
func changeNamePure(p Person, newName string) Person {
	return Person{
		Age:  p.Age,
		Name: newName,
	}
}
