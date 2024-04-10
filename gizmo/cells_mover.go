package gizmo

import (
	"log"
	"math/rand"

	"github.com/centretown/xray/model"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type CellMoverItem struct {
	Bounds     rl.RectangleInt32
	PixelRateX float64
	playing    bool
	drawer     *Cells
}

type CellsMover struct {
	CellMoverItem
	Record *model.Record
}

var _ Actor = (*CellsMover)(nil)
var _ Inputer = (*CellsMover)(nil)
var _ model.Recorder = (*CellsMover)(nil)
var _ model.Linker = (*CellsMover)(nil)

func (cm *CellsMover) GetDrawer() Drawer        { return cm.drawer }
func (cm *CellsMover) GetRecord() *model.Record { return cm.Record }
func (cm *CellsMover) GetItem() any             { return &cm.CellMoverItem }

func (cm *CellsMover) Link(recs ...*model.Record) {
	err := MakeLink(cm.AddDrawer, 1, 1, recs...)
	if err != nil {
		log.Fatal(err)
	}
}

func (cm *CellsMover) Children() []model.Recorder {
	return []model.Recorder{cm.drawer}
}

func (cm *CellsMover) AddDrawer(dr *Cells) {
	cm.drawer = dr
}

func (cs *CellsMover) CountNeighbors(x, y int32) int {
	count := 0
	dr := cs.drawer
	cells := dr.GetCells()
	for i := int32(-1); i < 2; i++ {
		for j := int32(-1); j < 2; j++ {
			col := (x + i + (dr.Cols)) % (dr.Cols)
			row := (y + j + (dr.Rows)) % (dr.Rows)
			if cells[col][row].Alive {
				count++
			}
		}
	}

	if cells[x][y].Alive {
		count--
	}

	return count
}
func (cm *CellsMover) Refresh(now float64, rect rl.RectangleInt32) {
	cm.Init(true)
}

func (cs *CellsMover) Act(can_move bool, now float64) {
	if can_move {
		cs.Update()
	}
	cs.drawer.Draw(rl.Vector3{X: float32(cs.Bounds.X),
		Y: float32(cs.Bounds.Y),
		Z: 0})
}

func (cs *CellsMover) Update() {
	dr := cs.drawer
	cells := dr.GetCells()
	for i := int32(0); i <= dr.Cols; i++ {
		for j := int32(0); j <= dr.Rows; j++ {
			NeighborCount := cs.CountNeighbors(i, j)
			if dr.cells[i][j].Alive {
				if NeighborCount < 2 {
					cells[i][j].Next = false
				} else if NeighborCount > 3 {
					cells[i][j].Next = false
				} else {
					cells[i][j].Next = true
				}
			} else {
				if NeighborCount == 3 {
					cells[i][j].Next = true
					cells[i][j].Visited = true
				}
			}
		}
	}
	for i := int32(0); i <= dr.Cols; i++ {
		for j := int32(0); j < dr.Rows; j++ {
			cells[i][j].Alive = cells[i][j].Next
		}
	}
}

func (cs *CellsMover) Init(clear bool) {
	dr := cs.drawer
	cells := dr.GetCells()

	for x := int32(0); x <= dr.Cols; x++ {
		for y := int32(0); y <= dr.Rows; y++ {
			*cells[x][y] = Cell{}

			cells[x][y].Position = rl.NewVector2(float32(x*dr.SquareSize),
				float32(y*dr.SquareSize+1))
			cells[x][y].Size = rl.NewVector2(float32(dr.SquareSize-1),
				float32(dr.SquareSize-1))

			if rand.Float64() < 0.1 && !clear {
				cells[x][y].Alive = true
			}
		}
	}
}

// INPUT
func (cs *CellsMover) Input() {
	// control
	if rl.IsKeyPressed(rl.KeyR) {
		cs.Init(false)
	}
	if rl.IsKeyPressed(rl.KeyC) {
		cs.Init(true)
	}
	if rl.IsKeyDown(rl.KeyRight) && !cs.playing {
		cs.Update()
	}
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		cs.Click(rl.GetMouseX(), rl.GetMouseY())
	}
	if rl.IsKeyPressed(rl.KeySpace) {
		cs.playing = !cs.playing
	}

}

func (cs *CellsMover) Click(x, y int32) {
	dr := cs.drawer
	cells := dr.GetCells()
	for i := int32(0); i <= dr.Cols; i++ {
		for j := int32(0); j <= dr.Rows; j++ {
			cell := cells[i][j].Position
			if int32(cell.X) < x && int32(cell.X)+dr.SquareSize > x &&
				int32(cell.Y) < y && int32(cell.Y)+dr.SquareSize > y {

				cells[i][j].Alive = !cells[i][j].Alive
				cells[i][j].Next = cells[i][j].Alive
			}
		}
	}
}
