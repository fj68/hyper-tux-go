package slicetools

import (
	"golang.org/x/exp/constraints"
)

func Every[T any](xs []T, f func(int, T) bool) bool {
	for i, v := range xs {
		if !f(i, v) {
			return false
		}
	}
	return true
}

func Some[T any](xs []T, f func(int, T) bool) bool {
	for i, v := range xs {
		if f(i, v) {
			return true
		}
	}
	return false
}

func Map[T, U any](xs []T, f func(T) U) []U {
	rs := make([]U, len(xs))
	for i, x := range xs {
		rs[i] = f(x)
	}
	return rs
}

func FoldLeft[T any, U any](xs []T, init U, f func(U, int, T) U) U {
	r := init
	for i, x := range xs {
		r = f(r, i, x)
	}
	return r
}

func Sum[T constraints.Integer | constraints.Float](xs []T) T {
	var acc T
	for _, x := range xs {
		acc += x
	}
	return acc
}

func Equals[T comparable](xs, ys []T) bool {
	return EqualsFunc(xs, ys, func(x, y T) bool { return x == y })
}

func EqualsFunc[T any](xs, ys []T, f func(T, T) bool) bool {
	if len(xs) != len(ys) {
		return false
	}
	for i := range xs {
		if !f(xs[i], ys[i]) {
			return false
		}
	}
	return true
}
