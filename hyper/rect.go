package hyper

type Rect struct {
	TopLeft, BottomRight Point
}

func NewRect(topLeft Point, size Size) Rect {
	bottomRight := Point{
		topLeft.X + size.W,
		topLeft.Y + size.H,
	}
	return Rect{topLeft, bottomRight}
}

func (r Rect) Size() Size {
	return Size{
		r.BottomRight.X - r.TopLeft.X,
		r.BottomRight.Y - r.TopLeft.Y,
	}
}

func (r Rect) Contains(p Point) bool {
	return (r.TopLeft.X <= p.X && r.BottomRight.X < p.X &&
		r.TopLeft.Y <= p.Y && r.BottomRight.Y < p.Y)
}
