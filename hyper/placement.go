package hyper

import (
	"math/rand"
)

type PlacementAlgorithm func(*Board) Point

type Placement struct {
	Actor PlacementAlgorithm
	Goal  PlacementAlgorithm
}

func RandomPlace(b *Board) Point {
	return Point{
		b.rand.Intn(b.Mapdata.Size.W),
		b.rand.Intn(b.Mapdata.Size.H),
	}
}

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
