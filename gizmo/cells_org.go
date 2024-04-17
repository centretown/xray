package gizmo

import (
	"image/color"
	"math/rand"

	"github.com/centretown/xray/gizmo/class"
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

var (
	gridColor    = rl.LightGray
	aliveColor   = rl.Green
	visitedColor = rl.Beige
)

var _ Drawer = (*CellsOrg)(nil)
var _ Inputer = (*CellsOrg)(nil)

type CellsOrg struct {
	model.RecorderG[CellItemsOrg]
}

func NewCellsOrgFromRecord(record *model.Record) *CellsOrg {
	cs := &CellsOrg{}
	model.Decode(cs, record)
	return cs
}

func NewCellsOrg(width, height, squareSize int32) *CellsOrg {
	cs := &CellsOrg{}
	item := &cs.Content
	item.Height = height
	item.Width = width
	item.SquareSize = squareSize
	item.Cols = width / squareSize
	item.Rows = height / squareSize
	item.Colors = append(item.Colors, gridColor, aliveColor, visitedColor)
	model.InitRecorder[CellItemsOrg](&cs.RecorderG, class.CellsOrg.String(),
		int32(class.CellsOrg))
	cs.start()
	return cs
}

func (cs *CellsOrg) SetColors(aliveColor,
	visitedColor,
	gridColor color.RGBA) {
}

func (cs *CellsOrg) start() {
	item := &cs.Content
	item.cells = make([][]*CellOrg, int(item.Cols+1))

	for x := int32(0); x <= item.Cols; x++ {
		item.cells[x] = make([]*CellOrg, int(item.Rows+1))
		for y := int32(0); y <= item.Rows; y++ {
			item.cells[x][y] = &CellOrg{}
		}
	}

	cs.Init(true)
	item.setup = true
}

func (cs *CellsOrg) Init(clear bool) {
	item := &cs.Content

	for x := int32(0); x <= item.Cols; x++ {
		for y := int32(0); y <= item.Rows; y++ {
			*item.cells[x][y] = CellOrg{}

			item.cells[x][y].Position = rl.NewVector2(float32(x*item.SquareSize),
				float32(y*item.SquareSize+1))
			item.cells[x][y].Size = rl.NewVector2(float32(item.SquareSize-1),
				float32(item.SquareSize-1))

			if rand.Float64() < 0.1 && !clear {
				item.cells[x][y].Alive = true
			}
		}
	}
}

func (cs *CellsOrg) Refresh(float64, rl.Vector4, ...func(any)) {}
func (cs *CellsOrg) Bounds() rl.Rectangle {
	return rl.Rectangle{X: 0, Y: 0, Width: float32(cs.Content.Width), Height: float32(cs.Content.Height)}
}

func (cs *CellsOrg) GetCells() [][]*CellOrg {
	return cs.Content.cells
}

func (cs *CellsOrg) Draw(v rl.Vector4) {
	item := &cs.Content

	if !cs.Content.setup {
		cs.start()
	}

	cs.Update()
	// Draw cells
	for x := int32(0); x <= item.Cols; x++ {
		for y := int32(0); y <= item.Rows; y++ {
			if item.cells[x][y].Alive {
				rl.DrawRectangleV(item.cells[x][y].Position, item.cells[x][y].Size,
					aliveColor)
			} else if item.cells[x][y].Visited {
				rl.DrawRectangleV(item.cells[x][y].Position, item.cells[x][y].Size,
					visitedColor)
			}
		}
	}

	// Draw grid lines
	for i := int32(0); i < item.Cols+1; i++ {
		rl.DrawLineV(
			rl.NewVector2(float32(item.SquareSize*i), 0),
			rl.NewVector2(float32(item.SquareSize*i), float32(item.Height)),
			gridColor)
	}

	for i := int32(0); i < item.Rows+1; i++ {
		rl.DrawLineV(
			rl.NewVector2(0, float32(item.SquareSize*i)),
			rl.NewVector2(float32(item.Width), float32(item.SquareSize*i)),
			gridColor)
	}
}

func (cs *CellsOrg) CountNeighbors(x, y int32) int {
	item := &cs.Content
	count := 0

	for i := int32(-1); i < 2; i++ {
		for j := int32(-1); j < 2; j++ {
			col := (x + i + (item.Cols)) % (item.Cols)
			row := (y + j + (item.Rows)) % (item.Rows)
			if item.cells[col][row].Alive {
				count++
			}
		}
	}

	if item.cells[x][y].Alive {
		count--
	}

	return count
}

func (cs *CellsOrg) Update() {
	item := &cs.Content
	for i := int32(0); i <= item.Cols; i++ {
		for j := int32(0); j <= item.Rows; j++ {
			NeighborCount := cs.CountNeighbors(i, j)
			if item.cells[i][j].Alive {
				if NeighborCount < 2 {
					item.cells[i][j].Next = false
				} else if NeighborCount > 3 {
					item.cells[i][j].Next = false
				} else {
					item.cells[i][j].Next = true
				}
			} else {
				if NeighborCount == 3 {
					item.cells[i][j].Next = true
					item.cells[i][j].Visited = true
				}
			}
		}
	}
	for i := int32(0); i <= item.Cols; i++ {
		for j := int32(0); j < item.Rows; j++ {
			item.cells[i][j].Alive = item.cells[i][j].Next
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
	if rl.IsKeyDown(rl.KeyRight) && !cs.Content.Playing {
		cs.Update()
	}
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		cs.Click(rl.GetMouseX(), rl.GetMouseY())
	}
	if rl.IsKeyPressed(rl.KeySpace) {
		cs.Content.Playing = !cs.Content.Playing
	}

}

func (cs *CellsOrg) Click(x, y int32) {
	item := &cs.Content
	for i := int32(0); i <= item.Cols; i++ {
		for j := int32(0); j <= item.Rows; j++ {
			cell := item.cells[i][j].Position
			if int32(cell.X) < x && int32(cell.X)+item.SquareSize > x &&
				int32(cell.Y) < y && int32(cell.Y)+item.SquareSize > y {

				item.cells[i][j].Alive = !item.cells[i][j].Alive
				item.cells[i][j].Next = item.cells[i][j].Alive
			}
		}
	}
}
