package hyper

import (
	"fmt"
	"math/rand"
	"time"
)

type Board struct {
	rand *rand.Rand
	*Counter
	*Goal
	*Mapdata
	Actors       map[Color]*Actor
	ColorWeights []int
}

func NewBoard(size *Size) (*Board, error) {
	b := &Board{
		rand:    rand.New(rand.NewSource(time.Now().Unix())),
		Counter: &Counter{},
		Actors:  map[Color]*Actor{},
		Mapdata: NewMapdata(size),
	}

	// place actors
	for _, color := range AllColors {
		b.PlaceActorAtRandom(color)
	}

	return b, nil
}

func (b *Board) NewGame() error {
	b.Counter.Reset()
	b.PlaceGoalAtRandom()
	return nil
}

func (b *Board) ActorExists(p *Point) bool {
	for _, actor := range b.Actors {
		if actor.Point.Equals(p) {
			return true
		}
	}
	return false
}

func (b *Board) RandomPlace() (p *Point, ok bool) {
	for range 50 {
		p = &Point{
			b.rand.Intn(b.Mapdata.Size.W),
			b.rand.Intn(b.Mapdata.Size.H),
		}
		if !b.Mapdata.Center().Contains(p) {
			return p, true
		}
	}
	return
}

func (b *Board) PlaceActorAtRandom(color Color) error {
	actor, ok := b.Actors[color]
	if !ok {
		return fmt.Errorf("unable to find actor of color: %s", color)
	}

	pos, ok := b.RandomPlace()
	if !ok || b.ActorExists(pos) {
		return fmt.Errorf("unable to place actor: %s", color)
	}
	actor.MoveTo(pos)
	return nil
}

func (b *Board) PlaceGoalAtRandom() error {
	pos, ok := b.RandomPlace()
	if !ok || b.ActorExists(pos) {
		return fmt.Errorf("unable to place goal")
	}
	color := RandomColor()
	b.Goal = &Goal{color, pos}
	return nil
}

func (b *Board) MoveActor(color Color, d Direction) (ok bool, finished bool) {
	actor, ok := b.Actors[color]
	if !ok {
		return
	}

	pos := b.NextStop(actor.Point, d)
	if actor.Point.Equals(pos) {
		// unable to move to the direction
		return
	}
	ok = true
	b.Counter.Incr()
	actor.MoveTo(pos)

	if b.Goal.Reached(actor) {
		finished = true
	}
	return
}

func (b *Board) NextStop(current *Point, d Direction) *Point {
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
	return nil
}

func (b *Board) nextStopNorth(current *Point) *Point {
	walls := b.HWalls[current.X]
	y := 0
	for _, wall := range walls {
		if y < wall && wall <= current.Y {
			y = wall
		}
	}
	return &Point{current.X, y}
}

func (b *Board) nextStopSouth(current *Point) *Point {
	walls := b.HWalls[current.X]
	y := b.Size.H - 1
	for _, wall := range walls {
		if wall < y && current.Y+1 <= wall {
			y = wall
		}
	}
	return &Point{current.X, y}
}

func (b *Board) nextStopWest(current *Point) *Point {
	walls := b.VWalls[current.Y]
	x := 0
	for _, wall := range walls {
		if x < wall && wall <= current.X {
			x = wall
		}
	}
	return &Point{x, current.Y}
}

func (b *Board) nextStopEast(current *Point) *Point {
	walls := b.VWalls[current.Y]
	x := b.Size.W - 1
	for _, wall := range walls {
		if wall < x && current.X+1 <= wall {
			x = wall
		}
	}
	return &Point{x, current.Y}
}
