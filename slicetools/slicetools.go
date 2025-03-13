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

func Sum[T constraints.Integer | constraints.Float](xs []T) T {
	var total T
	for _, x := range xs {
		total += x
	}
	return total
}
