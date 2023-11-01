package a_type_aliases_for_primitive

type phoneNumber string
type age uint

type Person struct {
	name        string
	age         age
	phonenumber phoneNumber
}

func (a age) valid() bool {
	return a < 120
}

func isValidPerson(p Person) bool {
	return p.age.valid() && p.name != ""
}

func (p *Person) setPhoneNumber(s phoneNumber) {
	p.phonenumber = s
}

func (p *Person) update(name string, phonenumber phoneNumber) {
	p.name = name
	p.phonenumber = phonenumber
}
