package hyper

type Direction int

const (
	North Direction = iota
	West
	East
	South
)

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
