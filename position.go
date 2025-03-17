package main

import (
	"fmt"

	"github.com/fj68/hyper-tux-go/hyper"
)

type Position struct {
	X, Y float64
}

func NewPositionFromPoint(p hyper.Point) Position {
	return Position{float64(p.X), float64(p.Y)}
}

func (p Position) String() string {
	return fmt.Sprintf("(%f, %f)", p.X, p.Y)
}
