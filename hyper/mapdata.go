package hyper

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"

	"github.com/fj68/hyper-tux-go/internal/slicetools"
	"golang.org/x/exp/slices"
)

// walls
// - - - - - -
// -|-|-|-|-|-
// -|-|-|-|-|-
// - - - - - -

type Mapdata struct {
	Size
	HWalls [][]int
	VWalls [][]int
}

func NewMapdata(size Size) *Mapdata {
	HWalls := make([][]int, size.W)
	VWalls := make([][]int, size.H)

	m := &Mapdata{size, HWalls, VWalls}

	// place center walls
	m.initCenterWalls()

	return m
}

// each cell represents wall by int of (0|North|West)
func NewMapdataFromCSV(r *csv.Reader) (*Mapdata, error) {
	values := [][]int{}

	for y := 0; ; y++ {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		values[y] = []int{}

		for x, field := range row {
			n, err := strconv.Atoi(field)
			if err != nil {
				return nil, fmt.Errorf("error at L%dC%d: %w", y, x, err)
			}
			values[y][x] = n
		}
	}

	return NewMapdataFromSlice(values)
}

func NewMapdataFromSlice(rows [][]int) (*Mapdata, error) {
	if len(rows) < 1 {
		return NewMapdata(Size{0, 0}), nil
	}

	m := NewMapdata(Size{W: len(rows[0]), H: len(rows)})

	for y, row := range rows {
		for x, n := range row {
			if (n & int(North)) != 0 {
				m.PutHWall(Point{x, y})
			}
			if (n & int(West)) != 0 {
				m.PutVWall(Point{x, y})
			}
		}
	}

	// place center walls
	m.initCenterWalls()

	return m, nil
}

func (m *Mapdata) PutHWall(p Point) {
	if slices.Contains(m.HWalls[p.X], p.Y) {
		return
	}
	m.HWalls[p.X] = append(m.HWalls[p.X], p.Y)
}

func (m *Mapdata) PutVWall(p Point) {
	if slices.Contains(m.VWalls[p.Y], p.X) {
		return
	}
	m.VWalls[p.Y] = append(m.VWalls[p.Y], p.X)
}

func (m *Mapdata) Center() Rect {
	c := m.Size.Center()
	return NewRect(Point{c.X - 1, c.Y - 1}, Size{2, 2})
}

func (m *Mapdata) initCenterWalls() {
	r := m.Center()

	m.PutHWall(Point{r.TopLeft.X, r.TopLeft.Y})
	m.PutVWall(Point{r.TopLeft.X, r.TopLeft.Y})
	m.PutHWall(Point{r.TopLeft.X, r.BottomRight.Y})
	m.PutVWall(Point{r.TopLeft.X, r.BottomRight.Y - 1})
	m.PutHWall(Point{r.BottomRight.X - 1, r.TopLeft.Y})
	m.PutVWall(Point{r.BottomRight.X, r.TopLeft.Y})
	m.PutHWall(Point{r.BottomRight.X - 1, r.BottomRight.Y})
	m.PutVWall(Point{r.BottomRight.X, r.BottomRight.Y - 1})
}

func (m *Mapdata) Equals(other *Mapdata) bool {
	intSliceEquals := func(a, b []int) bool {
		return slicetools.Equals(a, b)
	}

	if !m.Size.Equals(other.Size) {
		return false
	}
	if !slicetools.EqualsFunc(m.HWalls, other.HWalls, intSliceEquals) {
		return false
	}
	if !slicetools.EqualsFunc(m.VWalls, other.VWalls, intSliceEquals) {
		return false
	}

	return true
}
