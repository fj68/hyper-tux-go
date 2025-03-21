//go:build guitests
// +build guitests

package main_test

import (
	"image/color"
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

	g := &main.StateMachine{Current: s}

	image := ebiten.NewImage(size.W*int(CELL_SIZE), size.H*int(CELL_SIZE))
	image.Fill(color.Black)

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
