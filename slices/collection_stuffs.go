package main

func Find[A any](items []A, predicate func(A) bool) (val A, found bool) {
	for _, v := range items {
		if predicate(v) {
			return v, true
		}
	}
	return //returns (nil, 0) I suppose?
}

func Reduce[A, B any](collection []A, f func(B, A) B, initialValue B) B {
	var result = initialValue
	for _, x := range collection {
		result = f(result, x)
	}
	return result
}
