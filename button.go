package main

import (
	"image/color"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)

func createButton(r *ResourceLoader, label string) (*widget.Button, error) {
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
		widget.ButtonOpts.Image(&widget.ButtonImage{
			Idle: image.NewNineSliceColor(color.White),
		}),
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
