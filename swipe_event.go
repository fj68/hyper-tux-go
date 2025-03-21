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
	diff := e.End.Sub(e.Start)
	distance := diff.Abs()
	if distance.X > distance.Y {
		if diff.X < 0 {
			return hyper.West
		}
		return hyper.East
	}
	if diff.Y < 0 {
		return hyper.North
	}
	return hyper.South
}

type SwipeEventHandler interface {
	HandlePressed() (start *Position)
	HandleReleased() (end *Position)
}

type SwipeEventDispatcher struct {
	q              *list.List // of *SwipeEvent
	EventHandlers  []SwipeEventHandler
	currentHandler SwipeEventHandler
	start          *Position
}

func NewSwipeEventDispather(handlers ...SwipeEventHandler) *SwipeEventDispatcher {
	return &SwipeEventDispatcher{
		q:             list.New(),
		EventHandlers: handlers,
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
	for _, handler := range d.EventHandlers {
		d.start = handler.HandlePressed()
		if d.start != nil {
			d.currentHandler = handler
			break
		}
	}
	return
}

func (d *SwipeEventDispatcher) handleReleased() {
	pos := d.currentHandler.HandleReleased()
	if pos == nil {
		return
	}

	start := d.start.ToPoint(CELL_SIZE)
	end := pos.ToPoint(CELL_SIZE)

	d.start = nil
	d.currentHandler = nil

	if start.Equals(end) {
		return
	}

	d.q.PushBack(&SwipeEvent{start, end})
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
	ev, ok := v.(*SwipeEvent)
	if !ok {
		return nil
	}
	return ev
}

type MouseEventHandler struct{}

func (h *MouseEventHandler) HandlePressed() *Position {
	if !inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return nil
	}
	x, y := ebiten.CursorPosition()
	return &Position{X: float32(x), Y: float32(y)}
}

func (h *MouseEventHandler) HandleReleased() *Position {
	if !inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		return nil
	}
	x, y := ebiten.CursorPosition()
	return &Position{X: float32(x), Y: float32(y)}
}

type TouchEventHandler struct {
	id ebiten.TouchID
}

func (h *TouchEventHandler) HandlePressed() *Position {
	touchIDs := inpututil.AppendJustPressedTouchIDs([]ebiten.TouchID{})
	if len(touchIDs) < 1 {
		return nil
	}
	// handle only first input
	x, y := ebiten.TouchPosition(touchIDs[0])
	h.id = touchIDs[0]
	return &Position{X: float32(x), Y: float32(y)}
}

func (h *TouchEventHandler) HandleReleased() *Position {
	touchIDs := inpututil.AppendJustReleasedTouchIDs([]ebiten.TouchID{})
	for _, touchID := range touchIDs {
		if touchID == h.id {
			x, y := ebiten.TouchPosition(touchID)
			return &Position{X: float32(x), Y: float32(y)}
		}
	}
	return nil
}
