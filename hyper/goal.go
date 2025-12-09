package hyper

// Goal represents a target location for actors to reach.
type Goal struct {
	Color
	Point
}

// Reached returns true if the actor has reached this goal.
// An actor reaches a goal if the color matches and the position matches.
func (g Goal) Reached(actor Actor) bool {
	colorIsSame := g.Color == Black || g.Color == actor.Color
	posIsSame := g.Point.Equals(actor.Point)
	return (colorIsSame && posIsSame)
}
