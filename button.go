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

func createButton(r *ResourceLoader, label string, onclick widget.ButtonClickedHandlerFunc) (*widget.Button, error) {
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
			widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
				Stretch:  false,
			}),
			widget.WidgetOpts.MinSize(76, 32),
		),
		widget.ButtonOpts.Image(img),
		widget.ButtonOpts.Text(label, font, &widget.ButtonTextColor{
			Idle: color.NRGBA{R: 0, G: 0, B: 0, A: 255},
		}),
		widget.ButtonOpts.TextPadding(widget.NewInsetsSimple(5)),
		widget.ButtonOpts.ClickedHandler(onclick),
	)
	return b, nil
}
