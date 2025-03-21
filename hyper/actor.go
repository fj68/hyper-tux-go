package hyper

type Actor struct {
	Color
	Point
}

func (a *Actor) Equals(other Actor) bool {
	return a.Color == other.Color && a.Point.Equals(other.Point)
}

func (a *Actor) MoveTo(p Point) {
	a.Point = p
}
