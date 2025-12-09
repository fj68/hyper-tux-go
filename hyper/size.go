package hyper

import (
	"math"
)

// Size represents the dimensions of the game board.
type Size struct {
	W, H int
}

// Center returns the center point of a region of this size.
func (s *Size) Center() Point {
	return Point{
		(int)(math.Floor(float64(s.W) / 2)),
		(int)(math.Floor(float64(s.H) / 2)),
	}
}

// Equals returns true if both sizes have the same width and height.
func (s *Size) Equals(other Size) bool {
	return s.W == other.W && s.H == other.H
}
