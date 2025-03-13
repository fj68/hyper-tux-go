package hyper_test

import (
	"testing"

	"github.com/fj68/hyper-tux-go/hyper"
)

func TestMoveActor(t *testing.T) {
	board, err := hyper.NewBoard(hyper.Size{16, 16})
	if err != nil {
		t.Fatal(err)
	}
	board.NewGame()

	type result struct {
		Ok       bool
		Finished bool
	}

	testcases := []struct {
		Name      string
		Goal      hyper.Goal
		Actor     hyper.Actor
		Direction hyper.Direction
		Walls     func(b *hyper.Board)
		Expected  result
	}{
		{
			"unable to move to north",
			hyper.Goal{hyper.Red, hyper.Point{3, 3}},
			hyper.Actor{hyper.Red, hyper.Point{5, 5}},
			hyper.North,
			func(b *hyper.Board) {
				b.PutHWall(hyper.Point{5, 5})
			},
			result{false, false},
		},
		{
			"unable to move to west",
			hyper.Goal{hyper.Red, hyper.Point{3, 3}},
			hyper.Actor{hyper.Red, hyper.Point{5, 5}},
			hyper.West,
			func(b *hyper.Board) {
				b.PutVWall(hyper.Point{5, 5})
			},
			result{false, false},
		},
		{
			"unable to move to south",
			hyper.Goal{hyper.Red, hyper.Point{3, 3}},
			hyper.Actor{hyper.Red, hyper.Point{5, 5}},
			hyper.South,
			func(b *hyper.Board) {
				b.PutVWall(hyper.Point{5, 5})
			},
			result{false, false},
		},
		{
			"unable to move to east",
			hyper.Goal{hyper.Red, hyper.Point{3, 3}},
			hyper.Actor{hyper.Red, hyper.Point{5, 5}},
			hyper.East,
			func(b *hyper.Board) {
				b.PutVWall(hyper.Point{5, 5})
			},
			result{false, false},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			board.Goal = testcase.Goal

			i, _ := board.Actor(testcase.Actor.Color)
			board.Actors[i].Point = testcase.Actor.Point

			testcase.Walls(board)

			ok, finished := board.MoveActor(testcase.Actor.Color, testcase.Direction)

			if testcase.Expected.Ok != ok {
				t.Error("unexpected return value: ok")
			}
			if testcase.Expected.Finished != finished {
				t.Error("unexpected return value: finished")
			}
		})
	}
}
