package main

import "github.com/hajimehoshi/ebiten/v2"

type State interface {
	Update(*Game) error
	Draw(g *Game, screen *ebiten.Image)
}

type StateMachine struct {
	Current State
}

func (s *StateMachine) Update(g *Game) error {
	return s.Current.Update(g)
}

func (s *StateMachine) Draw(g *Game, screen *ebiten.Image) {
	s.Current.Draw(g, screen)
}
