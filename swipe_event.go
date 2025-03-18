package main

import (
	"container/list"

	"github.com/fj68/hyper-tux-go/hyper"
	"github.com/fj68/hyper-tux-go/set"
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

type SwipeEventHandler interface {
	HandlePressed() bool
	HandleReleased() *SwipeEvent
}

type SwipeEventDispatcher struct {
	q              *list.List // of *SwipeEvent
	EventHandlers  set.Set[SwipeEventHandler]
	currentHandler SwipeEventHandler
}

func NewSwipeEventDispather() *SwipeEventDispatcher {
	return &SwipeEventDispatcher{
		q: list.New(),
	}
}

func (d *SwipeEventDispatcher) Update() error {
	if d.currentHandler == nil {
		d.handlePressed()
	} else {
		d.handleReleased()
	}
	return nil
}

func (d *SwipeEventDispatcher) handlePressed() {
	for handler := range d.EventHandlers {
		if handler.HandlePressed() {
			d.currentHandler = handler
			break
		}
	}
	return
}

func (d *SwipeEventDispatcher) handleReleased() {
	if e := d.currentHandler.HandleReleased(); e != nil {
		d.q.PushBack(e)
	}
}

func (d *SwipeEventDispatcher) Len() int {
	return d.q.Len()
}

func (d *SwipeEventDispatcher) Pop() *SwipeEvent {
	front := d.q.Front()
	if front == nil {
		return nil
	}
	v := d.q.Remove(front)
	e, ok := v.(*SwipeEvent)
	if !ok {
		return nil
	}
	return e
}

type MouseEventHandler struct {
	start hyper.Point
}

func (h *MouseEventHandler) HandlePressed() bool {
	if !inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return false
	}
	x, y := ebiten.CursorPosition()
	h.start = hyper.Point{X: x, Y: y}
	return true
}

func (h *MouseEventHandler) HandleReleased() *SwipeEvent {
	if !inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		return nil
	}
	x, y := ebiten.CursorPosition()
	end := hyper.Point{X: x, Y: y}
	if h.start.Equals(end) {
		return nil
	}
	return &SwipeEvent{h.start, end}
}

type TouchEventHandler struct {
	start hyper.Point
	id    ebiten.TouchID
}

func (h *TouchEventHandler) HandlePressed() bool {
	touchIDs := inpututil.AppendJustPressedTouchIDs([]ebiten.TouchID{})
	if len(touchIDs) < 1 {
		return false
	}
	// handle only first input
	x, y := ebiten.TouchPosition(touchIDs[0])
	h.id = touchIDs[0]
	h.start = hyper.Point{X: x, Y: y}
	return true
}

func (h *TouchEventHandler) HandleReleased() *SwipeEvent {
	touchIDs := inpututil.AppendJustReleasedTouchIDs([]ebiten.TouchID{})
	for _, touchID := range touchIDs {
		if touchID == h.id {
			x, y := ebiten.TouchPosition(touchID)
			end := hyper.Point{X: x, Y: y}
			if h.start.Equals(end) {
				return nil
			}
			return &SwipeEvent{h.start, end}
		}
	}
	return nil
}
