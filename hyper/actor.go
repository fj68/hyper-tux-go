package hyper

type Actor struct {
	Color
	*Point
}

func (a *Actor) MoveTo(p *Point) {
	a.Point = p
}
