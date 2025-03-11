package hyper

type Goal struct {
	Color
	Point
}

func (g Goal) Reached(actor Actor) bool {
	return (g.Color == actor.Color &&
		actor.Point.Equals(actor.Point))
}
