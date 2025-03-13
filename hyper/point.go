package hyper

import (
	"fmt"
)

type Point struct {
	X, Y int
}

func (p *Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func (p *Point) Equals(other *Point) bool {
	return (p.X == other.X && p.Y == other.Y)
}
