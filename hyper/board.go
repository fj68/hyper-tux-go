package hyper

import (
	"fmt"
	"math/rand"
	"time"
)

type Board struct {
	rand *rand.Rand
	Counter
	Goal
	*Mapdata
	Actors       map[Color]Actor
	ColorWeights []int
}

func NewBoard(size Size) (*Board, error) {
	b := &Board{
		rand:    rand.New(rand.NewSource(time.Now().Unix())),
		Counter: Counter{},
		Actors:  map[Color]Actor{},
		Mapdata: NewMapdata(size),
	}

	// place actors
	for _, color := range AllColors {
		b.Actors[color] = Actor{Color: color, Point: Point{0, 0}}
		b.PlaceActorAtRandom(color)
	}

	return b, nil
}

func (b *Board) NewGame() error {
	b.Counter.Reset()
	b.PlaceGoalAtRandom()
	return nil
}

func (b *Board) ActorAt(p Point) (Actor, bool) {
	for _, actor := range b.Actors {
		if actor.Point.Equals(p) {
			return actor, true
		}
	}
	return Actor{}, false
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
	actor, ok := b.Actors[color]
	if !ok {
		return fmt.Errorf("unable to find actor of color: %s", color)
	}

	pos, ok := b.RandomPlace()
	actor, exists := b.ActorAt(pos)
	if !ok || !exists {
		return fmt.Errorf("unable to place actor: %s", color)
	}
	actor.MoveTo(pos)
	return nil
}

func (b *Board) PlaceGoalAtRandom() error {
	pos, ok := b.RandomPlace()
	_, exists := b.ActorAt(pos)
	if !ok || !exists {
		return fmt.Errorf("unable to place goal")
	}
	color := RandomColor()
	b.Goal = Goal{color, pos}
	return nil
}

func (b *Board) MoveActor(actor Actor, d Direction) (pos Point, ok bool, finished bool) {
	pos = b.NextStop(actor.Point, d)
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
	// min of y-index
	y := 0

	// find y-index of actor who is:
	//   1. on the current column
	//   2. nearer to the north than current
	//   3. nearer to the current position than ever before
	for _, actor := range b.Actors {
		if actor.Point.X == current.X && actor.Y != current.Y && y < actor.Point.Y {
			y = actor.Point.Y
		}
	}

	// find y-index of wall which is:
	//   1. on the current column
	//   2. nearer to the west than current
	//   3. nearer to the current position than ever before
	for wall := range b.HWalls[current.X] {
		if y < wall && wall <= current.Y {
			y = wall
		}
	}

	return Point{current.X, y}
}

func (b *Board) nextStopSouth(current Point) Point {
	// max of x-index
	y := b.Size.H - 1

	// find y-index of actor who is:
	//   1. on the current column
	//   2. nearer to the south than current
	//   3. nearer to the current position than ever before
	for _, actor := range b.Actors {
		if actor.Point.X == current.X && actor.Y != current.Y && actor.Point.Y < y {
			y = actor.Point.Y
		}
	}

	// find y-index of wall which is:
	//   1. on the current column
	//   2. nearer to the south than current
	//   3. nearer to the current position than ever before
	for wall := range b.HWalls[current.X] {
		if wall < y && current.Y+1 <= wall {
			y = wall
		}
	}

	return Point{current.X, y}
}

func (b *Board) nextStopWest(current Point) Point {
	// min of x-index
	x := 0

	// find x-index of actor who is:
	//   1. on the current row
	//   2. nearer to the west than current
	//   3. nearer to the current position than ever before
	for _, actor := range b.Actors {
		if actor.Point.Y == current.Y && actor.X != current.X && x < actor.Point.X {
			x = actor.Point.X
		}
	}

	// find x-index of wall which is:
	//   1. on the current row
	//   2. nearer to the west than current
	//   3. nearer to the current position than ever before
	for wall := range b.VWalls[current.Y] {
		if x < wall && wall <= current.X {
			x = wall
		}
	}

	return Point{x, current.Y}
}

func (b *Board) nextStopEast(current Point) Point {
	// max of x-index
	x := b.Mapdata.Size.W - 1

	// find x-index of actor who is:
	//   1. on the current row
	//   2. nearer to the east than current
	//   3. nearer to the current position than ever before
	for _, actor := range b.Actors {
		if actor.Point.Y == current.Y && current.X+1 <= actor.Point.X && actor.Point.X < x {
			x = actor.Point.X
		}
	}

	// find x-index of wall which is:
	//   1. on the current row
	//   2. nearer to the east than current
	//   3. nearer to the current position than ever before
	for wall := range b.VWalls[current.Y] {
		if wall < x && current.X+1 <= wall {
			x = wall
		}
	}

	return Point{x, current.Y}
}
