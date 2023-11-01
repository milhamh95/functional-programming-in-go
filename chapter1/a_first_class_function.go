package chapter1

type predicate func(int) bool

func filter(is []int, condition predicate) []int {
	out := []int{}

	for _, i := range is {
		if condition(i) {
			out = append(out, i)
		}
	}

	return out
}

func largerThan5(i int) bool {
	return i > 5
}
