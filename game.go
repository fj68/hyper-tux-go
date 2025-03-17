package main

import (
	"fmt"

	"github.com/fj68/hyper-tux-go/hyper"
)

type Game struct {
	*hyper.Board
	*SwipeEventDispatcher
	Actions []Action
	Records []*ActionRecord
}

func (g *Game) handleInput() error {
	if err := g.SwipeEventDispatcher.Update(); err != nil {
		return err
	}

	for g.SwipeEventDispatcher.Len() > 0 {
		e, ok := g.SwipeEventDispatcher.Pop()
		if !ok {
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

func (g *Game) Update() error {
	if err := g.handleInput(); err != nil {
		return err
	}

	return nil
}
