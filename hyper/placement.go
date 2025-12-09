package hyper

import (
	"math/rand"
)

// PlacementAlgorithm is a function type that determines where to place an actor or goal on the board.
type PlacementAlgorithm func(*Board) Point

// Placement contains algorithms for placing actors and goals on the board.
type Placement struct {
	Actor PlacementAlgorithm
	Goal  PlacementAlgorithm
}

// RandomPlace returns a random point on the board.
func RandomPlace(b *Board) Point {
	return Point{
		b.rand.Intn(b.Mapdata.Size.W),
		b.rand.Intn(b.Mapdata.Size.H),
	}
}

// RandomPlaceNearByWalls returns a random point adjacent to a wall on the board.
func RandomPlaceNearByWalls(b *Board) Point {
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
	wall := walls[rand.Intn(len(walls))]
	return wall.Add(Point{X: rand.Intn(2) - 1, Y: rand.Intn(2) - 1})
}
