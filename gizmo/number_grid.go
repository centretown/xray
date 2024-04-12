package gizmo

import (
	"image/color"
	"time"

	"github.com/centretown/xray/check"
	"github.com/centretown/xray/gizmo/categories"
	"github.com/centretown/xray/model"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type NumberGridItem[T check.NumberType] struct {
	Cols       int32
	Rows       int32
	CellWidth  int32
	CellHeight int32
	Start      time.Time
	Duration   time.Duration
	Colors     []color.RGBA

	StateCount int
	cells      [][]*NumberCell[T]
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

type NumberGrid[T check.NumberType] struct {
	NumberGridItem[T]
	Record *model.Record
}

func NewGrid[T check.NumberType](bounds rl.RectangleInt32,
	columns, rows int32,
	colors ...color.RGBA) *NumberGrid[T] {

	cs := &NumberGrid[T]{}
	var _ Drawer = cs

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

	cs.StateCount = len(cs.Colors) - int(CellColorState)
	cs.SetupCells()
	cs.Record = model.NewRecord("generic_grid",
		int32(categories.NumberGridi8), &cs.NumberGridItem, model.JSON)

	return cs
}

func (cs *NumberGrid[T]) GetRecord() *model.Record { return cs.Record }
func (cs *NumberGrid[T]) GetItem() any             { return &cs.NumberGridItem }
func (cs *NumberGrid[T]) Refresh(rect rl.RectangleInt32, options ...bool) {
	if cs.cells == nil {
		cs.SetupCells()
	}
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

func (cs *NumberGrid[T]) SetColors(aliveColor,
	visitedColor,
	gridColor color.RGBA) {
}

func (cs *NumberGrid[T]) SetupCells() {
	cs.cells = make([][]*NumberCell[T], int(cs.Cols+1))
	for x := int32(0); x <= cs.Cols; x++ {
		cs.cells[x] = make([]*NumberCell[T], int(cs.Rows+1))
		for y := int32(0); y <= cs.Rows; y++ {
			cs.cells[x][y] = NewNumberCell[T](cs.StateCount)
		}
	}
}

func (cs *NumberGrid[T]) GetCells() [][]*NumberCell[T] {
	return cs.cells
}

func (cs *NumberGrid[T]) Position(x, y int32) (int32, int32) {
	return cs.bounds.X + x*cs.CellWidth, cs.bounds.Y + y*cs.CellHeight
}

func (cs *NumberGrid[T]) getCell(x, y int32) *NumberCell[T] {
	return cs.cells[x][y]
}

func (cs *NumberGrid[T]) Draw(rl.Vector3) {
	var clr color.RGBA

	for y := range cs.Rows {
		for x := range cs.Cols {
			cell := cs.getCell(x, y)
			clr = Black
			for s := range cell.States {
				clr = cs.Colors[CellColorOff+int(cell.Get(int32(s)))]
			}
			// clr = cs.Colors[CellColorOff+int(cell.Get(x))]
			posX, posY := cs.Position(x, y)
			rl.DrawRectangle(posX, posY, cs.CellWidth, cs.CellHeight, clr)
		}
	}

	for x := range cs.Cols {
		fromX, fromY := cs.Position(x, 0)
		toX, toY := cs.Position(x, cs.Rows)
		rl.DrawLine(fromX, fromY, toX, toY, cs.Colors[CellColorLineVertical])
	}

	for y := range cs.Rows {
		fromX, fromY := cs.Position(0, y)
		toX, toY := cs.Position(cs.Cols, y)
		rl.DrawLine(fromX, fromY, toX, toY, cs.Colors[CellColorLineHorizontal])
	}
}

func (cs *NumberGrid[T]) Bounds() rl.RectangleInt32 {
	return cs.bounds
}
