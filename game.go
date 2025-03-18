package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/ebitenui/ebitenui"
	"github.com/fj68/hyper-tux-go/hyper"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type GameState struct {
	*hyper.Board
	*SwipeEventDispatcher
	*ebitenui.UI
	labels *Labels
}

func NewGameState(size hyper.Size) (*GameState, error) {
	board, err := hyper.NewBoard(size)
	if err != nil {
		return nil, err
	}
	// debug
	board.Actors[hyper.Black].Point = hyper.Point{X: 0, Y: 0}
	for _, actor := range board.Actors {
		log.Println(actor)
	}

	ui, labels, err := createUI(board)
	if err != nil {
		return nil, err
	}

	return &GameState{
		Board: board,
		SwipeEventDispatcher: NewSwipeEventDispather(
			&MouseEventHandler{},
			&TouchEventHandler{},
		),
		UI:     ui,
		labels: labels,
	}, nil
}

func (g *GameState) handleInput() error {
	if err := g.SwipeEventDispatcher.Update(); err != nil {
		return err
	}

	for g.SwipeEventDispatcher.Len() > 0 {
		e := g.SwipeEventDispatcher.Pop()
		if e == nil {
			return fmt.Errorf("SwipeEvent is nil")
		}
		actor, ok := g.Board.ActorAt(e.Start)
		if !ok {
			// TODO: this should not be an error
			return fmt.Errorf("No actor at %+v", e.Start)
		}
		_, ok = g.Board.MoveActor(actor, e.Direction())
		if !ok {
			// TODO: this should not be an error
			return fmt.Errorf("Unable to move: %+v to %s", actor, e.Direction())
		}
	}

	return nil
}

func (g *GameState) Update() error {
	if err := g.handleInput(); err != nil {
		return err
	}

	g.labels.Update(g.Board)

	g.UI.Update()

	return nil
}

func (g *GameState) Draw(screen *ebiten.Image) {
	g.drawBoard(screen)
	g.drawActors(screen)
	g.drawGoal(screen)
	g.UI.Draw(screen)
}

func (g *GameState) drawBoard(screen *ebiten.Image) {
	lineColor := color.Gray{200}
	// background
	vector.DrawFilledRect(screen, 0, 0, float32(g.W)*CELL_SIZE, float32(g.H)*CELL_SIZE, color.White, false)
	// lines
	for y := range g.Board.H - 1 {
		vector.StrokeLine(screen, 0, float32(y+1)*CELL_SIZE, float32(g.Board.W)*CELL_SIZE, float32(y+1)*CELL_SIZE, 1, lineColor, false)
	}
	vector.StrokeLine(screen, 0, float32(g.Board.H)*CELL_SIZE, float32(g.Board.W)*CELL_SIZE, float32(g.Board.H)*CELL_SIZE, 1, color.Black, false)
	for x := range g.Board.W - 1 {
		vector.StrokeLine(screen, float32(x+1)*CELL_SIZE, 0, float32(x+1)*CELL_SIZE, float32(g.Board.H)*CELL_SIZE, 1, lineColor, false)
	}
	vector.StrokeLine(screen, float32(g.Board.W)*CELL_SIZE, 0, float32(g.Board.W)*CELL_SIZE, float32(g.Board.H)*CELL_SIZE, 1, color.Black, false)
	// walls
	for y, rows := range g.Board.VWalls {
		for _, x := range rows {
			vector.StrokeLine(screen, float32(x)*CELL_SIZE, float32(y)*CELL_SIZE, float32(x)*CELL_SIZE, float32(y+1)*CELL_SIZE, 1, color.Black, false)
		}
	}
	for x, cols := range g.Board.HWalls {
		for _, y := range cols {
			vector.StrokeLine(screen, float32(x)*CELL_SIZE, float32(y)*CELL_SIZE, float32(x+1)*CELL_SIZE, float32(y)*CELL_SIZE, 1, color.Black, false)
		}
	}
	// center box
	c := g.Board.Center()
	vector.DrawFilledRect(screen, float32(c.TopLeft.X)*CELL_SIZE, float32(c.TopLeft.Y)*CELL_SIZE, float32(c.Size().W)*CELL_SIZE-1, float32(c.Size().H)*CELL_SIZE-1, lineColor, false)
}

func (g *GameState) drawActors(screen *ebiten.Image) {
	for _, actor := range g.Board.Actors {
		g.drawActor(screen, actor)
	}
}

func (g *GameState) drawActor(screen *ebiten.Image, actor *hyper.Actor) {
	p := NewPosition(actor.Point, CELL_SIZE)
	halfCellSize := CELL_SIZE / 2
	p = p.Add(Position{halfCellSize, halfCellSize})
	r := halfCellSize - 2
	vector.DrawFilledCircle(screen, p.X, p.Y, r, Color(actor.Color), true)
}

func (g *GameState) drawGoal(screen *ebiten.Image) {
	goal := g.Board.Goal
	vector.DrawFilledRect(screen, float32(goal.X)*CELL_SIZE, float32(goal.Y)*CELL_SIZE, CELL_SIZE-1, CELL_SIZE-1, Color(goal.Color), false)
}
