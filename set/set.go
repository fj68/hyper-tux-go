package set

import (
	"iter"
	"maps"
	"slices"
)

type Set[T comparable] map[T]struct{}

func (s Set[T]) Add(v T) {
	((map[T]struct{})(s))[v] = struct{}{}
}

func (s Set[T]) Remove(v T) {
	delete(s, v)
}

func (s Set[T]) Contains(v T) bool {
	_, ok := ((map[T]struct{})(s))[v]
	return ok
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) Values() iter.Seq[T] {
	return maps.Keys(s)
}

func (s Set[T]) Collect() []T {
	return slices.Collect(maps.Keys(s))
}
