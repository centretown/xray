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
var _ Inputer = (*Cells)(nil)

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

	// AliveColor   color.RGBA
	// VisitedColor color.RGBA
	// GridColor    color.RGBA

	cells  [][]*Cell
	Colors []color.RGBA
	setup  bool
}

type Cells struct {
	CellItems
	Record *model.Record
}

const (
// squareSize = 8
)

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

func NewCells(width, height, squareSize int32) *Cells {

	cs := &Cells{}
	cs.Height = height
	cs.Width = width
	cs.SquareSize = squareSize
	cs.Cols = width / squareSize
	cs.Rows = height / squareSize
	cs.Colors = append(cs.Colors, gridColor, aliveColor, visitedColor)

	cs.Record = model.NewRecord("cells",
		int32(categories.Cells), &cs.CellItems, model.JSON)
	cs.setup = false
	cs.start()
	return cs
}

func (cs *Cells) GetRecord() *model.Record { return cs.Record }
func (cs *Cells) GetItem() any             { return &cs.CellItems }

func (cs *Cells) SetColors(aliveColor,
	visitedColor,
	gridColor color.RGBA) {
}

func (cs *Cells) start() {
	cs.cells = make([][]*Cell, int(cs.Cols+1))
	for i := int32(0); i <= cs.Cols; i++ {
		cs.cells[i] = make([]*Cell, int(cs.Rows+1))
	}

	cs.Init(true)
	cs.setup = true
}

func (cs *Cells) Init(clear bool) {
	for x := int32(0); x <= cs.Cols; x++ {
		for y := int32(0); y <= cs.Rows; y++ {
			cs.cells[x][y] = &Cell{}

			cs.cells[x][y].Position = rl.NewVector2(float32(x*cs.SquareSize),
				float32(y*cs.SquareSize+1))
			cs.cells[x][y].Size = rl.NewVector2(float32(cs.SquareSize-1),
				float32(cs.SquareSize-1))

			if rand.Float64() < 0.1 && !clear {
				cs.cells[x][y].Alive = true
			}
		}
	}
}

func (cs *Cells) Bounds() rl.RectangleInt32 {
	return rl.RectangleInt32{X: 0, Y: 0, Width: cs.Width, Height: cs.Height}
}

func (cs *Cells) Draw(v rl.Vector3) {

	if !cs.setup {
		cs.start()
	}

	cs.Update()
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

func (cs *Cells) CountNeighbors(x, y int32) int {
	count := 0

	for i := int32(-1); i < 2; i++ {
		for j := int32(-1); j < 2; j++ {
			col := (x + i + (cs.Cols)) % (cs.Cols)
			row := (y + j + (cs.Rows)) % (cs.Rows)
			if cs.cells[col][row].Alive {
				count++
			}
		}
	}

	if cs.cells[x][y].Alive {
		count--
	}

	return count
}

func (cs *Cells) Update() {
	for i := int32(0); i <= cs.Cols; i++ {
		for j := int32(0); j <= cs.Rows; j++ {
			NeighborCount := cs.CountNeighbors(i, j)
			if cs.cells[i][j].Alive {
				if NeighborCount < 2 {
					cs.cells[i][j].Next = false
				} else if NeighborCount > 3 {
					cs.cells[i][j].Next = false
				} else {
					cs.cells[i][j].Next = true
				}
			} else {
				if NeighborCount == 3 {
					cs.cells[i][j].Next = true
					cs.cells[i][j].Visited = true
				}
			}
		}
	}
	for i := int32(0); i <= cs.Cols; i++ {
		for j := int32(0); j < cs.Rows; j++ {
			cs.cells[i][j].Alive = cs.cells[i][j].Next
		}
	}
}

func (cs *Cells) Input() {
	// control
	if rl.IsKeyPressed(rl.KeyR) {
		cs.Init(false)
	}
	if rl.IsKeyPressed(rl.KeyC) {
		cs.Init(true)
	}
	if rl.IsKeyDown(rl.KeyRight) && !cs.Playing {
		cs.Update()
	}
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		cs.Click(rl.GetMouseX(), rl.GetMouseY())
	}
	if rl.IsKeyPressed(rl.KeySpace) {
		cs.Playing = !cs.Playing
	}

}

func (cs *Cells) Click(x, y int32) {
	for i := int32(0); i <= cs.Cols; i++ {
		for j := int32(0); j <= cs.Rows; j++ {
			cell := cs.cells[i][j].Position
			if int32(cell.X) < x && int32(cell.X)+cs.SquareSize > x &&
				int32(cell.Y) < y && int32(cell.Y)+cs.SquareSize > y {

				cs.cells[i][j].Alive = !cs.cells[i][j].Alive
				cs.cells[i][j].Next = cs.cells[i][j].Alive
			}
		}
	}
}
