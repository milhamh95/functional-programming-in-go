package e_todo

import "os"

type authorizationFunc func() bool

type Db struct {
	AuthorizationF authorizationFunc
}

func (d *Db) IsAuthorized() bool {
	return d.AuthorizationF()
}

func NewDB() *Db {
	return &Db{
		AuthorizationF: argsAuthorization,
	}
}

func argsAuthorization() bool {
	user := os.Args[1]
	if user == "admin" {
		return true
	}

	return false
}
