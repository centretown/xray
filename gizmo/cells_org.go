package gizmo

import (
	"image/color"
	"math/rand"

	"github.com/centretown/xray/gizmo/categories"
	"github.com/centretown/xray/model"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type CellOrg struct {
	Position rl.Vector2
	Size     rl.Vector2
	Alive    bool
	Next     bool
	Visited  bool
}

type CellItemsOrg struct {
	Cols       int32
	Rows       int32
	Width      int32
	Height     int32
	SquareSize int32
	Playing    bool

	cells  [][]*CellOrg
	Colors []color.RGBA
	setup  bool
}

var _ Drawer = (*CellsOrg)(nil)
var _ Inputer = (*CellsOrg)(nil)

type CellsOrg struct {
	CellItemsOrg
	Record *model.Record
}

func NewCellsOrg(width, height, squareSize int32) *CellsOrg {

	cs := &CellsOrg{}
	cs.Height = height
	cs.Width = width
	cs.SquareSize = squareSize
	cs.Cols = width / squareSize
	cs.Rows = height / squareSize
	cs.Colors = append(cs.Colors, gridColor, aliveColor, visitedColor)

	cs.Record = model.NewRecord("cells",
		int32(categories.CellsOrg), &cs.CellItemsOrg, model.JSON)
	cs.setup = false
	cs.start()
	return cs
}

func (cs *CellsOrg) GetRecord() *model.Record { return cs.Record }
func (cs *CellsOrg) GetItem() any             { return &cs.CellItemsOrg }

func (cs *CellsOrg) SetColors(aliveColor,
	visitedColor,
	gridColor color.RGBA) {
}

func (cs *CellsOrg) start() {
	cs.cells = make([][]*CellOrg, int(cs.Cols+1))
	for x := int32(0); x <= cs.Cols; x++ {
		cs.cells[x] = make([]*CellOrg, int(cs.Rows+1))
		for y := int32(0); y <= cs.Rows; y++ {
			cs.cells[x][y] = &CellOrg{}
		}
	}

	cs.Init(true)
	cs.setup = true
}

func (cs *CellsOrg) Init(clear bool) {
	for x := int32(0); x <= cs.Cols; x++ {
		for y := int32(0); y <= cs.Rows; y++ {
			*cs.cells[x][y] = CellOrg{}

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

func (cs *CellsOrg) Refresh(rect rl.RectangleInt32, options ...bool) {}
func (cs *CellsOrg) Bounds() rl.RectangleInt32 {
	return rl.RectangleInt32{X: 0, Y: 0, Width: cs.Width, Height: cs.Height}
}

func (cs *CellsOrg) GetCells() [][]*CellOrg {
	return cs.cells
}

func (cs *CellsOrg) Draw(v rl.Vector3) {

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

func (cs *CellsOrg) CountNeighbors(x, y int32) int {
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

func (cs *CellsOrg) Update() {
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

func (cs *CellsOrg) Input() {
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

func (cs *CellsOrg) Click(x, y int32) {
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
