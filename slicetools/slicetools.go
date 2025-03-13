package slicetools

import (
	"golang.org/x/exp/constraints"
)

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

func FoldLeft[T any, U any](xs []T, init U, f func(acc U, v T, i int) U) U {
	r := init
	for i, x := range xs {
		r = f(r, x, i)
	}
	return r
}

func Sum[T constraints.Integer | constraints.Float](xs []T) T {
	return FoldLeft(xs, 0, func(acc T, v T, i int) T {
		return acc + v
	})
}
