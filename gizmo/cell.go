package gizmo

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/centretown/xray/gizmo/categories"
	"github.com/centretown/xray/model"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var _ Drawer = (*Cells)(nil)
var _ model.Recorder = (*Cells)(nil)

type Cell struct {
	Position rl.Vector2
	Size     rl.Vector2
	Alive    bool
	Next     bool
	Visited  bool
}

type CellItems struct {
	Cols   int32
	Rows   int32
	Width  int32
	Height int32
	cells  [][]*Cell
	setup  bool
}

type Cells struct {
	CellItems
	Record *model.Record
}

const (
	squareSize = 8
)

var _ = rand.New(rand.NewSource(time.Now().UnixNano()))

func NewCells(width, height int32) *Cells {

	cs := &Cells{}
	cs.Height = height
	cs.Width = width
	cs.Cols = width / squareSize
	cs.Rows = height / squareSize

	cs.Record = model.NewRecord("cells",
		int32(categories.Cells), &cs.CellItems, model.JSON)
	cs.setup = false
	cs.start()
	return cs
}

func (cs *Cells) GetRecord() *model.Record { return cs.Record }
func (cs *Cells) GetItem() any             { return &cs.CellItems }

func (cs *Cells) clear(clear bool) {
	for x := int32(0); x <= cs.Cols; x++ {
		for y := int32(0); y <= cs.Rows; y++ {
			cs.cells[x][y] = &Cell{}
			cs.cells[x][y].Position = rl.NewVector2((float32(x) * squareSize), (float32(y)*squareSize)+1)
			cs.cells[x][y].Size = rl.NewVector2(squareSize-1, squareSize-1)
			if rand.Float64() < 0.1 && !clear {
				cs.cells[x][y].Alive = true
			}
		}
	}
}

func (cs *Cells) start() {
	cs.cells = make([][]*Cell, int(cs.Cols+1))
	for i := int32(0); i <= cs.Cols; i++ {
		cs.cells[i] = make([]*Cell, int(cs.Rows+1))
	}

	cs.clear(true)
	cs.setup = true
}

func (cs *Cells) Bounds() rl.RectangleInt32 {
	return rl.RectangleInt32{X: 0, Y: 0, Width: cs.Width, Height: cs.Height}
}

func (cs *Cells) Draw(v rl.Vector3) {

	if !cs.setup {
		fmt.Println("UNINITIALIZED")
		cs.start()
	}
	// Draw cells
	for x := int32(0); x <= cs.Cols; x++ {
		for y := int32(0); y <= cs.Rows; y++ {
			if cs.cells[x][y].Alive {
				rl.DrawRectangleV(cs.cells[x][y].Position, cs.cells[x][y].Size, rl.Blue)
			} else if cs.cells[x][y].Visited {
				rl.DrawRectangleV(cs.cells[x][y].Position, cs.cells[x][y].Size, rl.Color{R: 128, G: 177, B: 136, A: 255})
			}
		}
	}

	// Draw grid lines
	for i := int32(0); i < cs.Cols+1; i++ {
		rl.DrawLineV(
			rl.NewVector2(float32(squareSize*i), 0),
			rl.NewVector2(float32(squareSize*i), float32(cs.Height)),
			rl.LightGray,
		)
	}

	for i := int32(0); i < cs.Rows+1; i++ {
		rl.DrawLineV(
			rl.NewVector2(0, float32(squareSize*i)),
			rl.NewVector2(float32(cs.Width), float32(squareSize*i)),
			rl.LightGray,
		)
	}
}
