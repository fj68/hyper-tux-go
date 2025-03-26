package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/fj68/hyper-tux-go/hyper"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type GameState struct {
	*hyper.Board
	*SwipeEventDispatcher
	*ResourceLoader
	UI       *ebitenui.UI
	stage    *ebiten.Image
	controls *ebiten.Image
}

func NewGameState(size hyper.Size) (*GameState, error) {
	b, err := hyper.NewBoard(size, hyper.Placement{Actor: hyper.RandomPlace, Goal: hyper.RandomPlaceNearByWalls})
	if err != nil {
		return nil, err
	}
	// debug
	b.Actors[hyper.Black].Point = hyper.Point{X: 0, Y: 0}
	for _, actor := range b.Actors {
		log.Println(actor)
	}

	r := NewResourceLoader()
	ui, err := createUI(r, b)
	if err != nil {
		return nil, err
	}

	stageWidth := b.W * int(CELL_SIZE)
	stageHeight := b.H * int(CELL_SIZE)
	stage := ebiten.NewImage(stageWidth, stageHeight)
	controls := ebiten.NewImage(stageWidth, CONTROLS_HEIGHT)

	return &GameState{
		Board: b,
		SwipeEventDispatcher: NewSwipeEventDispather(
			&MouseEventHandler{},
			&TouchEventHandler{},
		),
		ResourceLoader: r,
		UI:             ui,
		stage:          stage,
		controls:       controls,
	}, nil
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
			return fmt.Errorf("no actor at %+v", e.Start)
		}
		_, ok = g.Board.MoveActor(actor, e.Direction())
		if !ok {
			// TODO: this should not be an error
			return fmt.Errorf("unable to move: %+v to %s", actor, e.Direction())
		}
	}

	return nil
}

func (g *GameState) Update() error {
	if err := g.handleInput(); err != nil {
		return err
	}

	g.UI.Update()

	return nil
}

func (g *GameState) clear(screen *ebiten.Image) {
	screen.Fill(color.White)
}

func (g *GameState) Draw(screen *ebiten.Image) {
	g.clear(screen)

	g.drawStage(g.stage)
	screen.DrawImage(g.stage, &ebiten.DrawImageOptions{})

	g.drawUI(g.controls)
	controlsOp := &ebiten.DrawImageOptions{}
	controlsOp.GeoM.Translate(0, float64(g.stage.Bounds().Dy()))
	screen.DrawImage(g.controls, controlsOp)
}

func (g *GameState) drawStage(screen *ebiten.Image) {
	g.clear(g.stage)
	g.drawBoard(screen)
	g.drawActors(screen)
	g.drawHistory(screen)
	g.drawGoal(screen)
	// bottom border
	vector.StrokeLine(screen, 0, float32(screen.Bounds().Dy()), float32(screen.Bounds().Dx()), float32(screen.Bounds().Dy()), 1, color.Black, false)
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
	vector.StrokeCircle(screen, p.X, p.Y, r, 1, color.Black, true)
}

func (g *GameState) drawHistory(screen *ebiten.Image) {
	for _, record := range g.History() {
		g.drawRecord(screen, record)
	}
}

func (g *GameState) drawRecord(screen *ebiten.Image, record *hyper.Record) {
	lineColor := Color(record.Color)
	start := adjust(Offset(record.Color), record.Start)
	end := adjust(Offset(record.Color), record.End)
	vector.StrokeLine(screen, start.X, start.Y, end.X, end.Y, 1, lineColor, false)
}

func adjust(n float32, p hyper.Point) Position {
	diff := n + CELL_SIZE/2
	pos := NewPosition(p, CELL_SIZE)
	return pos.Add(Position{diff, diff})
}

func (g *GameState) drawGoal(screen *ebiten.Image) {
	goal := g.Board.Goal
	vector.DrawFilledRect(screen, float32(goal.X)*CELL_SIZE, float32(goal.Y)*CELL_SIZE, CELL_SIZE-1, CELL_SIZE-1, Color(goal.Color), false)
}

func (g *GameState) drawUI(screen *ebiten.Image) {
	g.clear(g.controls)
	g.UI.Draw(g.controls)
}

func createUI(r *ResourceLoader, b *hyper.Board) (*ebitenui.UI, error) {
	root := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout(
			widget.AnchorLayoutOpts.Padding(widget.NewInsetsSimple(0)),
		)),
	)

	btnContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionHorizontal),
			widget.RowLayoutOpts.Padding(widget.NewInsetsSimple(8)),
			widget.RowLayoutOpts.Spacing(8),
		)),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			HorizontalPosition: widget.AnchorLayoutPositionCenter,
			VerticalPosition:   widget.AnchorLayoutPositionCenter,
		})),
	)
	root.AddChild(btnContainer)

	undoBtn, err := createButton(r, "Undo", func(args *widget.ButtonClickedEventArgs) {
		b.Undo()
	})
	if err != nil {
		return nil, err
	}
	btnContainer.AddChild(undoBtn)

	redoBtn, err := createButton(r, "Redo", func(args *widget.ButtonClickedEventArgs) {
		b.Redo()
	})
	if err != nil {
		return nil, err
	}
	btnContainer.AddChild(redoBtn)

	resetBtn, err := createButton(r, "Reset", func(args *widget.ButtonClickedEventArgs) {
		b.Reset()
	})
	if err != nil {
		return nil, err
	}
	btnContainer.AddChild(resetBtn)

	newGameBtn, err := createButton(r, "New Game", func(args *widget.ButtonClickedEventArgs) {
		if err := b.NewGame(); err != nil {
			log.Fatal(err)
		}
	})
	if err != nil {
		return nil, err
	}
	btnContainer.AddChild(newGameBtn)

	return &ebitenui.UI{
		Container: root,
	}, nil
}
