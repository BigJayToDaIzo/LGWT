package main

type Person struct {
	Name string
}

func Find[A comparable](collection []A, reducer func(A) bool) (A, bool) {
	for _, x := range collection {
		if reducer(x) {
			return x, true
		}
	}
	var zero A
	return zero, false
}
