package main

import "github.com/fj68/hyper-tux-go/hyper"

type ActionRecord struct {
	hyper.Color
	Start, End hyper.Point
}

type Action interface {
	Perform(*Game) *ActionRecord
}

type MoveAction struct {
	hyper.Actor
	hyper.Direction
}

func (a *MoveAction) Perform(g *Game) *ActionRecord {
	pos, ok, finished := g.MoveActor(a.Actor, a.Direction)
	if !ok {
		return nil
	}
	if finished {
		// goal!
	}
	return &ActionRecord{
		Color: a.Actor.Color,
		Start: a.Actor.Point,
		End: pos,
	}
}
