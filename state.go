package main

import "github.com/hajimehoshi/ebiten/v2"

// State is an interface for game states that can be updated and drawn.
type State interface {
	Update() error
	Draw(screen *ebiten.Image)
}

// StateMachine manages state transitions and delegates Update and Draw calls to the current state.
type StateMachine struct {
	Current State
}

// Update delegates the update call to the current state.
func (s *StateMachine) Update() error {
	return s.Current.Update()
}

// Draw delegates the draw call to the current state.
func (s *StateMachine) Draw(screen *ebiten.Image) {
	s.Current.Draw(screen)
}
