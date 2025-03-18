package main

import (
	"bytes"
	"fmt"
	"image/color"
	"log"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/fj68/hyper-tux-go/hyper"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/gofont/goregular"
)

type GameState struct {
	*hyper.Board
	*SwipeEventDispatcher
	*ebitenui.UI
}

func NewGameState(size hyper.Size) (*GameState, error) {
	board, err := hyper.NewBoard(size)
	if err != nil {
		return nil, err
	}
	// debug
	board.Actors[hyper.Black].Point = hyper.Point{X: 0, Y: 0}
	for _, actor := range board.Actors {
		log.Println(actor)
	}

	ui, err := createUI(board)
	if err != nil {
		return nil, err
	}

	return &GameState{
		Board: board,
		SwipeEventDispatcher: NewSwipeEventDispather(
			&MouseEventHandler{},
			&TouchEventHandler{},
		),
		UI: ui,
	}, nil
}

func newImageFromFile(path string) (*ebiten.Image, error) {
	f, err := embeddedAssets.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	i, _, err := ebitenutil.NewImageFromReader(f)
	return i, err
}

func loadImageNineSlice(path string, centerWidth int, centerHeight int) (*image.NineSlice, error) {
	i, err := newImageFromFile(path)
	if err != nil {
		return nil, err
	}
	w := i.Bounds().Dx()
	h := i.Bounds().Dy()
	return image.NewNineSlice(i,
			[3]int{(w - centerWidth) / 2, centerWidth, w - (w-centerWidth)/2 - centerWidth},
			[3]int{(h - centerHeight) / 2, centerHeight, h - (h-centerHeight)/2 - centerHeight}),
		nil
}

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

func (g *GameState) handleInput() error {
	if err := g.SwipeEventDispatcher.Update(); err != nil {
		return err
	}

	for g.SwipeEventDispatcher.Len() > 0 {
		e := g.SwipeEventDispatcher.Pop()
		if e == nil {
			return fmt.Errorf("SwipeEvent is nil")
		}
		actor, ok := g.Board.ActorAt(e.Start)
		if !ok {
			// TODO: this should not be an error
			return fmt.Errorf("No actor at %+v", e.Start)
		}
		_, ok = g.Board.MoveActor(actor, e.Direction())
		if !ok {
			// TODO: this should not be an error
			return fmt.Errorf("Unable to move: %+v to %s", actor, e.Direction())
		}
	}

	return nil
}

func (g *GameState) Update() error {
	g.UI.Update()

	if err := g.handleInput(); err != nil {
		return err
	}

	return nil
}

func (g *GameState) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen, 0, 0, float32(g.W)*CELL_SIZE, float32(g.H)*CELL_SIZE, color.White, false)
	g.drawBoard(screen)
	g.drawActors(screen)
	g.UI.Draw(screen)
}

func (g *GameState) drawBoard(screen *ebiten.Image) {
	lineColor := color.Gray{200}
	// lines
	for y := range g.Board.H - 1 {
		vector.StrokeLine(screen, 0, float32(y+1)*CELL_SIZE, float32(g.Board.W)*CELL_SIZE, float32(y+1)*CELL_SIZE, 1, lineColor, false)
	}
	for x := range g.Board.W - 1 {
		vector.StrokeLine(screen, float32(x+1)*CELL_SIZE, 0, float32(x+1)*CELL_SIZE, float32(g.Board.H)*CELL_SIZE, 1, lineColor, false)
	}
	// walls
	for y, rows := range g.Board.VWalls {
		for _, x := range rows {
			vector.StrokeLine(screen, float32(x)*CELL_SIZE, float32(y)*CELL_SIZE, float32(x)*CELL_SIZE, float32(y+1)*CELL_SIZE, 1, color.Black, false)
		}
	}
	for x, cols := range g.Board.HWalls {
		for _, y := range cols {
			vector.StrokeLine(screen, float32(x)*CELL_SIZE, float32(y)*CELL_SIZE, float32(x+1)*CELL_SIZE, float32(y)*CELL_SIZE, 1, color.Black, false)
		}
	}
	// center box
	c := g.Board.Center()
	vector.DrawFilledRect(screen, float32(c.TopLeft.X)*CELL_SIZE, float32(c.TopLeft.Y)*CELL_SIZE, float32(c.Size().W)*CELL_SIZE-1, float32(c.Size().H)*CELL_SIZE-1, lineColor, false)
}

func (g *GameState) drawActors(screen *ebiten.Image) {
	for _, actor := range g.Board.Actors {
		g.drawActor(screen, actor)
	}
}

func (g *GameState) drawActor(screen *ebiten.Image, actor *hyper.Actor) {
	p := NewPosition(actor.Point, CELL_SIZE)
	halfCellSize := CELL_SIZE / 2
	p = p.Add(Position{halfCellSize, halfCellSize})
	r := halfCellSize - 2
	vector.DrawFilledCircle(screen, p.X, p.Y, r, Color(actor.Color), true)
}
