package main

import (
	"container/list"

	"github.com/fj68/hyper-tux-go/hyper"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type SwipeEvent struct {
	Start, End hyper.Point
}

func (e *SwipeEvent) Direction() hyper.Direction {
	distance := e.End.Sub(e.Start)
	if distance.Abs().X > distance.Abs().Y {
		if distance.X < 0 {
			return hyper.West
		}
		return hyper.East
	}
	if distance.Y < 0 {
		return hyper.North
	}
	return hyper.South
}

type SwipeEventDispatcher struct {
	q     *list.List // of SwipeEvent
	id    ebiten.TouchID
	start hyper.Point
}

func NewSwipeEventDispather() SwipeEventDispatcher {
	return SwipeEventDispatcher{
		q: list.New(),
	}
}

func (d *SwipeEventDispatcher) handleTouchPressed() error {
	touchIDs := inpututil.AppendJustPressedTouchIDs([]ebiten.TouchID{})
	for _, touchID := range touchIDs {
		// handle only first input
		if touchID == 0 || touchID == d.id {
			x, y := ebiten.TouchPosition(touchID)
			d.id = touchID
			d.start = hyper.Point{X: x, Y: y}
			break
		}
	}
	return nil
}

func (d *SwipeEventDispatcher) Update() error {
	touchIDs := inpututil.AppendJustReleasedTouchIDs([]ebiten.TouchID{})
	for _, touchID := range touchIDs {
		if touchID == d.id {
			x, y := ebiten.TouchPosition(touchID)
			end := hyper.Point{X: x, Y: y}
			d.id = 0
			d.q.PushBack(SwipeEvent{Start: d.start, End: end})
			break
		}
	}
	return nil
}

func (d *SwipeEventDispatcher) Len() int {
	return d.q.Len()
}

func (d *SwipeEventDispatcher) Pop() (SwipeEvent, bool) {
	front := d.q.Front()
	if front == nil {
		return SwipeEvent{}, false
	}
	v := d.q.Remove(front)
	e, ok := v.(SwipeEvent)
	return e, ok
}
