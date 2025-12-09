// Package hyper contains the core game model types and utilities used by
// the Hyper Tux game. It defines geometric primitives, board/map data,
// actors, goals, and related algorithms such as movement, placement and
// history tracking.
package hyper

import (
	"fmt"
	"maps"
	"math/rand"
	"slices"
	"time"

	"github.com/fj68/hyper-tux-go/internal/slicetools"
)

// Board represents the game board with actors, walls, and goals.
type Board struct {
	rand    *rand.Rand
	history *History
	Placement
	Goal
	*Mapdata
	Actors       map[Color]*Actor
	ColorWeights []int
	Goaled       bool
}

// NewBoard creates and initializes a new game board with the given size and placement algorithms.
func NewBoard(size Size, p Placement) (*Board, error) {
	b := &Board{
		rand:      rand.New(rand.NewSource(time.Now().Unix())),
		history:   &History{},
		Actors:    map[Color]*Actor{},
		Mapdata:   NewMapdata(size),
		Placement: p,
	}

	// place actors
	for _, color := range AllColors {
		b.Actors[color] = &Actor{Color: color, Point: Point{0, 0}}
		if err := b.PlaceActor(color); err != nil {
			return nil, err
		}
	}

	b.NewGame()

	return b, nil
}

// NewGame resets the board for a new game, clearing history and placing a new goal.
func (b *Board) NewGame() error {
	b.Goaled = false
	b.history.Reset()
	if err := b.PlaceGoal(); err != nil {
		return err
	}
	return nil
}

// SomethingExists returns true if an actor, goal, or center box exists at the given position.
func (b *Board) SomethingExists(pos Point) bool {
	_, exists := b.ActorAt(pos)
	c := b.Mapdata.Center()
	return exists || pos.Equals(b.Goal.Point) || c.Contains(pos)
}

// PlaceActor places an actor on a random unoccupied position on the board.
func (b *Board) PlaceActor(color Color) error {
	for range 50 {
		pos := b.Placement.Actor(b)
		if !b.SomethingExists(pos) {
			b.Actors[color].MoveTo(pos)
			return nil
		}
	}
	return fmt.Errorf("unable to place %s actor", color)
}

// PlaceGoal places a goal with a random color on a random unoccupied position.
func (b *Board) PlaceGoal() error {
	var pos Point
	for range 50 {
		pos = b.Placement.Goal(b)
		if !b.SomethingExists(pos) {
			b.Goal = Goal{
				Color: RandomColor(),
				Point: pos,
			}
			return nil
		}
	}
	return fmt.Errorf("unable to place goal")
}

// History returns all recorded moves on this board.
func (b *Board) History() []*Record {
	return b.history.Records()
}

// Steps returns the number of moves taken so far.
func (b *Board) Steps() int {
	return b.history.Len()
}

// PlaceActors positions multiple actors at specific locations.
func (b *Board) PlaceActors(actors map[Color]Point) {
	for c, p := range actors {
		b.Actors[c] = &Actor{c, p}
	}
}

// ActorAt returns the actor at the given position, if any.
func (b *Board) ActorAt(p Point) (actor *Actor, exists bool) {
	for _, actor = range b.Actors {
		if actor.Point.Equals(p) {
			return actor, true
		}
	}
	return
}

// MoveActor moves an actor in the given direction and returns success and goal-reached status.
func (b *Board) MoveActor(actor *Actor, d Direction) (pos Point, ok bool) {
	pos = b.NextStop(actor.Point, d)
	if actor.Point.Equals(pos) {
		// unable to move to the direction
		return
	}
	ok = true
	b.history.Push(&Record{
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

// Reset undoes all moves, returning the board to its initial state.
func (b *Board) Reset() {
	for b.history.Len() > 0 {
		b.Undo()
	}
}

// Undo reverts the last move.
func (b *Board) Undo() {
	r := b.history.Undo()
	if r == nil {
		return
	}
	b.Actors[r.Color].Point = r.Start
	b.Goaled = false
}

// Redo replays the next move in history.
func (b *Board) Redo() {
	r := b.history.Redo()
	if r == nil {
		return
	}
	b.Actors[r.Color].Point = r.End
	if b.Goal.Reached(*b.Actors[r.Color]) {
		b.Goaled = true
	}
}

// NextStop calculates where an actor moving in a direction would stop.
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
