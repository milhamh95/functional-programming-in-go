package b_referential_transparency

import (
	"fmt"
	"math/rand"
)

/*
referential transparency

function is referential tranparent if you can
replace the function call with its output.
without chnging the result of the program.

x = 1 + (2 * 2)
x = 1 + 4

replace (2 * 2) with 4
*/

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Printf("%v\n", add(10, add(10, 5)))
	fmt.Printf("%v\n", add(10, 15))
}

type Player string

const (
	PlayerOne Player = "Remi"
	PlayerTwo Player = "Yvonne"
)

func selectStartingPlayer() Player {
	randomized := rand.Intn(2)
	switch randomized {
	case 0:
		return PlayerOne
	case 1:
		return PlayerTwo
	}
	panic("No further player available")
}

func PlayerSelectPure(i int) (Player, error) {
	switch i {
	case 0:
		return PlayerOne, nil
	case 1:
		return PlayerTwo, nil
	}

	return Player(""), fmt.Errorf("no player matching input: %v", i)
}
