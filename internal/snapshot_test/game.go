package snapshot_test

import (
	"errors"
	"os"
	"testing"

	"github.com/hajimehoshi/ebiten/v2"
)

var errRegularTermination = errors.New("regular termination")

type game struct {
	m    *testing.M
	code int
}

func (g *game) Update() error {
	g.code = g.m.Run()
	return errRegularTermination
}

func (*game) Draw(*ebiten.Image) {
}

func (g *game) Layout(int, int) (int, int) {
	return 1, 1
}

// RunTestGame runs the provided test suite in the Ebiten game window environment.
func RunTestGame(m *testing.M) {
	ebiten.SetWindowSize(128, 72)
	ebiten.SetWindowTitle("Testing...")

	g := &game{
		m: m,
	}
	if err := ebiten.RunGameWithOptions(g, &ebiten.RunGameOptions{
		InitUnfocused: true,
	}); err != nil && err != errRegularTermination {
		panic(err)
	}
	os.Exit(g.code)
}
