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
		g.Actions = append(g.Actions, &MoveAction{actor, e.Direction()})
	}

	return nil
}

func (g *Game) handleActions() error {
	for _, action := range g.Actions {
		r := action.Perform(g)
		if r != nil {
			g.Records = append(g.Records, r)
		}
	}
}

func (g *Game) Update() error {
	if err := g.handleInput(); err != nil {
		return err
	}

	if err := g.handleActions(); err != nil {
		return err
	}

	return nil
}
