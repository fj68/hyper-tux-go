package hyper

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y int
}

func (p *Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func (p *Point) Equals(other Point) bool {
	return (p.X == other.X && p.Y == other.Y)
}

func (p *Point) Add(other Point) Point {
	return Point{X: p.X + other.X, Y: p.Y + other.Y}
}

func (p *Point) Sub(other Point) Point {
	return Point{X: p.X - other.X, Y: p.Y - other.Y}
}

func (p *Point) Abs() Point {
	return Point{
		X: int(math.Abs(float64(p.X))),
		Y: int(math.Abs(float64(p.Y))),
	}
}
