package hyper

import (
	"maps"
	"math/rand"
)

// PlacementAlgorithm defines a function type for placing actors and goals on the board.
type ActorPlacementAlgorithm func(*Board, Color) (Point, bool)

// GoalPlacementAlgorithm defines a function type for placing goals on the board.
type GoalPlacementAlgorithm func(*Board) (Goal, bool)

// Placement contains algorithms for placing actors and goals on the board.
type Placement struct {
	Actor ActorPlacementAlgorithm
	Goal  GoalPlacementAlgorithm
}

// PlaceActorAt returns an ActorPlacementAlgorithm that places actors at specified points.
func PlaceActorAt(ps... map[Color]Point) ActorPlacementAlgorithm {
	p := map[Color]Point{}
	for _, placement := range ps {
		maps.Copy(p, placement)
	}
	return func(b *Board, c Color) (Point, bool) {
		pos, ok := p[c]
		return pos, ok
	}
}

// PlaceGoalAt returns a GoalPlacementAlgorithm that places the goal at a specified point and color.
func PlaceGoalAt(color Color, pos Point) GoalPlacementAlgorithm {
	return func(b *Board) (Goal, bool) {
		return Goal{color, pos}, true
	}
}

// PlaceAtRandom returns a random point on the board.
func PlaceAtRandom(b *Board) Point {
	return Point{
		b.rand.Intn(b.Mapdata.Size.W),
		b.rand.Intn(b.Mapdata.Size.H),
	}
}

// PlaceActorAtRandom returns a random point on the board.
func PlaceActorAtRandom(b *Board, _ Color) (Point, bool) {
	return PlaceAtRandom(b), true
}

// PlaceGoalAtRandom returns a random point and color for the goal.
func PlaceGoalAtRandom(b *Board) (Goal, bool) {
	return Goal{ColorAtRandom(), PlaceAtRandom(b)}, true
}

// PlaceNearByWalls returns a point near existing walls on the board.
func PlaceNearByWalls(b *Board) (Point, bool) {
	walls := []Point{}
	for y, row := range b.Mapdata.HWalls {
		for _, x := range row {
			walls = append(walls, Point{x, y})
		}
	}
	for y, row := range b.Mapdata.VWalls {
		for _, x := range row {
			walls = append(walls, Point{x, y})
		}
	}
	if len(walls) == 0 {
		return Point{}, false
	}
	wall := walls[rand.Intn(len(walls))]
	return wall.Add(Point{X: rand.Intn(2) - 1, Y: rand.Intn(2) - 1}), true
}

// PlaceActorNearByWalls returns a point near existing walls on the board.
func PlaceActorNearByWalls(b *Board, _ Color) (Point, bool) {
	return PlaceNearByWalls(b)
}

// PlaceGoalNearByWalls returns a point near existing walls on the board along with a random color.
func PlaceGoalNearByWalls(b *Board) (Goal, bool) {
	if pos, ok := PlaceNearByWalls(b); ok {
		return Goal{ColorAtRandom(), pos}, true
	}
	return b.Goal, false
}
