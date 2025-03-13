package hyper

import (
	"fmt"
	"math"
)

type Size struct {
	W, H int
}

func (s *Size) String() string {
	return fmt.Sprintf("Size{%d, %d}", s.W, s.H)
}

func (s *Size) Center() Point {
	return Point{
		(int)(math.Floor(float64(s.W) / 2)),
		(int)(math.Floor(float64(s.H) / 2)),
	}
}
