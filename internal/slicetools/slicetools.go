// Package slicetools provides generic functional programming utilities for working with slices.
// It includes functions for mapping, filtering, folding, and other common slice operations.
package slicetools

import (
	"golang.org/x/exp/constraints"
)

// Every returns true if the predicate f returns true for all elements in the slice.
func Every[T any](xs []T, f func(int, T) bool) bool {
	for i, v := range xs {
		if !f(i, v) {
			return false
		}
	}
	return true
}

// Some returns true if the predicate f returns true for at least one element in the slice.
func Some[T any](xs []T, f func(int, T) bool) bool {
	for i, v := range xs {
		if f(i, v) {
			return true
		}
	}
	return false
}

// Map applies function f to each element of the slice and returns a new slice with the results.
func Map[T, U any](xs []T, f func(T) U) []U {
	rs := make([]U, len(xs))
	for i, x := range xs {
		rs[i] = f(x)
	}
	return rs
}

// Mapi applies function f to the index of each element and returns a new slice with the results.
func Mapi[T, U any](xs []T, f func(int) U) []U {
	rs := make([]U, len(xs))
	for i := range xs {
		rs[i] = f(i)
	}
	return rs
}

// Filter returns a new slice containing only elements for which the predicate f returns true.
func Filter[T any](xs []T, f func(T) bool) []T {
	rs := make([]T, 0, len(xs))
	for _, x := range xs {
		if f(x) {
			rs = append(rs, x)
		}
	}
	return rs
}

// FilterMap filters elements using predicate f and maps matching elements using function m.
func FilterMap[T, U any](xs []T, f func(T) bool, m func(T) U) []U {
	rs := make([]U, 0, len(xs))
	for _, x := range xs {
		if f(x) {
			rs = append(rs, m(x))
		}
	}
	return rs
}

// Flat flattens a slice of slices into a single slice.
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

// FoldLeft reduces a slice to a single value by applying function f from left to right, starting with init.
func FoldLeft[T any, U any](xs []T, init U, f func(U, int, T) U) U {
	r := init
	for i, x := range xs {
		r = f(r, i, x)
	}
	return r
}

// Sum returns the sum of all elements in the slice.
func Sum[T constraints.Integer | constraints.Float](xs []T) T {
	var acc T
	for _, x := range xs {
		acc += x
	}
	return acc
}

// Equals returns true if both slices have the same elements in the same order.
func Equals[T comparable](xs, ys []T) bool {
	return EqualsFunc(xs, ys, func(x, y T) bool { return x == y })
}

// EqualsFunc returns true if both slices have the same length and elements satisfy the equality function f.
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
