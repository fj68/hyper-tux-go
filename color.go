package main

import (
	"image/color"

	"github.com/fj68/hyper-tux-go/hyper"
)

func Color(c hyper.Color) color.Color {
	switch c {
	case hyper.Black:
		return color.Black
	case hyper.Red:
		return color.RGBA{255, 0, 0, 255}
	case hyper.Green:
		return color.RGBA{0, 255, 0, 255}
	case hyper.Blue:
		return color.RGBA{0, 0, 255, 255}
	case hyper.Yellow:
		return color.RGBA{255, 255, 0, 255}
	}
	return color.Transparent
}

func Offset(color hyper.Color) float32 {
	switch color {
	case hyper.Red:
		return 0
	case hyper.Green:
		return 1
	case hyper.Blue:
		return -1
	case hyper.Yellow:
		return 2
	case hyper.Black:
		return -2
	}
	return 0
}
