package hyper

import (
	"fmt"
	"math"
)

// Point represents a point in game board grid coordinates.
type Point struct {
	X, Y int
}

// String returns the string representation of the point.
func (p *Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

// Equals returns true if this point and another point have the same coordinates.
func (p *Point) Equals(other Point) bool {
	return (p.X == other.X && p.Y == other.Y)
}

// Add returns the sum of this point and another point.
func (p *Point) Add(other Point) Point {
	return Point{X: p.X + other.X, Y: p.Y + other.Y}
}

// Sub returns the difference of this point and another point.
func (p *Point) Sub(other Point) Point {
	return Point{X: p.X - other.X, Y: p.Y - other.Y}
}

// Abs returns the absolute (component-wise) value of the point.
func (p *Point) Abs() Point {
	return Point{
		X: int(math.Abs(float64(p.X))),
		Y: int(math.Abs(float64(p.Y))),
	}
}
