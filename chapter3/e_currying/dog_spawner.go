package e_currying

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

func DogSpawnerCurry(breed Breed) func(gender Gender) NameToDogFunc {
	return func(gender Gender) NameToDogFunc {
		return func(name Name) Dog {
			return Dog{
				Name:   name,
				Breed:  breed,
				Gender: gender,
			}
		}
	}
}

// without type alias
//func DogSpawnerCurry(breed Breed) func(Gender) func(Name) Dog {
//	return func(gender Gender) func(Name) Dog {
//		return func(name Name) Dog {
//			return Dog{
//				Breed:  breed,
//				Gender: gender,
//				Name:   name,
//			}
//		}
//	}
//}
