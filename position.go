package main

import (
	"fmt"

	"github.com/fj68/hyper-tux-go/hyper"
)

// Position represents a position in screen space as floating-point coordinates.
type Position struct {
	X, Y float32
}

// NewPosition converts a grid point to screen coordinates given a cell size.
func NewPosition(p hyper.Point, cellSize float32) Position {
	return Position{float32(p.X) * cellSize, float32(p.Y) * cellSize}
}

// ToPoint converts screen coordinates to a grid point given a cell size.
func (p *Position) ToPoint(cellSize float32) hyper.Point {
	return hyper.Point{
		X: int(p.X / cellSize),
		Y: int(p.Y / cellSize),
	}
}

// Add returns the sum of this position and another position.
func (p *Position) Add(other Position) Position {
	return Position{
		X: p.X + other.X,
		Y: p.Y + other.Y,
	}
}

// Sub returns the difference of this position and another position.
func (p *Position) Sub(other Position) Position {
	return Position{
		X: p.X - other.X,
		Y: p.Y - other.Y,
	}
}

// Mul returns the component-wise product of this position and another position.
func (p *Position) Mul(other Position) Position {
	return Position{
		X: p.X * other.X,
		Y: p.Y * other.Y,
	}
}

// Equals returns true if this position and another position have the same coordinates.
func (p *Position) Equals(other Position) bool {
	return p.X == other.X && p.Y == other.Y
}

// String returns a string representation of the position.
func (p *Position) String() string {
	return fmt.Sprintf("(%f, %f)", p.X, p.Y)
}
