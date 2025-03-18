package hyper

type Record struct {
	Color
	Direction
	Start, End Point
	Goaled     bool
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
	h.last--
	return h.records[h.last+1]
}

func (h *History) Redo() *Record {
	if h.last+1 < len(h.records) {
		h.last++
		return h.records[h.last]
	}
	return nil
}
