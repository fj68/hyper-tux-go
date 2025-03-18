package hyper_test

import (
	"testing"

	"github.com/fj68/hyper-tux-go/hyper"
)

func TestBoard_NextStop(t *testing.T) {
	size := hyper.Size{16, 16}

	testcases := []struct {
		Name string
		*hyper.Actor
		hyper.Direction
		Walls    func(b *hyper.Board)
		Expected hyper.Point
	}{
		{
			"move north",
			&hyper.Actor{hyper.Red, hyper.Point{5, 5}},
			hyper.North,
			func(b *hyper.Board) {
				b.PutHWall(hyper.Point{5, 3})
			},
			hyper.Point{5, 3},
		},
		{
			"move north and stop at the edge",
			&hyper.Actor{hyper.Red, hyper.Point{5, 5}},
			hyper.North,
			func(b *hyper.Board) {
				b.PutHWall(hyper.Point{5, 6})
			},
			hyper.Point{5, 0},
		},
		{
			"move north and stop by another actor",
			&hyper.Actor{hyper.Red, hyper.Point{5, 5}},
			hyper.North,
			func(b *hyper.Board) {
				b.PutHWall(hyper.Point{5, 2})
				b.Actors[hyper.Blue] = &hyper.Actor{hyper.Blue, hyper.Point{5, 3}}
			},
			hyper.Point{5, 3},
		},
		{
			"move south",
			&hyper.Actor{hyper.Red, hyper.Point{5, 5}},
			hyper.South,
			func(b *hyper.Board) {
				b.PutHWall(hyper.Point{5, 10})
			},
			hyper.Point{5, 10},
		},
		{
			"move south and stop at the edge",
			&hyper.Actor{hyper.Red, hyper.Point{5, 5}},
			hyper.South,
			func(b *hyper.Board) {
				b.PutHWall(hyper.Point{5, 5})
			},
			hyper.Point{5, size.H - 1},
		},
		{
			"move south and stop by another actor",
			&hyper.Actor{hyper.Red, hyper.Point{5, 5}},
			hyper.South,
			func(b *hyper.Board) {
				b.PutHWall(hyper.Point{5, 10})
				b.Actors[hyper.Blue] = &hyper.Actor{hyper.Blue, hyper.Point{5, 7}}
			},
			hyper.Point{5, 7},
		},
		{
			"move west",
			&hyper.Actor{hyper.Red, hyper.Point{5, 5}},
			hyper.West,
			func(b *hyper.Board) {
				b.PutVWall(hyper.Point{3, 5})
			},
			hyper.Point{3, 5},
		},
		{
			"move west and stop at the edge",
			&hyper.Actor{hyper.Red, hyper.Point{5, 5}},
			hyper.West,
			func(b *hyper.Board) {
				b.PutVWall(hyper.Point{6, 5})
			},
			hyper.Point{0, 5},
		},
		{
			"move west and stop by another actor",
			&hyper.Actor{hyper.Red, hyper.Point{5, 5}},
			hyper.West,
			func(b *hyper.Board) {
				b.PutVWall(hyper.Point{2, 5})
				b.Actors[hyper.Blue] = &hyper.Actor{hyper.Blue, hyper.Point{3, 5}}
			},
			hyper.Point{3, 5},
		},
		{
			"move east",
			&hyper.Actor{hyper.Red, hyper.Point{5, 5}},
			hyper.East,
			func(b *hyper.Board) {
				b.PutVWall(hyper.Point{7, 5})
			},
			hyper.Point{7, 5},
		},
		{
			"move east and stop at the edge",
			&hyper.Actor{hyper.Red, hyper.Point{5, 5}},
			hyper.East,
			func(b *hyper.Board) {
				b.PutVWall(hyper.Point{5, 5})
			},
			hyper.Point{size.W - 1, 5},
		},
		{
			"move east and stop by another actor",
			&hyper.Actor{hyper.Red, hyper.Point{5, 5}},
			hyper.East,
			func(b *hyper.Board) {
				b.PutVWall(hyper.Point{10, 5})
				b.Actors[hyper.Blue] = &hyper.Actor{hyper.Blue, hyper.Point{7, 5}}
			},
			hyper.Point{7, 5},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			board, err := hyper.NewBoard(size)
			if err != nil {
				t.Fatal(err)
			}
			board.NewGame()

			// delete all actors and add neccessary one only
			// because they are randomly placed and may block others accidentally
			// which leads tests failed
			board.Actors = map[hyper.Color]*hyper.Actor{}
			board.Actors[testcase.Actor.Color] = testcase.Actor

			actor, ok := board.Actors[testcase.Actor.Color]
			if !ok {
				t.Fatalf("unable to find actor of color: %s", testcase.Actor.Color)
			}
			actor.MoveTo(testcase.Point)

			testcase.Walls(board)

			actual := board.NextStop(testcase.Actor.Point, testcase.Direction)
			if !actual.Equals(testcase.Expected) {
				t.Errorf("unexpected value: Expected = %+v, Actual = %+v", testcase.Expected, actual)
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
		Actor     *hyper.Actor
		Direction hyper.Direction
		Walls     func(b *hyper.Board)
		Expected  result
	}{
		{
			"unable to move north",
			hyper.Goal{hyper.Red, hyper.Point{3, 3}},
			&hyper.Actor{hyper.Red, hyper.Point{5, 5}},
			hyper.North,
			func(b *hyper.Board) {
				b.PutHWall(hyper.Point{5, 5})
			},
			result{false, false, hyper.Point{5, 5}},
		},
		{
			"unable to move west",
			hyper.Goal{hyper.Red, hyper.Point{3, 3}},
			&hyper.Actor{hyper.Red, hyper.Point{5, 5}},
			hyper.West,
			func(b *hyper.Board) {
				b.PutVWall(hyper.Point{5, 5})
			},
			result{false, false, hyper.Point{5, 5}},
		},
		{
			"unable to move south",
			hyper.Goal{hyper.Red, hyper.Point{3, 3}},
			&hyper.Actor{hyper.Red, hyper.Point{5, 5}},
			hyper.South,
			func(b *hyper.Board) {
				b.PutHWall(hyper.Point{5, 5})
			},
			result{false, false, hyper.Point{5, 5}},
		},
		{
			"unable to move east",
			hyper.Goal{hyper.Red, hyper.Point{3, 3}},
			&hyper.Actor{hyper.Red, hyper.Point{5, 5}},
			hyper.East,
			func(b *hyper.Board) {
				b.PutVWall(hyper.Point{5, 5})
			},
			result{false, false, hyper.Point{5, 5}},
		},
		{
			"reached goal",
			hyper.Goal{hyper.Red, hyper.Point{5, 5}},
			&hyper.Actor{hyper.Red, hyper.Point{5, 4}},
			hyper.South,
			func(b *hyper.Board) {
				b.PutHWall(hyper.Point{5, 5})
			},
			result{true, true, hyper.Point{5, 5}},
		},
		{
			"reached black goal",
			hyper.Goal{hyper.Black, hyper.Point{5, 5}},
			&hyper.Actor{hyper.Red, hyper.Point{4, 5}},
			hyper.East,
			func(b *hyper.Board) {
				b.PutVWall(hyper.Point{5, 5})
			},
			result{true, true, hyper.Point{5, 5}},
		},
		{
			"move south",
			hyper.Goal{hyper.Red, hyper.Point{5, 5}},
			&hyper.Actor{hyper.Red, hyper.Point{5, 4}},
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

			actor, ok := board.Actors[testcase.Actor.Color]
			if !ok {
				t.Fatalf("unable to find actor of color: %s", testcase.Actor.Color)
			}
			actor.MoveTo(testcase.Actor.Point)

			testcase.Walls(board)

			_, ok = board.MoveActor(testcase.Actor, testcase.Direction)

			if testcase.Expected.Ok != ok {
				t.Errorf("unexpected return value ok: Actor = %+v, Goal = %+v", actor, board.Goal)
			}
			if testcase.Expected.Finished != board.Goaled {
				t.Errorf("unexpected return value finished: Actor = %+v, Goal = %+v", actor, board.Goal)
			}
			if !testcase.Expected.Point.Equals(actor.Point) {
				t.Errorf("unexpected return value position: Expected = %+v, Actor = %+v, Goal = %+v", testcase.Expected.Point, actor, board.Goal)
			}
		})
	}
}
