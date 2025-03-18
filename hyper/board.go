package hyper

import (
	"fmt"
	"maps"
	"math/rand"
	"slices"
	"time"

	"github.com/fj68/hyper-tux-go/slicetools"
)

type Board struct {
	rand *rand.Rand
	*History
	Goal
	*Mapdata
	Actors       map[Color]*Actor
	ColorWeights []int
	Goaled       bool
}

func NewBoard(size Size) (*Board, error) {
	b := &Board{
		rand:    rand.New(rand.NewSource(time.Now().Unix())),
		History: &History{},
		Actors:  map[Color]*Actor{},
		Mapdata: NewMapdata(size),
	}

	// place actors
	for _, color := range AllColors {
		b.Actors[color] = &Actor{Color: color, Point: Point{0, 0}}
		if err := b.PlaceActorAtRandom(color); err != nil {
			return nil, err
		}
	}

	return b, nil
}

func (b *Board) NewGame() error {
	b.Goaled = false
	b.History.Reset()
	return b.PlaceGoalAtRandom()
}

func (b *Board) ActorAt(p Point) (actor *Actor, exists bool) {
	for _, actor = range b.Actors {
		if actor.Point.Equals(p) {
			return actor, true
		}
	}
	return
}

func (b *Board) RandomPlace() (p Point, ok bool) {
	for range 50 {
		p = Point{
			b.rand.Intn(b.Mapdata.Size.W),
			b.rand.Intn(b.Mapdata.Size.H),
		}
		c := b.Mapdata.Center()
		if !c.Contains(p) {
			return p, true
		}
	}
	return
}

func (b *Board) PlaceActorAtRandom(color Color) error {
	_, ok := b.Actors[color]
	if !ok {
		return fmt.Errorf("unable to find actor of color: %s", color)
	}

	for range 50 {
		pos, ok := b.RandomPlace()
		_, exists := b.ActorAt(pos)
		if ok && !exists {
			b.Actors[color].MoveTo(pos)
			return nil
		}
	}
	return fmt.Errorf("unable to place actor: %s", color)
}

func (b *Board) PlaceGoalAtRandom() error {
	for range 50 {
		pos, ok := b.RandomPlace()
		_, exists := b.ActorAt(pos)
		if ok && !exists {
			color := RandomColor()
			b.Goal = Goal{color, pos}
			return nil
		}
	}
	return fmt.Errorf("unable to place goal")
}

func (b *Board) MoveActor(actor *Actor, d Direction) (pos Point, ok bool) {
	pos = b.NextStop(actor.Point, d)
	if actor.Point.Equals(pos) {
		// unable to move to the direction
		return
	}
	ok = true
	b.History.Push(&Record{
		Color:     actor.Color,
		Direction: d,
		Start:     actor.Point,
		End:       pos,
	})
	actor.MoveTo(pos)

	if b.Goal.Reached(*actor) {
		b.Goaled = true
	}
	return
}

func (b *Board) NextStop(current Point, d Direction) Point {
	switch d {
	case North:
		return b.nextStopNorth(current)
	case West:
		return b.nextStopWest(current)
	case South:
		return b.nextStopSouth(current)
	case East:
		return b.nextStopEast(current)
	}
	return Point{}
}

func (b *Board) nextStopNorth(current Point) Point {
	// find y-index of actor who is:
	//   1. on the current column
	//   2. nearer to the north than current
	actors := slicetools.FilterMap(
		slices.Collect(maps.Values(b.Actors)),
		func(actor *Actor) bool {
			return actor.X == current.X && actor.Y < current.Y
		},
		func(actor *Actor) int {
			return actor.Y + 1
		},
	)

	// find y-index of wall which is:
	//   1. on the current column
	//   2. nearer to the west than current
	walls := slicetools.Filter(b.HWalls[current.X], func(wall int) bool {
		return wall <= current.Y
	})

	// find x which is nearest to the current position
	ys := []int{0}
	ys = append(ys, actors...)
	ys = append(ys, walls...)
	y := slices.Max(ys)

	return Point{current.X, y}
}

func (b *Board) nextStopSouth(current Point) Point {
	// find y-index of actor who is:
	//   1. on the current column
	//   2. nearer to the south than current
	actors := slicetools.FilterMap(
		slices.Collect(maps.Values(b.Actors)),
		func(actor *Actor) bool {
			return actor.X == current.X && actor.Y > current.Y
		},
		func(actor *Actor) int {
			return actor.Y
		},
	)

	// find y-index of wall which is:
	//   1. on the current column
	//   2. nearer to the south than current
	walls := slicetools.Filter(b.HWalls[current.X], func(wall int) bool {
		return wall > current.Y
	})

	// find x which is nearest to the current position
	ys := []int{b.Mapdata.H}
	ys = append(ys, actors...)
	ys = append(ys, walls...)
	y := slices.Min(ys) - 1

	return Point{current.X, y}
}

func (b *Board) nextStopWest(current Point) Point {
	// find x-indices of actors who are:
	//   1. on the current row
	//   2. nearer to the west than current
	actors := slicetools.FilterMap(
		slices.Collect(maps.Values(b.Actors)),
		func(actor *Actor) bool {
			return actor.Y == current.Y && actor.X < current.X
		},
		func(actor *Actor) int {
			return actor.X + 1
		},
	)

	// find x-indices of walls which are:
	//   1. on the current row
	//   2. nearer to the west than current
	walls := slicetools.Filter(b.VWalls[current.Y], func(wall int) bool {
		return wall <= current.X
	})

	// find x which is nearest to the current position
	xs := []int{0}
	xs = append(xs, actors...)
	xs = append(xs, walls...)
	x := slices.Max(xs)

	return Point{x, current.Y}
}

func (b *Board) nextStopEast(current Point) Point {
	// find x-index of actor who is:
	//   1. on the current row
	//   2. nearer to the east than current
	actors := slicetools.FilterMap(
		slices.Collect(maps.Values(b.Actors)),
		func(actor *Actor) bool {
			return actor.Y == current.Y && actor.X > current.X
		},
		func(actor *Actor) int {
			return actor.X
		},
	)

	// find x-index of wall which is:
	//   1. on the current row
	//   2. nearer to the east than current
	walls := slicetools.Filter(b.VWalls[current.Y], func(wall int) bool {
		return wall > current.X
	})

	// find x which is nearest to the current position
	xs := []int{b.Mapdata.W}
	xs = append(xs, actors...)
	xs = append(xs, walls...)
	x := slices.Min(xs) - 1

	return Point{x, current.Y}
}
