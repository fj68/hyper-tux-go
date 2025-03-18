package main

import (
	"bytes"
	"image/color"
	"log"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/fj68/hyper-tux-go/hyper"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/gofont/goregular"
)

func createButton(label string, onClick widget.ButtonClickedHandlerFunc) (*widget.Button, error) {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		log.Fatal(err)
	}

	face := &text.GoTextFace{
		Source: s,
		Size:   32,
	}

	idle, err := loadImageNineSlice("assets/button-idle.png", 12, 0)
	if err != nil {
		return nil, err
	}

	buttonImage := &widget.ButtonImage{
		Idle:         idle,
		Pressed:      idle,
		Hover:        idle,
		PressedHover: idle,
		Disabled:     idle,
	}

	buttonTextColor := &widget.ButtonTextColor{
		Idle: color.RGBA{255, 255, 255, 255},
	}

	return widget.NewButton(
		widget.ButtonOpts.Image(buttonImage),
		widget.ButtonOpts.Text(label, face, buttonTextColor),
		widget.ButtonOpts.TextPadding(widget.Insets{
			Left:  30,
			Right: 30,
		}),
		widget.ButtonOpts.ClickedHandler(onClick),
	), nil
}

func createUI(b *hyper.Board) (*ebitenui.UI, error) {
	root := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout(
			widget.AnchorLayoutOpts.Padding(widget.NewInsetsSimple(10)),
		)),
	)

	ui := &ebitenui.UI{
		Container: root,
	}

	buttonGroup := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Spacing(10),
		)),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionEnd,
				VerticalPosition:   widget.AnchorLayoutPositionStart,
			}),
		),
	)
	root.AddChild(buttonGroup)

	undoButton, err := createButton("Undo", func(args *widget.ButtonClickedEventArgs) {
		b.Undo()
	})
	if err != nil {
		return nil, err
	}
	buttonGroup.AddChild(undoButton)

	redoButton, err := createButton("Redo", func(args *widget.ButtonClickedEventArgs) {
		b.Redo()
	})
	if err != nil {
		return nil, err
	}
	buttonGroup.AddChild(redoButton)

	return ui, nil
}
