package hyper

// Record represents a move made by an actor.
type Record struct {
	Color
	Direction
	Start, End Point
}

// Equals returns true if both records describe the same move.
func (r *Record) Equals(other *Record) bool {
	return (r.Color == other.Color &&
		r.Direction == other.Direction &&
		r.Start.Equals(other.Start) &&
		r.End.Equals(other.End))
}

// History manages undo and redo functionality for game moves.
type History struct {
	records []*Record
	last    int
}

// Records returns all recorded moves.
func (h *History) Records() []*Record {
	return h.records
}

// Push adds a new move record to the history.
func (h *History) Push(r *Record) {
	h.records = append(h.records[:h.last], r)
	h.last++
}

// Reset clears all move history.
func (h *History) Reset() {
	h.records = nil
	h.last = 0
}

// Len returns the number of recorded moves.
func (h *History) Len() int {
	return h.last
}

// Undo reverts the last move and returns it.
func (h *History) Undo() *Record {
	if h.last < 1 {
		return nil
	}
	r := h.records[h.last-1]
	h.last--
	return r
}

// Redo replays the next move in history and returns it.
func (h *History) Redo() *Record {
	if h.last+1 <= len(h.records) {
		h.last++
		return h.records[h.last-1]
	}
	return nil
}
