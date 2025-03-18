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

type SwipeEventHandler interface {
	HandlePressed() bool
	HandleReleased() (e *SwipeEvent, released bool)
}

type SwipeEventDispatcher struct {
	q              *list.List // of *SwipeEvent
	EventHandlers  []SwipeEventHandler
	currentHandler SwipeEventHandler
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
		if handler.HandlePressed() {
			d.currentHandler = handler
			break
		}
	}
	return
}

func (d *SwipeEventDispatcher) handleReleased() {
	e, released := d.currentHandler.HandleReleased()
	if e != nil {
		d.q.PushBack(e)
	}
	if released {
		d.currentHandler = nil
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
	start *Position
}

func (h *MouseEventHandler) HandlePressed() bool {
	if !inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return false
	}
	x, y := ebiten.CursorPosition()
	h.start = &Position{X: float32(x), Y: float32(y)}
	return true
}

func (h *MouseEventHandler) HandleReleased() (*SwipeEvent, bool) {
	if !inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		return nil, false
	}
	x, y := ebiten.CursorPosition()
	pos := Position{X: float32(x), Y: float32(y)}
	start := h.start.ToPoint(CELL_SIZE)
	end := pos.ToPoint(CELL_SIZE)
	h.start = nil
	if start.Equals(end) {
		return nil, true
	}
	return &SwipeEvent{start, end}, true
}

type TouchEventHandler struct {
	start Position
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
	h.start = Position{X: float32(x), Y: float32(y)}
	return true
}

func (h *TouchEventHandler) HandleReleased() (*SwipeEvent, bool) {
	touchIDs := inpututil.AppendJustReleasedTouchIDs([]ebiten.TouchID{})
	for _, touchID := range touchIDs {
		if touchID == h.id {
			x, y := ebiten.TouchPosition(touchID)
			pos := Position{X: float32(x), Y: float32(y)}
			start := h.start.ToPoint(CELL_SIZE)
			end := pos.ToPoint(CELL_SIZE)
			if start.Equals(end) {
				return nil, true
			}
			return &SwipeEvent{start, end}, true
		}
	}
	return nil, false
}
