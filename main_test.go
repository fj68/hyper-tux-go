//go:build guitests
// +build guitests

package main_test

import (
	"testing"

	"github.com/fj68/hyper-tux-go"
	"github.com/fj68/hyper-tux-go/hyper"
	"github.com/fj68/hyper-tux-go/internal/snapshot_test"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	CELL_SIZE = main.CELL_SIZE
)

func TestGameState(t *testing.T) {
	size := hyper.Size{W: 16, H: 16}
	s, err := main.NewGameState(size)
	if err != nil {
		panic(err)
	}

	m, err := hyper.NewMapdataFromSlice([][]int{
		{0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0},
	})
	if err != nil {
		panic(err)
	}
	s.Board.Mapdata = m
	s.Board.Actors[hyper.Red].Point = hyper.Point{X: 1, Y: 0}
	s.Board.Actors[hyper.Blue].Point = hyper.Point{X: 1, Y: 3}
	s.Board.Actors[hyper.Yellow].Point = hyper.Point{X: 5, Y: 2}

	g := &main.StateMachine{Current: s}

	s.SwipeEventDispatcher.Push(&main.SwipeEvent{
		Start: hyper.Point{X: 1, Y: 0},
		End:   hyper.Point{X: 1, Y: 2},
	})
	s.SwipeEventDispatcher.Push(&main.SwipeEvent{
		Start: hyper.Point{X: 1, Y: 2},
		End:   hyper.Point{X: 2, Y: 2},
	})
	s.SwipeEventDispatcher.Push(&main.SwipeEvent{
		Start: hyper.Point{X: 1, Y: 3},
		End:   hyper.Point{X: 1, Y: 0},
	})
	s.SwipeEventDispatcher.Push(&main.SwipeEvent{
		Start: hyper.Point{X: 4, Y: 2},
		End:   hyper.Point{X: 4, Y: 0},
	})
	s.SwipeEventDispatcher.Push(&main.SwipeEvent{
		Start: hyper.Point{X: 5, Y: 2},
		End:   hyper.Point{X: 3, Y: 2},
	})

	image := ebiten.NewImage(main.SCREEN_WIDTH, main.SCREEN_HEIGHT)

	if err := g.Update(); err != nil {
		t.Error(err)
	}
	g.Draw(image)

	if err := snapshot_test.CheckSnapshot(t, image); err != nil {
		t.Error(err)
	}
}

func TestMain(m *testing.M) {
	snapshot_test.RunTestGame(m)
}
