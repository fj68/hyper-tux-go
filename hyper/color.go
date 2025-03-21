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

var ColorWeights = []int{
	22, // Red
	22, // Green
	22, // Blue
	22, // Yellow
	12, // Black
}

func RandomColor() Color {
	return Choice(AllColors, ColorWeights)
}
