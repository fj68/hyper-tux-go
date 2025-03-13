package hyper_test

import (
	"testing"

	"github.com/fj68/hyper-tux-go/hyper"
)

func TestBoard_NextStop(t *testing.T) {
	size := hyper.Size{16, 16}

	testcases := []struct {
		Name string
		hyper.Actor
		hyper.Direction
		Walls    func(b *hyper.Board)
		Expected hyper.Point
	}{
		{
			"move north",
			hyper.Actor{hyper.Red, hyper.Point{5, 5}},
			hyper.North,
			func(b *hyper.Board) {
				b.PutHWall(hyper.Point{5, 3})
			},
			hyper.Point{5, 3},
		},
		{
			"move north and stop at the edge",
			hyper.Actor{hyper.Red, hyper.Point{5, 5}},
			hyper.North,
			func(b *hyper.Board) {
				b.PutHWall(hyper.Point{5, 6})
			},
			hyper.Point{5, 0},
		},
		{
			"move south",
			hyper.Actor{hyper.Red, hyper.Point{5, 5}},
			hyper.South,
			func(b *hyper.Board) {
				b.PutHWall(hyper.Point{5, 10})
			},
			hyper.Point{5, 10},
		},
		{
			"move south and stop at the edge",
			hyper.Actor{hyper.Red, hyper.Point{5, 5}},
			hyper.South,
			func(b *hyper.Board) {
				b.PutHWall(hyper.Point{5, 5})
			},
			hyper.Point{5, size.H - 1},
		},
		{
			"move west",
			hyper.Actor{hyper.Red, hyper.Point{5, 5}},
			hyper.West,
			func(b *hyper.Board) {
				b.PutVWall(hyper.Point{3, 5})
			},
			hyper.Point{3, 5},
		},
		{
			"move west and stop at the edge",
			hyper.Actor{hyper.Red, hyper.Point{5, 5}},
			hyper.West,
			func(b *hyper.Board) {
				b.PutVWall(hyper.Point{6, 5})
			},
			hyper.Point{0, 5},
		},
		{
			"move east",
			hyper.Actor{hyper.Red, hyper.Point{5, 5}},
			hyper.East,
			func(b *hyper.Board) {
				b.PutVWall(hyper.Point{7, 5})
			},
			hyper.Point{7, 5},
		},
		{
			"move east and stop at the edge",
			hyper.Actor{hyper.Red, hyper.Point{5, 5}},
			hyper.East,
			func(b *hyper.Board) {
				b.PutVWall(hyper.Point{5, 5})
			},
			hyper.Point{size.W - 1, 5},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			board, err := hyper.NewBoard(size)
			if err != nil {
				t.Fatal(err)
			}
			board.NewGame()

			i, _ := board.Actor(testcase.Actor.Color)
			board.Actors[i].Point = testcase.Actor.Point

			testcase.Walls(board)

			actual := board.NextStop(testcase.Actor.Point, testcase.Direction)
			if actual != testcase.Expected {
				t.Errorf("unexpected value: Expected = %s, Actual = %s", testcase.Expected, actual)
			}
		})
	}
}

func TestBoard_MoveActor(t *testing.T) {
	type result struct {
		Ok       bool
		Finished bool
		hyper.Point
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
			"unable to move north",
			hyper.Goal{hyper.Red, hyper.Point{3, 3}},
			hyper.Actor{hyper.Red, hyper.Point{5, 5}},
			hyper.North,
			func(b *hyper.Board) {
				b.PutHWall(hyper.Point{5, 5})
			},
			result{false, false, hyper.Point{5, 5}},
		},
		{
			"unable to move west",
			hyper.Goal{hyper.Red, hyper.Point{3, 3}},
			hyper.Actor{hyper.Red, hyper.Point{5, 5}},
			hyper.West,
			func(b *hyper.Board) {
				b.PutVWall(hyper.Point{5, 5})
			},
			result{false, false, hyper.Point{5, 5}},
		},
		{
			"unable to move south",
			hyper.Goal{hyper.Red, hyper.Point{3, 3}},
			hyper.Actor{hyper.Red, hyper.Point{5, 5}},
			hyper.South,
			func(b *hyper.Board) {
				b.PutHWall(hyper.Point{5, 5})
			},
			result{false, false, hyper.Point{5, 5}},
		},
		{
			"unable to move east",
			hyper.Goal{hyper.Red, hyper.Point{3, 3}},
			hyper.Actor{hyper.Red, hyper.Point{5, 5}},
			hyper.East,
			func(b *hyper.Board) {
				b.PutVWall(hyper.Point{5, 5})
			},
			result{false, false, hyper.Point{5, 5}},
		},
		{
			"reached goal",
			hyper.Goal{hyper.Red, hyper.Point{5, 5}},
			hyper.Actor{hyper.Red, hyper.Point{5, 4}},
			hyper.South,
			func(b *hyper.Board) {
				b.PutHWall(hyper.Point{5, 5})
			},
			result{true, true, hyper.Point{5, 5}},
		},
		{
			"reached black goal",
			hyper.Goal{hyper.Black, hyper.Point{5, 5}},
			hyper.Actor{hyper.Red, hyper.Point{4, 5}},
			hyper.East,
			func(b *hyper.Board) {
				b.PutVWall(hyper.Point{5, 5})
			},
			result{true, true, hyper.Point{5, 5}},
		},
		{
			"move south",
			hyper.Goal{hyper.Red, hyper.Point{5, 5}},
			hyper.Actor{hyper.Red, hyper.Point{5, 4}},
			hyper.South,
			func(b *hyper.Board) {
				b.PutHWall(hyper.Point{5, 10})
			},
			result{true, false, hyper.Point{5, 10}},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			board, err := hyper.NewBoard(hyper.Size{16, 16})
			if err != nil {
				t.Fatal(err)
			}
			board.NewGame()

			board.Goal = testcase.Goal

			i, _ := board.Actor(testcase.Actor.Color)
			board.Actors[i].Point = testcase.Actor.Point

			testcase.Walls(board)

			ok, finished := board.MoveActor(testcase.Actor.Color, testcase.Direction)

			if testcase.Expected.Ok != ok {
				t.Errorf("unexpected return value ok: Actor = %s, Goal = %s", board.Actors[i], board.Goal)
			}
			if testcase.Expected.Finished != finished {
				t.Errorf("unexpected return value finished: Actor = %s, Goal = %s", board.Actors[i], board.Goal)
			}
			if !testcase.Expected.Point.Equals(board.Actors[i].Point) {
				t.Errorf("unexpected return value position: Expected = %s, Actor = %s, Goal = %s", testcase.Expected.Point, board.Actors[i], board.Goal)
			}
		})
	}
}
