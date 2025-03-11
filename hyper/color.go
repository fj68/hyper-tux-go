package hyper

type Color int

const (
	Red Color = iota
	Green
	Blue
	Yellow
	Black
)

func (c Color) String() string {
	switch c {
	case Red:
		return "Red"
	case Green:
		return "Green"
	case Blue:
		return "Blue"
	case Yellow:
		return "Yellow"
	case Black:
		return "Black"
	}
	return "unknown Color"
}

var AllColors = []Color{
	Red,
	Green,
	Blue,
	Yellow,
	Black,
}
