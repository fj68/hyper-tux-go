package hyper

type Record struct {
	Color
	Direction
	Start, End Point
	Goaled     bool
}

func (r *Record) Equals(other *Record) bool {
	return (r.Color == other.Color &&
		r.Direction == other.Direction &&
		r.Start.Equals(other.Start) &&
		r.End.Equals(other.End) &&
		r.Goaled == other.Goaled)
}

type History struct {
	records []*Record
	last    int
}

func (h *History) Push(r *Record) {
	h.records = append(h.records[:h.last], r)
	h.last++
}

func (h *History) Reset() {
	h.records = nil
	h.last = 0
}

func (h *History) Len() int {
	return h.last
}

func (h *History) Undo() *Record {
	if h.last < 1 {
		return nil
	}
	r := h.records[h.last-1]
	h.last--
	return r
}

func (h *History) Redo() *Record {
	if h.last+1 <= len(h.records) {
		h.last++
		return h.records[h.last-1]
	}
	return nil
}
