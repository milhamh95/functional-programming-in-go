package a_higher_order_functions

type Afn func() string

func A() string {
	return "hello"
}

func B(a Afn) string {
	return a() + " world"
}
