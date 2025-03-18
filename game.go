package main

import (
	"fmt"

	"github.com/fj68/hyper-tux-go/hyper"
	"github.com/hajimehoshi/ebiten/v2"
)

type GameState struct {
	*hyper.Board
	*SwipeEventDispatcher
	Actions []Action
	Records []*ActionRecord
}

func NewGameState(size hyper.Size) (*GameState, error) {
	board, err := hyper.NewBoard(size)
	if err != nil {
		return nil, err
	}
	return &GameState{
		Board:                board,
		SwipeEventDispatcher: NewSwipeEventDispather(),
	}, nil
}

func (g *GameState) handleInput() error {
	if err := g.SwipeEventDispatcher.Update(); err != nil {
		return err
	}

	for g.SwipeEventDispatcher.Len() > 0 {
		e := g.SwipeEventDispatcher.Pop()
		if e == nil {
			return fmt.Errorf("")
		}
		actor, ok := g.Board.ActorAt(e.Start)
		if !ok {
			return fmt.Errorf("")
		}
		action := MoveAction{actor, e.Direction()}
		if r := action.Perform(g); r != nil {
			g.Records = append(g.Records, r)
		}
	}

	return nil
}

func (g *GameState) Update() error {
	if err := g.handleInput(); err != nil {
		return err
	}

	return nil
}

func (g *GameState) Draw(screen *ebiten.Image) {}
