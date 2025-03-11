package hyper

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fj68/hyper-tux-go/slicetools"
)

type VWall = []int // | : list of column indices where the wall exists
type HWall = []int // _ : list of row indices where the wall exists

type Board struct {
	rand *rand.Rand
	Size
	Goal
	Counter
	VWalls []VWall
	HWalls []HWall
	Actors []Actor
}

func NewBoard(size Size) *Board {
	b := &Board{
		Size:   size,
		VWalls: make([]VWall, size.H),
		HWalls: make([]HWall, size.W),
	}
	b.initCenterWalls()
	return b
}

func (b *Board) Center() Rect {
	c := b.Size.Center()
	return NewRect(Point{c.X - 1, c.Y - 1}, Size{2, 2})
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

func (b *Board) Actor(color Color) (int, Actor) {
	for i, actor := range b.Actors {
		if actor.Color == color {
			return i, actor
		}
	}
	return -1, Actor{}
}

func (b *Board) RandomPlace(validator func(Point) bool) (p Point, ok bool) {
	for i := 0; i < 50; i++ {
		p = Point{
			b.rand.Intn(b.Size.W),
			b.rand.Intn(b.Size.H),
		}
		if !b.Center().Contains(p) && validator(p) {
			return p, true
		}
	}
	return
}

func (b *Board) RandomColor(colors []Color) Color {
	return colors[b.rand.Intn(len(colors))]
}

func (b *Board) NewGame() error {
	b.rand = rand.New(rand.NewSource(time.Now().Unix()))
	b.Counter.Reset()

	notOverlapped := func(p Point) bool {
		return slicetools.Every(b.Actors, func(actor Actor) bool {
			return !actor.Point.Equals(p)
		})
	}

	// place actors
	for i := range b.Actors {
		if p, ok := b.RandomPlace(notOverlapped); ok {
			b.Actors[i].Point = p
		} else {
			return fmt.Errorf("unable to place actor: %s", b.Actors[i].Color)
		}
	}

	// place goal
	if p, ok := b.RandomPlace(notOverlapped); ok {
		b.Goal.Point = p
	} else {
		return fmt.Errorf("unable to place goal")
	}
	b.Goal.Color = b.RandomColor(AllColors)

	return nil
}

func (b *Board) MoveActor(color Color, d Direction) (ok bool, finished bool) {
	i, actor := b.Actor(color)

	pos := b.NextStop(actor.Point, d)
	if actor.Point.Equals(pos) {
		// unable to move to the direction
		return
	}
	ok = true
	b.Counter.Incr()
	b.Actors[i].Point = pos

	if b.Goal.Reached(b.Actors[i]) {
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
	walls := b.HWalls[current.X]
	y := 0
	for _, wall := range walls {
		if y < wall && wall <= current.Y {
			y = wall
		}
	}
	return Point{current.X, y}
}

func (b *Board) nextStopSouth(current Point) Point {
	walls := b.HWalls[current.X]
	y := b.Size.H - 1
	for _, wall := range walls {
		if wall < y && current.Y <= wall {
			y = wall
		}
	}
	return Point{current.X, y}
}

func (b *Board) nextStopWest(current Point) Point {
	walls := b.VWalls[current.Y]
	x := 0
	for _, wall := range walls {
		if x < wall && wall <= current.X {
			x = wall
		}
	}
	return Point{x, current.Y}
}

func (b *Board) nextStopEast(current Point) Point {
	walls := b.VWalls[current.Y]
	x := b.Size.W - 1
	for _, wall := range walls {
		if wall < x && current.X <= wall {
			x = wall
		}
	}
	return Point{x, current.Y}
}
