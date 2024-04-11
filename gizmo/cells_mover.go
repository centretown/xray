package gizmo

import (
	"log"
	"math/rand"

	"github.com/centretown/xray/gizmo/categories"
	"github.com/centretown/xray/model"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type CellMoverItem struct {
	Bounds     rl.RectangleInt32
	PixelRateX float64
	playing    bool
	drawer     *Cells
}

var _ Actor = (*CellsMover)(nil)
var _ Inputer = (*CellsMover)(nil)
var _ model.Recorder = (*CellsMover)(nil)
var _ model.Linker = (*CellsMover)(nil)

type CellsMover struct {
	CellMoverItem
	Record *model.Record
}

func NewCellsMover(bounds rl.RectangleInt32, pixelRateX float64) *CellsMover {

	mv := &CellsMover{}
	mv.Bounds = bounds
	mv.PixelRateX = pixelRateX
	mv.Record = model.NewRecord("cellsmover", int32(categories.CellsMover), &mv.CellMoverItem, model.JSON)
	return mv
}

func (cm *CellsMover) GetDrawer() Drawer        { return cm.drawer }
func (cm *CellsMover) GetRecord() *model.Record { return cm.Record }
func (cm *CellsMover) GetItem() any             { return &cm.CellMoverItem }

func (cm *CellsMover) Link(recs ...*model.Record) {
	log.Println("MAKELINK", len(recs))
	err := MakeLink(cm.AddDrawer, 1, 1, recs...)
	if err != nil {
		log.Fatal(err)
	}
}

func (cm *CellsMover) Children() []model.Recorder {
	return []model.Recorder{cm.drawer}
}

func (cm *CellsMover) AddDrawer(dr *Cells) {
	if dr == nil {
		panic("AddDrawer")
	}
	log.Println("AddDrawer")
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
	dr := cm.drawer
	if dr == nil {
		log.Fatalln("nil drawer")
	}
	cm.Init(true)
}

func (cm *CellsMover) Act(can_move bool, now float64) {
	if can_move {
		cm.Update()
	}
	cm.drawer.Draw(rl.Vector3{X: float32(cm.Bounds.X),
		Y: float32(cm.Bounds.Y),
		Z: 0})
}

func (cm *CellsMover) Update() {
	dr := cm.drawer
	cells := dr.GetCells()
	for i := int32(0); i <= dr.Cols; i++ {
		for j := int32(0); j <= dr.Rows; j++ {
			NeighborCount := cm.CountNeighbors(i, j)
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

func (cm *CellsMover) Init(clear bool) {
	dr := cm.drawer
	if dr == nil {
		log.Fatalln("nil drawer")
	}
	cells := dr.GetCells()
	if cells == nil {
		log.Fatalln("nil cells")
	}

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
func (cm *CellsMover) Input() {
	// control
	if rl.IsKeyPressed(rl.KeyR) {
		cm.Init(false)
	}
	if rl.IsKeyPressed(rl.KeyC) {
		cm.Init(true)
	}
	if rl.IsKeyDown(rl.KeyRight) && !cm.playing {
		cm.Update()
	}
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		cm.Click(rl.GetMouseX(), rl.GetMouseY())
	}
	if rl.IsKeyPressed(rl.KeySpace) {
		cm.playing = !cm.playing
	}

}

func (cm *CellsMover) Click(x, y int32) {
	dr := cm.drawer
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
