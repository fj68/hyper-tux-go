package hyper

// Actor represents a player character on the game board.
type Actor struct {
	Color
	Point
}

// Equals returns true if both actors have the same color and position.
func (a *Actor) Equals(other Actor) bool {
	return a.Color == other.Color && a.Point.Equals(other.Point)
}

// MoveTo moves the actor to a new position.
func (a *Actor) MoveTo(p Point) {
	a.Point = p
}
