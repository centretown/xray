package gizmo

import (
	"image/color"
	"time"

	"github.com/centretown/xray/gizmo/categories"
	"github.com/centretown/xray/model"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var _ Drawer = (*Grid[bool])(nil)

type GridItem[T any] struct {
	Cols       int32
	Rows       int32
	CellWidth  int32
	CellHeight int32
	Start      time.Time
	Duration   time.Duration
	Playing    bool
	Colors     []color.RGBA

	cells  [][]*Cell[T]
	bounds rl.RectangleInt32
}

type CellColor int

const (
	CellColorAlive CellColor = iota
	CellColorVisited
	CellColorGrid
)

var (
	aliveColor   = rl.Blue
	visitedColor = rl.Color{R: 128, G: 177, B: 136, A: 255}
	gridColor    = rl.LightGray
)

type Grid[T any] struct {
	GridItem[T]
	Record *model.Record
}

func NewGrid[T any](rect rl.RectangleInt32, columns, rows int32) *Grid[T] {
	cs := &Grid[T]{}
	cs.bounds = rect
	cs.Cols = columns
	cs.Rows = rows
	cs.CellWidth = cs.bounds.Width / columns
	cs.CellHeight = cs.bounds.Height / rows
	cs.Colors = append(cs.Colors, gridColor, aliveColor, visitedColor)
	cs.SetupCells()
	cs.Record = model.NewRecord("cells",
		int32(categories.Grid), &cs.GridItem, model.JSON)
	return cs
}

func (cs *Grid[T]) GetRecord() *model.Record { return cs.Record }
func (cs *Grid[T]) GetItem() any             { return &cs.GridItem }

func (cs *Grid[T]) SetColors(aliveColor,
	visitedColor,
	gridColor color.RGBA) {
}

func (cs *Grid[T]) SetupCells() {
	cs.cells = make([][]*Cell[T], int(cs.Cols+1))
	for x := int32(0); x <= cs.Cols; x++ {
		cs.cells[x] = make([]*Cell[T], int(cs.Rows+1))
		for y := int32(0); y <= cs.Rows; y++ {
			cs.cells[x][y] = &Cell[T]{}
		}
	}
}

func (cs *Grid[T]) GetCells() [][]*Cell[T] {
	if cs.cells == nil {
		cs.SetupCells()
	}
	return cs.cells
}

func (cs *Grid[T]) Position(x, y int32) (int32, int32) {
	return cs.bounds.X + x*cs.CellWidth, cs.bounds.Y + y*cs.CellHeight
}

func (cs *Grid[T]) getCell(x, y int32) *Cell[T] {
	return cs.cells[x][y]
}

func (cs *Grid[T]) Draw(v rl.Vector3) {
	var clr color.RGBA
	// Draw cells
	for y := int32(0); y < cs.Rows; y++ {
		for x := int32(0); x < cs.Cols; x++ {
			// cell := cs.getCell(x, y)
			alive, visited := false, false //cell.Alive, cell.Visited
			if alive {
				clr = cs.Colors[CellColorAlive]
			} else if visited {
				clr = cs.Colors[CellColorVisited]
			}
			if alive || visited {
				posX, posY := cs.Position(x, y)
				rl.DrawRectangle(posX, posY, cs.CellWidth, cs.CellHeight, clr)
			}
		}
	}

	for x := int32(0); x < cs.Cols; x++ {
		fromX, fromY := cs.Position(x, 0)
		toX, toY := cs.Position(x, cs.Rows)
		rl.DrawLine(fromX, fromY, toX, toY, cs.Colors[CellColorGrid])
	}

	for y := int32(0); y < cs.Rows; y++ {
		fromX, fromY := cs.Position(0, y)
		toX, toY := cs.Position(cs.Cols, y)
		rl.DrawLine(fromX, fromY, toX, toY, cs.Colors[CellColorGrid])
	}
}

func (cs *Grid[T]) Bounds() rl.RectangleInt32 {
	return cs.bounds
}

func (cs *Grid[T]) Refresh(rect rl.RectangleInt32, options ...bool) {
	cs.bounds = rect
	cs.CellWidth = rect.Width / cs.Cols
	cs.CellHeight = rect.Height / cs.Rows

	if len(options) == 0 || !options[0] {
		return
	}

	cells := cs.GetCells()
	for y := int32(0); y <= cs.Rows; y++ {
		for x := int32(0); x <= cs.Cols; x++ {
			cells[x][y].Clear()
		}
	}
}
