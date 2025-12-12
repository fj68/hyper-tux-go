package hyper_test

import (
	"testing"

	"github.com/fj68/hyper-tux-go/hyper"
)

var defaultActorPlacement = map[hyper.Color]hyper.Point{
	hyper.Red:    {1, 1},
	hyper.Green:  {14, 1},
	hyper.Blue:   {1, 14},
	hyper.Yellow: {14, 14},
	hyper.Black:  {13, 14},
}

func TestBoard_NextStop(t *testing.T) {
	size := hyper.Size{16, 16}

	testcases := []struct {
		Name       string
		ActorColor hyper.Color
		Direction hyper.Direction
		Walls func(b *hyper.Board)
		Placement hyper.Placement
		Expected hyper.Point
	}{
		{
			"move north",
			hyper.Red,
			hyper.North,
			func(b *hyper.Board) {
				b.PutHWall(hyper.Point{5, 3})
			},
			hyper.Placement{
				Actor: hyper.PlaceActorAt(defaultActorPlacement, map[hyper.Color]hyper.Point{hyper.Red: {5, 5}}),
				Goal:  hyper.PlaceGoalAt(hyper.Black, hyper.Point{0, 0}),
			},
			hyper.Point{5, 3},
		},
		{
			"move north and stop at the edge",
			hyper.Red,
			hyper.North,
			func(b *hyper.Board) {
				b.PutHWall(hyper.Point{5, 6})
			},
			hyper.Placement{
				Actor: hyper.PlaceActorAt(defaultActorPlacement, map[hyper.Color]hyper.Point{hyper.Red: {5, 5}}),
				Goal:  hyper.PlaceGoalAt(hyper.Black, hyper.Point{0, 0}),
			},
			hyper.Point{5, 0},
		},
		{
			"move north and stop by another actor",
			hyper.Red,
			hyper.North,
			func(b *hyper.Board) {
				b.PutHWall(hyper.Point{5, 2})
				b.Actors[hyper.Blue] = &hyper.Actor{hyper.Blue, hyper.Point{5, 3}}
			},
			hyper.Placement{
				Actor: hyper.PlaceActorAt(defaultActorPlacement, map[hyper.Color]hyper.Point{hyper.Red: {5, 5}}),
				Goal:  hyper.PlaceGoalAt(hyper.Black, hyper.Point{0, 0}),
			},
			hyper.Point{5, 3},
		},
		{
			"move south",
			hyper.Red,
			hyper.South,
			func(b *hyper.Board) {
				b.PutHWall(hyper.Point{5, 10})
			},
			hyper.Placement{
				Actor: hyper.PlaceActorAt(defaultActorPlacement, map[hyper.Color]hyper.Point{hyper.Red: {5, 5}}),
				Goal:  hyper.PlaceGoalAt(hyper.Black, hyper.Point{0, 0}),
			},
			hyper.Point{5, 10},
		},
		{
			"move south and stop at the edge",
			hyper.Red,
			hyper.South,
			func(b *hyper.Board) {
				b.PutHWall(hyper.Point{5, 5})
			},
			hyper.Placement{
				Actor: hyper.PlaceActorAt(defaultActorPlacement, map[hyper.Color]hyper.Point{hyper.Red: {5, 5}}),
				Goal:  hyper.PlaceGoalAt(hyper.Black, hyper.Point{0, 0}),
			},
			hyper.Point{5, size.H - 1},
		},
		{
			"move south and stop by another actor",
			hyper.Red,
			hyper.South,
			func(b *hyper.Board) {
				b.PutHWall(hyper.Point{5, 10})
				b.Actors[hyper.Blue] = &hyper.Actor{hyper.Blue, hyper.Point{5, 7}}
			},
			hyper.Placement{
				Actor: hyper.PlaceActorAt(defaultActorPlacement, map[hyper.Color]hyper.Point{hyper.Red: {5, 5}}),
				Goal:  hyper.PlaceGoalAt(hyper.Black, hyper.Point{0, 0}),
			},
			hyper.Point{5, 7},
		},
		{
			"move west",
			hyper.Red,
			hyper.West,
			func(b *hyper.Board) {
				b.PutVWall(hyper.Point{3, 5})
			},
			hyper.Placement{
				Actor: hyper.PlaceActorAt(defaultActorPlacement, map[hyper.Color]hyper.Point{hyper.Red: {5, 5}}),
				Goal:  hyper.PlaceGoalAt(hyper.Black, hyper.Point{0, 0}),
			},
			hyper.Point{3, 5},
		},
		{
			"move west and stop at the edge",
			hyper.Red,
			hyper.West,
			func(b *hyper.Board) {
				b.PutVWall(hyper.Point{6, 5})
			},
			hyper.Placement{
				Actor: hyper.PlaceActorAt(defaultActorPlacement, map[hyper.Color]hyper.Point{hyper.Red: {5, 5}}),
				Goal:  hyper.PlaceGoalAt(hyper.Black, hyper.Point{0, 0}),
			},
			hyper.Point{0, 5},
		},
		{
			"move west and stop by another actor",
			hyper.Red,
			hyper.West,
			func(b *hyper.Board) {
				b.PutVWall(hyper.Point{2, 5})
				b.Actors[hyper.Blue] = &hyper.Actor{hyper.Blue, hyper.Point{3, 5}}
			},
			hyper.Placement{
				Actor: hyper.PlaceActorAt(defaultActorPlacement, map[hyper.Color]hyper.Point{hyper.Red: {5, 5}}),
				Goal:  hyper.PlaceGoalAt(hyper.Black, hyper.Point{0, 0}),
			},
			hyper.Point{3, 5},
		},
		{
			"move east",
			hyper.Red,
			hyper.East,
			func(b *hyper.Board) {
				b.PutVWall(hyper.Point{7, 5})
			},
			hyper.Placement{
				Actor: hyper.PlaceActorAt(defaultActorPlacement, map[hyper.Color]hyper.Point{hyper.Red: {5, 5}}),
				Goal:  hyper.PlaceGoalAt(hyper.Black, hyper.Point{0, 0}),
			},
			hyper.Point{7, 5},
		},
		{
			"move east and stop at the edge",
			hyper.Red,
			hyper.East,
			func(b *hyper.Board) {
				b.PutVWall(hyper.Point{5, 5})
			},
			hyper.Placement{
				Actor: hyper.PlaceActorAt(defaultActorPlacement, map[hyper.Color]hyper.Point{hyper.Red: {5, 5}}),
				Goal:  hyper.PlaceGoalAt(hyper.Black, hyper.Point{0, 0}),
			},
			hyper.Point{size.W - 1, 5},
		},
		{
			"move east and stop by another actor",
			hyper.Red,
			hyper.East,
			func(b *hyper.Board) {
				b.PutVWall(hyper.Point{10, 5})
				b.Actors[hyper.Blue] = &hyper.Actor{hyper.Blue, hyper.Point{7, 5}}
			},
			hyper.Placement{
				Actor: hyper.PlaceActorAt(defaultActorPlacement, map[hyper.Color]hyper.Point{hyper.Red: {5, 5}}),
				Goal:  hyper.PlaceGoalAt(hyper.Black, hyper.Point{0, 0}),
			},
			hyper.Point{7, 5},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			board, err := hyper.NewBoard(size, testcase.Placement)
			if err != nil {
				t.Fatal(err)
			}
			board.NewGame()

			testcase.Walls(board)

			actor, ok := board.Actors[testcase.ActorColor]
			if !ok {
				t.Fatalf("unable to find actor of color: %s", testcase.ActorColor)
			}

			actual := board.NextStop(actor.Point, testcase.Direction)
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
		ActorColor hyper.Color
		Direction hyper.Direction
		Walls     func(b *hyper.Board)
		Placement hyper.Placement
		Expected  result
	}{
		{
			"unable to move north",
			hyper.Red,
			hyper.North,
			func(b *hyper.Board) {
				b.PutHWall(hyper.Point{5, 5})
			},
			hyper.Placement{
				Actor: hyper.PlaceActorAt(defaultActorPlacement, map[hyper.Color]hyper.Point{hyper.Red: {5, 5}}),
				Goal:  hyper.PlaceGoalAt(hyper.Red, hyper.Point{3, 3}),
			},
			result{false, false, hyper.Point{5, 5}},
		},
		{
			"unable to move west",
			hyper.Red,
			hyper.West,
			func(b *hyper.Board) {
				b.PutVWall(hyper.Point{5, 5})
			},
			hyper.Placement{
				Actor: hyper.PlaceActorAt(defaultActorPlacement, map[hyper.Color]hyper.Point{hyper.Red: {5, 5}}),
				Goal:  hyper.PlaceGoalAt(hyper.Red, hyper.Point{3, 3}),
			},
			result{false, false, hyper.Point{5, 5}},
		},
		{
			"unable to move south",
			hyper.Red,
			hyper.South,
			func(b *hyper.Board) {
				b.PutHWall(hyper.Point{5, 5})
			},
			hyper.Placement{
				Actor: hyper.PlaceActorAt(defaultActorPlacement, map[hyper.Color]hyper.Point{hyper.Red: {5, 5}}),
				Goal:  hyper.PlaceGoalAt(hyper.Red, hyper.Point{3, 3}),
			},
			result{false, false, hyper.Point{5, 5}},
		},
		{
			"unable to move east",
			hyper.Red,
			hyper.East,
			func(b *hyper.Board) {
				b.PutVWall(hyper.Point{5, 5})
			},
			hyper.Placement{
				Actor: hyper.PlaceActorAt(defaultActorPlacement, map[hyper.Color]hyper.Point{hyper.Red: {5, 5}}),
				Goal:  hyper.PlaceGoalAt(hyper.Red, hyper.Point{3, 3}),
			},
			result{false, false, hyper.Point{5, 5}},
		},
		{
			"reached goal",
			hyper.Red,
			hyper.South,
			func(b *hyper.Board) {
				b.PutHWall(hyper.Point{5, 5})
			},
			hyper.Placement{
				Actor: hyper.PlaceActorAt(defaultActorPlacement, map[hyper.Color]hyper.Point{hyper.Red: {5, 4}}),
				Goal:  hyper.PlaceGoalAt(hyper.Red, hyper.Point{5, 5}),
			},
			result{true, true, hyper.Point{5, 5}},
		},
		{
			"reached black goal",
			hyper.Red,
			hyper.East,
			func(b *hyper.Board) {
				b.PutVWall(hyper.Point{5, 5})
			},
			hyper.Placement{
				Actor: hyper.PlaceActorAt(defaultActorPlacement, map[hyper.Color]hyper.Point{hyper.Red: {4, 5}}),
				Goal:  hyper.PlaceGoalAt(hyper.Black, hyper.Point{5, 5}),
			},
			result{true, true, hyper.Point{5, 5}},
		},
		{
			"move south",
			hyper.Red,
			hyper.South,
			func(b *hyper.Board) {
				b.PutHWall(hyper.Point{5, 10})
			},
			hyper.Placement{
				Actor: hyper.PlaceActorAt(defaultActorPlacement, map[hyper.Color]hyper.Point{hyper.Red: {5, 4}}),
				Goal:  hyper.PlaceGoalAt(hyper.Red, hyper.Point{5, 5}),
			},
			result{true, false, hyper.Point{5, 10}},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			board, err := hyper.NewBoard(hyper.Size{16, 16}, testcase.Placement)
			if err != nil {
				t.Fatal(err)
			}
			board.NewGame()
			testcase.Walls(board)

			actor, ok := board.Actors[testcase.ActorColor]
			if !ok {
				t.Fatalf("unable to find actor of color: %s", testcase.ActorColor)
			}


			_, ok = board.MoveActor(actor, testcase.Direction)

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
