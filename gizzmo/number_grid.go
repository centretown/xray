package gizzmo

import (
	"fmt"
	"image/color"
	"time"

	rl "github.com/centretown/raylib-go/raylib"
	"github.com/centretown/xray/gizzmo/class"
	"github.com/centretown/xray/gizzmodb/model"
	"github.com/centretown/xray/numbers"
)

type NumberGridItem[T numbers.NumberType] struct {
	Cols            int32
	Rows            int32
	CellWidth       int32
	CellHeight      int32
	Start           time.Time
	Duration        time.Duration
	HorizontalColor color.RGBA
	VerticalColor   color.RGBA
	StateColors     []color.RGBA

	StateCount int
	cells      [][]*NumberCell[T]
	bounds     rl.Vector4
}

type NumberGrid[T numbers.NumberType] struct {
	model.RecorderClass[NumberGridItem[T]]
}

func NewGrid[T numbers.NumberType](bounds rl.Vector4,
	columns, rows int32, horizontalColor color.RGBA, verticalColor color.RGBA,
	colors ...color.RGBA) *NumberGrid[T] {

	cs := &NumberGrid[T]{}
	var _ Drawer = cs

	item := &cs.Content
	item.bounds = bounds
	item.Cols = int32(columns)
	item.Rows = int32(rows)
	item.CellWidth = int32(item.bounds.X / float32(columns))
	item.CellHeight = int32(item.bounds.Y / float32(rows))
	item.HorizontalColor = horizontalColor
	item.VerticalColor = verticalColor
	item.StateColors = colors
	item.StateCount = len(item.StateColors)

	cs.SetupCells()
	model.SetupRecorder[NumberGrid[T]](cs, class.LifeGrid.String(),
		int32(class.LifeGrid))
	return cs
}

func (cs *NumberGrid[T]) Refresh(now float64, v rl.Vector4, funcs ...func(any)) {
	item := &cs.Content
	if item.cells == nil {
		cs.SetupCells()
	}
	item.bounds = v
	item.CellWidth = int32(v.X / float32(item.Cols))
	item.CellHeight = int32(v.Y / float32(item.Rows))

	if len(funcs) < 1 {
		return
	}

	cells := cs.GetCells()
	for y := int32(0); y <= item.Rows; y++ {
		for x := int32(0); x <= item.Cols; x++ {
			for _, f := range funcs {
				f(cells[x][y])
			}
		}
	}
}

func (cs *NumberGrid[T]) SetupCells() {
	fmt.Println("SETUP CELLS")
	item := &cs.Content
	item.cells = make([][]*NumberCell[T], int(item.Cols+1))
	for x := int32(0); x <= item.Cols; x++ {
		item.cells[x] = make([]*NumberCell[T], int(item.Rows+1))
		for y := int32(0); y <= item.Rows; y++ {
			item.cells[x][y] = NewNumberCell[T](item.StateCount)
		}
	}
	fmt.Println("SETUP CELLS", len(item.cells))
}

func (cs *NumberGrid[T]) GetCells() [][]*NumberCell[T] {
	return cs.Content.cells
}

func (cs *NumberGrid[T]) Position(x, y int32) (int32, int32) {
	item := &cs.Content
	return int32(item.bounds.X) + x*item.CellWidth, int32(item.bounds.Y) + y*item.CellHeight
}

func (cs *NumberGrid[T]) PositionToCell(posX, posY int32) (x, y int32) {
	return posX / cs.Content.CellWidth, posY / cs.Content.CellHeight
}

func (cs *NumberGrid[T]) getCell(x, y int32) *NumberCell[T] {
	return cs.Content.cells[x][y]
}

func (cs *NumberGrid[T]) Draw(rl.Vector4) {
	var clr color.RGBA
	item := &cs.Content
	for y := range item.Rows {
		for x := range item.Cols {
			cell := cs.getCell(x, y)
			clr = Black
			for s := range cell.States {
				clr = item.StateColors[int(cell.Get(int32(s)))]
			}
			// clr = cs.Colors[CellColorOff+int(cell.Get(x))]
			posX, posY := cs.Position(x, y)
			rl.DrawRectangle(posX, posY, item.CellWidth, item.CellHeight, clr)
		}
	}

	for x := range item.Cols {
		fromX, fromY := cs.Position(x, 0)
		toX, toY := cs.Position(x, item.Rows)
		rl.DrawLine(fromX, fromY, toX, toY, item.VerticalColor)
	}

	for y := range item.Rows {
		fromX, fromY := cs.Position(0, y)
		toX, toY := cs.Position(item.Cols, y)
		rl.DrawLine(fromX, fromY, toX, toY, item.HorizontalColor)
	}
}

func (cs *NumberGrid[T]) Bounds() rl.Vector4 {
	return cs.Content.bounds
}
