package main

import (
	"image/color"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)

func loadButtonImage() (*widget.ButtonImage, error) {

	idle := image.NewNineSliceColor(color.NRGBA{R: 170, G: 170, B: 180, A: 255})

	hover := image.NewNineSliceColor(color.NRGBA{R: 130, G: 130, B: 150, A: 255})

	pressed := image.NewNineSliceColor(color.NRGBA{R: 130, G: 130, B: 150, A: 255})

	return &widget.ButtonImage{
		Idle:    idle,
		Hover:   hover,
		Pressed: pressed,
	}, nil
}

func createButton(r *ResourceLoader, label string) (*widget.Button, error) {
	img, err := loadButtonImage()
	if err != nil {
		return nil, err
	}
	font, err := r.FontFace(12)
	if err != nil {
		return nil, err
	}
	b := widget.NewButton(
		widget.ButtonOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
				VerticalPosition:   widget.AnchorLayoutPositionCenter,
			}),
		),
		widget.ButtonOpts.Image(img),
		widget.ButtonOpts.Text(label, font, &widget.ButtonTextColor{
			Idle: color.NRGBA{0xff, 0xff, 0xff, 0xff},
		}),
		widget.ButtonOpts.TextPadding(widget.Insets{
			Left:   30,
			Right:  30,
			Top:    5,
			Bottom: 5,
		}),
	)
	return b, nil
}
