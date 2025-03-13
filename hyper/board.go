package hyper

import (
	"fmt"
	"math/rand"
	"time"
)

type VWall = []int // | : list of column indices where the wall exists
type HWall = []int // _ : list of row indices where the wall exists

type Board struct {
	rand *rand.Rand
	*Counter
	*Size
	*Goal
	VWalls []VWall
	HWalls []HWall
	Actors map[Color]*Actor
}

func NewBoard(size *Size) (*Board, error) {
	b := &Board{
		rand:    rand.New(rand.NewSource(time.Now().Unix())),
		Counter: &Counter{},
		Size:    size,
		Actors:  map[Color]*Actor{},
		VWalls:  make([]VWall, size.H),
		HWalls:  make([]HWall, size.W),
	}

	// place walls
	b.initCenterWalls()

	// place actors
	for _, color := range AllColors {
		b.PlaceActorAtRandom(color)
	}

	return b, nil
}

func (b *Board) ActorExists(p *Point) bool {
	for _, actor := range b.Actors {
		if actor.Point.Equals(p) {
			return true
		}
	}
	return false
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

func (b *Board) Center() *Rect {
	c := b.Size.Center()
	return NewRect(&Point{c.X - 1, c.Y - 1}, &Size{2, 2})
}

func (b *Board) PutHWall(p *Point) bool {
	for _, wall := range b.HWalls[p.X] {
		if wall == p.Y {
			return false
		}
	}
	b.HWalls[p.X] = append(b.HWalls[p.X], p.Y)
	return true
}

func (b *Board) PutVWall(p *Point) bool {
	for _, wall := range b.VWalls[p.Y] {
		if wall == p.X {
			return false
		}
	}
	b.VWalls[p.Y] = append(b.VWalls[p.Y], p.X)
	return true
}

func (b *Board) initCenterWalls() {
	r := b.Center()

	for x := r.TopLeft.X; x < r.BottomRight.X; x++ {
		b.HWalls[x] = []int{r.TopLeft.Y, r.BottomRight.Y - 1}
	}
	for y := r.TopLeft.Y; y < r.BottomRight.Y; y++ {
		b.VWalls[y] = []int{r.TopLeft.X, r.BottomRight.X - 1}
	}
}

func (b *Board) RandomPlace() (p *Point, ok bool) {
	for i := 0; i < 50; i++ {
		p = &Point{
			b.rand.Intn(b.Size.W),
			b.rand.Intn(b.Size.H),
		}
		if !b.Center().Contains(p) {
			return p, true
		}
	}
	return
}

func (b *Board) RandomColor(colors []Color) Color {
	return colors[b.rand.Intn(len(colors))]
}

func (b *Board) PlaceGoalAtRandom() error {
	pos, ok := b.RandomPlace()
	if !ok || b.ActorExists(pos) {
		return fmt.Errorf("unable to place goal")
	}
	color := b.RandomColor(AllColors)
	b.Goal = &Goal{color, pos}
	return nil
}

func (b *Board) NewGame() error {
	b.Counter.Reset()
	b.PlaceGoalAtRandom()
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
