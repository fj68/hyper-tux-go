package hyper

// Direction represents a direction in the game world.
type Direction int

// Direction constants for movement in cardinal directions.
const (
	North Direction = 1
	West Direction = 2
	East Direction = 4
	South Direction = 8
)

// String returns the string representation of the direction.
func (d Direction) String() string {
	switch d {
	case North:
		return "North"
	case West:
		return "West"
	case East:
		return "East"
	case South:
		return "South"
	}
	return "unknown Direction"
}
