package hyper

import (
	"fmt"
	"math/rand"
)

type PlacementAlgorithm func(*Board) (Point, error)

type Placement struct {
	Actor PlacementAlgorithm
	Goal  PlacementAlgorithm
}

func PlaceAtRandom(b *Board) (Point, error) {
	for range 50 {
		pos, ok := b.RandomPlace()
		_, exists := b.ActorAt(pos)
		c := b.Mapdata.Center()
		if ok && !exists && !pos.Equals(b.Goal.Point) && !c.Contains(pos) {
			return pos, nil
		}
	}
	var zero Point
	return zero, fmt.Errorf("unable to place goal")
}

func PlaceAtRandomNearByWalls(b *Board) (Point, error) {
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
	for range 50 {
		wall := walls[rand.Intn(len(walls))]
		pos := wall.Add(Point{X: rand.Intn(2) - 1, Y: rand.Intn(2) - 1})
		_, exists := b.ActorAt(pos)
		c := b.Mapdata.Center()
		if !exists && !b.Goal.Point.Equals(pos) && !c.Contains(pos) {
			return pos, nil
		}
	}
	var zero Point
	return zero, fmt.Errorf("unable to place goal")
}
