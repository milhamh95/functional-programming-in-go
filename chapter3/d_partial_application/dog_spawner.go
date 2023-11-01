package d_partial_application

import "fmt"

type (
	Name          string
	Breed         int
	Gender        int
	NameToDogFunc func(name Name) Dog
)

// define possible breeds
const (
	Bulldog Breed = iota
	Havanese
	Cavalier
	Poodle
)

// define possible geners
const (
	Male Gender = iota
	Female
)

var (
	maleHavaneseSpawner = DogSpwaner(Havanese, Male)
	femalePoodleSpawner = DogSpwaner(Poodle, Female)
)

type Dog struct {
	Name   Name
	Breed  Breed
	Gender Gender
}

func DogSpwaner(breed Breed, gender Gender) NameToDogFunc {
	return func(n Name) Dog {
		return Dog{
			Breed:  breed,
			Gender: gender,
			Name:   n,
		}
	}
}

func createDogsWithPartialApplication() {
	bucky := maleHavaneseSpawner("bucky")
	rocky := maleHavaneseSpawner("rocky")
	tipsy := femalePoodleSpawner("tipsy")

	fmt.Printf("%v\n", bucky)
	fmt.Printf("%v\n", rocky)
	fmt.Printf("%v\n", tipsy)
}

func createDogsWithoutPartialApplication() {
	bucky := Dog{
		Name:   "Bucky",
		Breed:  Havanese,
		Gender: Male,
	}

	rocky := Dog{
		Name:   "Rocky",
		Breed:  Havanese,
		Gender: Male,
	}

	tipsy := Dog{
		Name:   "Tipsy",
		Breed:  Poodle,
		Gender: Female,
	}
	_ = bucky
	_ = rocky
	_ = tipsy
}
