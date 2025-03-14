package hyper

import (
	"encoding/csv"
	"fmt"
	"strconv"
)

type VWall = []int // | : list of column indices where the wall exists
type HWall = []int // _ : list of row indices where the wall exists

type Mapdata struct {
	*Size
	HWalls []HWall
	VWalls []VWall
}

func NewMapdata(size *Size) *Mapdata {
	m := &Mapdata{
		Size:   size,
		VWalls: make([]VWall, size.H),
		HWalls: make([]HWall, size.W),
	}

	// place center walls
	m.initCenterWalls()

	return m
}

// each cell represents wall by int of (0|North|West)
func NewMapdataFromCSV(r *csv.Reader) (*Mapdata, error) {
	rows, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	return NewMapdataFromSlice(rows)
}

func NewMapdataFromSlice(rows [][]string) (*Mapdata, error) {
	if len(rows) < 1 {
		return NewMapdata(&Size{0, 0}), nil
	}

	size := &Size{W: len(rows[0]), H: len(rows)}
	m := NewMapdata(size)

	for y, fields := range rows {
		for x, field := range fields {
			n, err := strconv.Atoi(field)
			if err != nil {
				return nil, fmt.Errorf("error at L%dC%d: %w", y, x, err)
			}

			if (n & int(North)) != 0 {
				m.PutHWall(&Point{x, y})
			}
			if (n & int(West)) != 0 {
				m.PutVWall(&Point{x, y})
			}
		}
	}

	return m, nil
}

func (m *Mapdata) PutHWall(p *Point) bool {
	for _, wall := range m.HWalls[p.X] {
		if wall == p.Y {
			return false
		}
	}
	m.HWalls[p.X] = append(m.HWalls[p.X], p.Y)
	return true
}

func (m *Mapdata) PutVWall(p *Point) bool {
	for _, wall := range m.VWalls[p.Y] {
		if wall == p.X {
			return false
		}
	}
	m.VWalls[p.Y] = append(m.VWalls[p.Y], p.X)
	return true
}

func (m *Mapdata) Center() *Rect {
	c := m.Size.Center()
	return NewRect(&Point{c.X - 1, c.Y - 1}, &Size{2, 2})
}

func (m *Mapdata) initCenterWalls() {
	r := m.Center()
	s := r.Size()

	for i := range s.W {
		m.HWalls[i+r.TopLeft.X] = []int{r.TopLeft.Y, r.BottomRight.Y - 1}
	}
	for i := range s.H {
		m.VWalls[i+r.TopLeft.Y] = []int{r.TopLeft.X, r.BottomRight.X - 1}
	}
}
