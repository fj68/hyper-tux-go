package hyper

// Rect represents a rectangular region defined by top-left and bottom-right corners.
// The bottom-right corner is exclusive (not included in the rectangle).
type Rect struct {
	TopLeft, BottomRight Point // BottomRight is the edge of the Rect, so will not contained. e.g. Rect{(0, 0) (5, 5)}.Contains((5, 5)) == false
}

func NewRect(topLeft Point, size Size) Rect {
	bottomRight := Point{
		topLeft.X + size.W,
		topLeft.Y + size.H,
	}
	return Rect{topLeft, bottomRight}
}

// Size returns the dimensions of the rectangle.
func (r *Rect) Size() Size {
	return Size{
		r.BottomRight.X - r.TopLeft.X,
		r.BottomRight.Y - r.TopLeft.Y,
	}
}

// Contains returns true if the point is within the rectangle (excluding the bottom-right edge).
func (r *Rect) Contains(p Point) bool {
	return (r.TopLeft.X <= p.X && p.X < r.BottomRight.X &&
		r.TopLeft.Y <= p.Y && p.Y < r.BottomRight.Y)
}
