package hyper

type Goal struct {
	Color
	Point
}

func (g Goal) Reached(actor *Actor) bool {
	colorIsSame := g.Color == Black || g.Color == actor.Color
	posIsSame := g.Point.Equals(actor.Point)
	return (colorIsSame && posIsSame)
}
