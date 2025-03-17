package hyper

import (
	"math"
)

type Size struct {
	W, H int
}

func (s *Size) Center() Point {
	return Point{
		(int)(math.Floor(float64(s.W) / 2)),
		(int)(math.Floor(float64(s.H) / 2)),
	}
}

func (s *Size) Equals(other Size) bool {
	return s.W == other.W && s.H == other.H
}
