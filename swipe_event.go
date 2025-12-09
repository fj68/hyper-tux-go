package main

import (
	"container/list"

	"github.com/fj68/hyper-tux-go/hyper"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// SwipeEvent represents a swipe gesture with start and end points.
type SwipeEvent struct {
	Start, End hyper.Point
}

// Direction returns the direction of the swipe based on the start and end points.
// It calculates which direction (North, South, East, or West) has the largest distance.
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

// SwipeEventHandler is an interface for handling swipe input events.
type SwipeEventHandler interface {
	HandlePressed() (start *Position)
	HandleReleased() (end *Position)
}

// SwipeEventDispatcher manages and dispatches swipe events from multiple event handlers.
type SwipeEventDispatcher struct {
	q              *list.List // of *SwipeEvent
	EventHandlers  []SwipeEventHandler
	currentHandler SwipeEventHandler
	start          *Position
}

// NewSwipeEventDispather creates a new SwipeEventDispatcher with the given event handlers.
func NewSwipeEventDispather(handlers ...SwipeEventHandler) *SwipeEventDispatcher {
	return &SwipeEventDispatcher{
		q:             list.New(),
		EventHandlers: handlers,
	}
}

// Update processes input events and generates SwipeEvents from handlers.
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

// Len returns the number of pending swipe events in the queue.
func (d *SwipeEventDispatcher) Len() int {
	return d.q.Len()
}

// Push adds a swipe event to the queue.
func (d *SwipeEventDispatcher) Push(ev *SwipeEvent) {
	d.q.PushBack(ev)
}

// Pop removes and returns the next swipe event from the queue.
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

// MouseEventHandler handles mouse input events.
type MouseEventHandler struct{}

// HandlePressed returns the cursor position if the left mouse button is just pressed, otherwise nil.
func (h *MouseEventHandler) HandlePressed() *Position {
	if !inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return nil
	}
	x, y := ebiten.CursorPosition()
	return &Position{X: float32(x), Y: float32(y)}
}

// HandleReleased returns the cursor position if the left mouse button is just released, otherwise nil.
func (h *MouseEventHandler) HandleReleased() *Position {
	if !inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		return nil
	}
	x, y := ebiten.CursorPosition()
	return &Position{X: float32(x), Y: float32(y)}
}

// TouchEventHandler handles touch input events.
type TouchEventHandler struct {
	id ebiten.TouchID
}

// HandlePressed returns the position of the first newly pressed touch, otherwise nil.
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

// HandleReleased returns the position of a touch that was being tracked if it just released, otherwise nil.
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
