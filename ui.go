package main

import (
	"bytes"
	"fmt"
	"image/color"
	"log"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/fj68/hyper-tux-go/hyper"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/gofont/goregular"
)

type Labels struct {
	Steps  *widget.Label
	Goaled *widget.Label
}

func (ls *Labels) Update(b *hyper.Board) {
	ls.Steps.Label = fmt.Sprintf("Steps: %d", b.Steps())
	if b.Goaled {
		ls.Goaled.Label = "Goal"
	} else {
		ls.Goaled.Label = ""
	}
}

func createButton(label string, onClick widget.ButtonClickedHandlerFunc) (*widget.Button, error) {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		return nil, err
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

	textColor := &widget.ButtonTextColor{
		Idle: color.RGBA{255, 255, 255, 255},
	}

	return widget.NewButton(
		widget.ButtonOpts.Image(buttonImage),
		widget.ButtonOpts.Text(label, face, textColor),
		widget.ButtonOpts.TextPadding(widget.Insets{
			Left:  30,
			Right: 30,
		}),
		widget.ButtonOpts.ClickedHandler(onClick),
	), nil
}

func createLabel(content string) (*widget.Label, error) {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		return nil, err
	}

	face := &text.GoTextFace{
		Source: s,
		Size:   32,
	}

	textColor := &widget.LabelColor{
		Idle: color.RGBA{255, 255, 255, 255},
	}

	label := widget.NewLabel(
		widget.LabelOpts.Text(content, face, textColor),
	)

	return label, nil
}

func createUI(b *hyper.Board) (*ebitenui.UI, *Labels, error) {
	root := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout(
			widget.AnchorLayoutOpts.Padding(widget.NewInsetsSimple(10)),
		)),
	)

	ui := &ebitenui.UI{
		Container: root,
	}

	buttonGroup, err := createButtonGroup(b)
	if err != nil {
		return nil, nil, err
	}
	root.AddChild(buttonGroup)

	stepsLabel, err := createLabel("0")
	if err != nil {
		return nil, nil, err
	}
	buttonGroup.AddChild(stepsLabel)

	goaledLabel, err := createLabel("")
	if err != nil {
		return nil, nil, err
	}
	buttonGroup.AddChild(goaledLabel)

	labels := &Labels{
		Steps:  stepsLabel,
		Goaled: goaledLabel,
	}

	return ui, labels, nil
}

func createButtonGroup(b *hyper.Board) (*widget.Container, error) {
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

	resetButton, err := createButton("Reset", func(args *widget.ButtonClickedEventArgs) {
		b.Reset()
	})
	if err != nil {
		return nil, err
	}
	buttonGroup.AddChild(resetButton)

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

	newGameButton, err := createButton("New Game", func(args *widget.ButtonClickedEventArgs) {
		if err := b.NewGame(); err != nil {
			log.Println(err)
		}
	})
	if err != nil {
		return nil, err
	}
	buttonGroup.AddChild(newGameButton)

	return buttonGroup, nil
}
