package b_type_aliases_for_functions

type predicate func(int) bool

func filter(is []int, p predicate) []int {
	out := []int{}

	for _, i := range is {
		if p(i) {
			out = append(out, i)
		}
	}

	return out
}
