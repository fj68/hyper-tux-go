// Package set provides a generic set data structure implementation.
// It offers O(1) lookups, insertions, and deletions using an underlying map.
package set

import (
	"iter"
	"maps"
	"slices"
)

// Set is a generic set implementation using a map for O(1) lookups.
type Set[T comparable] map[T]struct{}

// New creates and returns an empty Set.
func New[T comparable]() Set[T] {
	return (Set[T])(map[T]struct{}{})
}

// Add inserts a value into the set.
func (s Set[T]) Add(v T) {
	((map[T]struct{})(s))[v] = struct{}{}
}

// Remove deletes a value from the set.
func (s Set[T]) Remove(v T) {
	delete(s, v)
}

// Contains returns true if the value is in the set.
func (s Set[T]) Contains(v T) bool {
	_, ok := ((map[T]struct{})(s))[v]
	return ok
}

// Len returns the number of elements in the set.
func (s Set[T]) Len() int {
	return len(s)
}

// Values returns an iterator over the values in the set.
func (s Set[T]) Values() iter.Seq[T] {
	return maps.Keys(s)
}

// Collect returns a slice containing all elements from the set.
func (s Set[T]) Collect() []T {
	return slices.Collect(maps.Keys(s))
}

// Equals returns true if both sets contain the same elements.
func (s Set[T]) Equals(other Set[T]) bool {
	if s.Len() != other.Len() {
		return false
	}
	for a := range s {
		if !other.Contains(a) {
			return false
		}
	}
	return true
}
