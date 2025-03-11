package hyper

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func (p Point) Equals(other Point) bool {
	return (p.X == other.X && p.Y == other.Y)
}

type Size struct {
	W, H int
}

func (s Size) String() string {
	return fmt.Sprintf("Size{%d, %d}", s.W, s.H)
}

func (s Size) Center() Point {
	return Point{
		(int)(math.Floor(float64(s.W) / 2)),
		(int)(math.Floor(float64(s.H) / 2)),
	}
}

type Rect struct {
	TopLeft, BottomRight Point
}

func NewRect(topLeft Point, size Size) Rect {
	bottomRight := Point{
		topLeft.X + size.W,
		topLeft.Y + size.H,
	}
	return Rect{topLeft, bottomRight}
}

func (r Rect) Size() Size {
	return Size{
		r.BottomRight.X - r.TopLeft.X,
		r.BottomRight.Y - r.TopLeft.Y,
	}
}

func (r Rect) Contains(p Point) bool {
	return (r.TopLeft.X <= p.X && r.BottomRight.X < p.X &&
		r.TopLeft.Y <= p.Y && r.BottomRight.Y < p.Y)
}

type Color int

const (
	Red Color = iota
	Green
	Blue
	Yellow
	Black
)

func (c Color) String() string {
	switch c {
	case Red:
		return "Red"
	case Green:
		return "Green"
	case Blue:
		return "Blue"
	case Yellow:
		return "Yellow"
	case Black:
		return "Black"
	}
	return "unknown Color"
}

var AllColors = []Color{
	Red,
	Green,
	Blue,
	Yellow,
	Black,
}

type Direction int

const (
	North Direction = iota
	West
	East
	South
)

func (d Direction) String() string {
	switch d {
	case North:
		return "North"
	case West:
		return "West"
	case East:
		return "East"
	case South:
		return "South"
	}
	return "unknown Direction"
}
