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

	stateCount int
	cells      [][]*Cell[T]
	bounds     rl.RectangleInt32
}

const (
	CellColorLineHorizontal = iota
	CellColorLineVertical
	CellColorOff
	CellColorOn
	CellColorMinimum
	CellColorState = CellColorOn
)

var (
	gridColors = []color.RGBA{
		rl.LightGray,
		rl.LightGray,
		rl.Black,
		rl.Green,
	}
	colorMin = len(gridColors)
)

type Grid[T comparable] struct {
	GridItem[T]
	Record *model.Record
}

func NewGrid[T comparable](bounds rl.RectangleInt32,
	columns, rows int32,
	colors ...color.RGBA) *Grid[T] {

	cs := &Grid[T]{}
	cs.bounds = bounds
	cs.Cols = columns
	cs.Rows = rows
	cs.CellWidth = cs.bounds.Width / columns
	cs.CellHeight = cs.bounds.Height / rows

	l := len(colors)
	if l < colorMin {
		cs.Colors = gridColors[l:]
		cs.Colors = append(cs.Colors, colors...)
	} else {
		cs.Colors = colors
	}
	cs.stateCount = len(cs.Colors) - CellColorState

	cs.SetupCells()
	cs.Record = model.NewRecord("generic_grid",
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
			cs.cells[x][y] = &Cell[T]{
				States: make([]T, cs.stateCount),
			}
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

func (cs *Grid[T]) Draw(rl.Vector3) {
	var clr color.RGBA
	var offValue, current T

	// Draw cells
	for y := int32(0); y < cs.Rows; y++ {
		for x := int32(0); x < cs.Cols; x++ {
			cell := cs.getCell(x, y)
			cell.SetState(0, current)
			clr = cs.Colors[CellColorOff]
			for i := range cs.stateCount {
				current = cell.GetState(i)
				if current != offValue {
					clr = cs.Colors[i+CellColorState]
				}
			}
			posX, posY := cs.Position(x, y)
			rl.DrawRectangle(posX, posY, cs.CellWidth, cs.CellHeight, clr)
		}
	}

	for x := int32(0); x < cs.Cols; x++ {
		fromX, fromY := cs.Position(x, 0)
		toX, toY := cs.Position(x, cs.Rows)
		rl.DrawLine(fromX, fromY, toX, toY, cs.Colors[CellColorLineVertical])
	}

	for y := int32(0); y < cs.Rows; y++ {
		fromX, fromY := cs.Position(0, y)
		toX, toY := cs.Position(cs.Cols, y)
		rl.DrawLine(fromX, fromY, toX, toY, cs.Colors[CellColorLineHorizontal])
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
