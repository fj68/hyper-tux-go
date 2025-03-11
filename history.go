package main

import (
	"math"

	"github.com/fj68/hyper-tux-go/hyper"
)

type Record struct {
	Actor         hyper.Color
	Before, After hyper.Point
	hyper.Direction
}

type History[T any] struct {
	values []T
	index  int
}

func (h *History[T]) Value() T {
	if len(h.values) < 1 {
		var zero T
		return zero
	}
	return h.values[h.index]
}

func (h *History[T]) Push(value T) {
	if h.index == len(h.values) {
		h.values = append(h.values, value)
		h.index++
		return
	}
	h.values = append(h.values[:h.index], value)
	h.index++
}

func (h *History[T]) Undo() {
	h.index = int(math.Max(float64(h.index-1), 0))
}

func (h *History[T]) Redo() {
	h.index = int(math.Min(float64(h.index+1), float64(len(h.values))))
}
