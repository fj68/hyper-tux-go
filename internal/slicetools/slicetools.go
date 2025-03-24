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

func Mapi[T, U any](xs []T, f func(int) U) []U {
	rs := make([]U, len(xs))
	for i := range xs {
		rs[i] = f(i)
	}
	return rs
}

func Filter[T any](xs []T, f func(T) bool) []T {
	rs := make([]T, 0, len(xs))
	for _, x := range xs {
		if f(x) {
			rs = append(rs, x)
		}
	}
	return rs
}

func FilterMap[T, U any](xs []T, f func(T) bool, m func(T) U) []U {
	rs := make([]U, 0, len(xs))
	for _, x := range xs {
		if f(x) {
			rs = append(rs, m(x))
		}
	}
	return rs
}

func Flat[T any](xs [][]T) []T {
	if len(xs) < 1 {
		return nil
	}
	rs := make([]T, 0, len(xs)*len(xs[0]))
	for _, x := range xs {
		rs = append(rs, x...)
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
