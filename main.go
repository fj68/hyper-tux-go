package main

import (
	"github.com/fj68/hyper-tux-go/hyper"
	"github.com/hajimehoshi/ebiten/v2"
)

const CONTROLS_HEIGHT = 64 + 16            // px incl. padding
const STAGE_WIDTH = 640 - CONTROLS_HEIGHT  // px
const STAGE_HEIGHT = 640 - CONTROLS_HEIGHT // px
const CELL_SIZE float32 = STAGE_WIDTH / 16

type Game struct {
	State
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 640
}

func main() {
	s, err := NewGameState(hyper.Size{W: 16, H: 16})
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

	game := &Game{&StateMachine{Current: s}}

	ebiten.SetWindowSize(640, 640)
	ebiten.SetWindowTitle("Hyper Tux")

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
