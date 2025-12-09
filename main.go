// Package main provides the entrypoint for the Hyper Tux game.
// It initializes the game state, sets up the window and runs the game loop.
package main

import (
	"github.com/fj68/hyper-tux-go/hyper"
	"github.com/hajimehoshi/ebiten/v2"
)

// CONTROLS_HEIGHT is the height of the controls panel in pixels including padding.
const CONTROLS_HEIGHT = 64 + 16 // px incl. padding
// STAGE_WIDTH is the width of the game board in pixels.
const STAGE_WIDTH = 640 - CONTROLS_HEIGHT // px
// STAGE_HEIGHT is the height of the game board in pixels.
const STAGE_HEIGHT = 640 - CONTROLS_HEIGHT // px
// CELL_SIZE is the width and height of each cell in the grid in pixels.
const CELL_SIZE float32 = STAGE_WIDTH / 16

// Game is the main game struct that implements ebiten.Game interface.
type Game struct {
	State
}

// Layout returns the logical screen dimensions for the game window.
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
