package slicetools

func Every[T any](xs []T, f func(T) bool) bool {
	for _, v := range xs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Some[T any](xs []T, f func(T) bool) bool {
	for _, v := range xs {
		if f(v) {
			return true
		}
	}
	return false
}
