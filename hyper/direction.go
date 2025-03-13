package hyper

type Direction int

const (
	North Direction = 1
	West Direction = 2
	East Direction = 4
	South Direction = 8
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
