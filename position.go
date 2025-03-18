package main

import (
	"fmt"

	"github.com/fj68/hyper-tux-go/hyper"
)

type Position struct {
	X, Y float32
}

func NewPosition(p hyper.Point, cellSize float32) Position {
	return Position{float32(p.X) * cellSize, float32(p.Y) * cellSize}
}

func (p *Position) ToPoint(cellSize float32) hyper.Point {
	return hyper.Point{
		X: int(p.X / cellSize),
		Y: int(p.Y / cellSize),
	}
}

func (p *Position) Add(other Position) Position {
	return Position{
		X: p.X + other.X,
		Y: p.Y + other.Y,
	}
}

func (p *Position) Sub(other Position) Position {
	return Position{
		X: p.X - other.X,
		Y: p.Y - other.Y,
	}
}

func (p *Position) Mul(other Position) Position {
	return Position{
		X: p.X * other.X,
		Y: p.Y * other.Y,
	}
}

func (p *Position) Equals(other Position) bool {
	return p.X == other.X && p.Y == other.Y
}

func (p *Position) String() string {
	return fmt.Sprintf("(%f, %f)", p.X, p.Y)
}
