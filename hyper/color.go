package hyper

type Color int

// Color constants for game actors and goals.
const (
	Red Color = iota
	Green
	Blue
	Yellow
	Black
)

// String returns the string representation of the color.
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

// AllColors is a slice containing all valid Color values.
var AllColors = []Color{
	Red,
	Green,
	Blue,
	Yellow,
	Black,
}

// ColorWeights defines the relative probability weights for random color selection.
var ColorWeights = []int{
	22, // Red
	22, // Green
	22, // Blue
	22, // Yellow
	12, // Black
}

// RandomColor returns a random Color selected using weighted probabilities.
func RandomColor() Color {
	return Choice(AllColors, ColorWeights)
}
