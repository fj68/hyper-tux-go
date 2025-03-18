package main

import "github.com/hajimehoshi/ebiten/v2"

type State interface {
	Update() error
	Draw(screen *ebiten.Image)
}

type StateMachine struct {
	Current State
}

func (s *StateMachine) Update() error {
	return s.Current.Update()
}

func (s *StateMachine) Draw(screen *ebiten.Image) {
	s.Current.Draw(screen)
}
