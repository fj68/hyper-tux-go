package hyper_test

import (
	"testing"

	"github.com/fj68/hyper-tux-go/hyper"
)

func TestHistoryUndo(t *testing.T) {
	t.Run("Undo nothing", func(t *testing.T) {
		h := &hyper.History{}
		if r := h.Undo(); r != nil {
			t.Errorf("not nil: %+v", r)
		}
	})
	t.Run("Undo 1 record", func(t *testing.T) {
		r := &hyper.Record{
			Color:     hyper.Yellow,
			Direction: hyper.East,
			Start:     hyper.Point{5, 5},
			End:       hyper.Point{10, 5},
		}
		h := &hyper.History{}
		h.Push(r)
		actual := h.Undo()
		if actual == nil {
			t.Errorf("nil: %+v", h)
		}
		if !actual.Equals(r) {
			t.Errorf("difference:\n\texpected = %+v\n\t  actual = %+v", r, actual)
		}
	})
}

func TestHistoryRedo(t *testing.T) {
	t.Run("Redo nothing", func(t *testing.T) {
		h := &hyper.History{}
		if r := h.Redo(); r != nil {
			t.Errorf("not nil: %+v", r)
		}
	})
	t.Run("Redo 1 record", func(t *testing.T) {
		r := &hyper.Record{
			Color:     hyper.Yellow,
			Direction: hyper.East,
			Start:     hyper.Point{5, 5},
			End:       hyper.Point{10, 5},
		}
		h := &hyper.History{}
		h.Push(r)
		h.Undo()
		actual := h.Redo()
		if actual == nil {
			t.Errorf("nil: %+v", h)
		}
		if !actual.Equals(r) {
			t.Errorf("difference:\n\texpected = %+v\n\t  actual = %+v", r, actual)
		}
	})
}
