package itertools

import (
	"iter"
)

func Filter[T any](it iter.Seq[T], f func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range it {
			if !f(v) {
				continue
			}
			if !yield(v) {
				break
			}
		}
	}
}

func Map[T, U any](it iter.Seq[T], f func(T) U) iter.Seq[U] {
	return func(yield func(U) bool) {
		for v := range it {
			if !yield(f(v)) {
				break
			}
		}
	}
}

func FilterMap[T, U any](it iter.Seq[T], ff func(T) bool, mf func(T) U) iter.Seq[U] {
	return func(yield func(U) bool) {
		for v := range it {
			if !ff(v) {
				continue
			}
			if !yield(mf(v)) {
				break
			}
		}
	}
}

func GroupBy[T any, K comparable](it iter.Seq[T], f func(T) K) map[K]T {
	m := map[K]T{}

	for v := range it {
		k := f(v)
		m[k] = v
	}

	return m
}

type Pair[L, R any] struct {
	Left  L
	Right R
}

func Seq2ToPairs[K comparable, V any](it iter.Seq2[K, V]) iter.Seq[Pair[K, V]] {
	return func(yield func(Pair[K, V]) bool) {
		for k, v := range it {
			if !yield(Pair[K, V]{Left: k, Right: v}) {
				break
			}
		}
	}
}

func PairsToSeq2[K comparable, V any](it iter.Seq[Pair[K, V]]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for pair := range it {
			if !yield(pair.Left, pair.Right) {
				break
			}
		}
	}
}

func Equals[T comparable](xs, ys iter.Seq[T]) bool {
	return EqualsFunc(xs, ys, func(x, y T) bool { return x == y })
}

func EqualsFunc[T any](xs, ys iter.Seq[T], f func(T, T) bool) bool {
	xNext, xStop := iter.Pull(xs)
	defer xStop()
	yNext, yStop := iter.Pull(ys)
	defer yStop()

	for {
		x, xOk := xNext()
		y, yOk := yNext()
		if !xOk || !yOk {
			return xOk == yOk
		}
		if !f(x, y) {
			return false
		}
	}
}
