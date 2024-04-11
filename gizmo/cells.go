package gizmo

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/centretown/xray/gizmo/categories"
	"github.com/centretown/xray/model"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var _ Drawer = (*Cells)(nil)

type Cell struct {
	Position rl.Vector2
	Size     rl.Vector2
	Alive    bool
	Next     bool
	Visited  bool
}

type CellItems struct {
	Cols       int32
	Rows       int32
	Width      int32
	Height     int32
	SquareSize int32
	Playing    bool

	cells  [][]*Cell
	Colors []color.RGBA
	setup  bool
}

var _ = rand.New(rand.NewSource(time.Now().UnixNano()))

type CellColor int

const (
	CellColorGrid CellColor = iota
	CellColorAlive
	CellColorVisited
)

var (
	aliveColor   = rl.Blue
	visitedColor = rl.Color{R: 128, G: 177, B: 136, A: 255}
	gridColor    = rl.LightGray
)

type Cells struct {
	CellItems
	Record *model.Record
}

func NewCells(width, height, squareSize int32) *Cells {

	cs := &Cells{}
	cs.Height = height
	cs.Width = width
	cs.SquareSize = squareSize
	cs.Cols = width / squareSize
	cs.Rows = height / squareSize
	cs.Colors = append(cs.Colors, gridColor, aliveColor, visitedColor)
	cs.SetupCells()
	cs.Record = model.NewRecord("cells",
		int32(categories.Cells), &cs.CellItems, model.JSON)
	cs.setup = false
	return cs
}

func (cs *Cells) GetRecord() *model.Record { return cs.Record }
func (cs *Cells) GetItem() any             { return &cs.CellItems }

func (cs *Cells) SetColors(aliveColor,
	visitedColor,
	gridColor color.RGBA) {
}

func (cs *Cells) SetupCells() {
	cs.cells = make([][]*Cell, int(cs.Cols+1))
	for x := int32(0); x <= cs.Cols; x++ {
		cs.cells[x] = make([]*Cell, int(cs.Rows+1))
		for y := int32(0); y <= cs.Rows; y++ {
			cs.cells[x][y] = &Cell{}
		}
	}

	cs.setup = true
}

func (cs *Cells) Bounds() rl.RectangleInt32 {
	return rl.RectangleInt32{X: 0, Y: 0, Width: cs.Width, Height: cs.Height}
}

func (cs *Cells) GetCells() [][]*Cell {
	if cs.cells == nil {
		cs.SetupCells()
	}
	return cs.cells
}

func (cs *Cells) Draw(v rl.Vector3) {
	// Draw cells
	for x := int32(0); x <= cs.Cols; x++ {
		for y := int32(0); y <= cs.Rows; y++ {
			if cs.cells[x][y].Alive {
				rl.DrawRectangleV(cs.cells[x][y].Position, cs.cells[x][y].Size,
					cs.Colors[CellColorAlive])
			} else if cs.cells[x][y].Visited {
				rl.DrawRectangleV(cs.cells[x][y].Position, cs.cells[x][y].Size,
					cs.Colors[CellColorVisited])
			}
		}
	}

	// Draw grid lines
	for i := int32(0); i < cs.Cols+1; i++ {
		rl.DrawLineV(
			rl.NewVector2(float32(cs.SquareSize*i), 0),
			rl.NewVector2(float32(cs.SquareSize*i), float32(cs.Height)),
			cs.Colors[CellColorGrid],
		)
	}

	for i := int32(0); i < cs.Rows+1; i++ {
		rl.DrawLineV(
			rl.NewVector2(0, float32(cs.SquareSize*i)),
			rl.NewVector2(float32(cs.Width), float32(cs.SquareSize*i)),
			cs.Colors[CellColorGrid],
		)
	}
}
