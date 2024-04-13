package gizmo

import (
	"log"

	"github.com/centretown/xray/check"
	"github.com/centretown/xray/gizmo/categories"
	"github.com/centretown/xray/model"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	Alive int32 = iota
	Visited
	Next
	GRIDMOVER_COUNT
)

type GridMoverItem[T check.NumberType] struct {
	Bounds     rl.RectangleInt32
	PixelRateX float64
	playing    bool
	drawer     *NumberGrid[T]
}

type GridMover[T check.NumberType] struct {
	GridMoverItem[T]
	Record *model.Record
}

func NewGridMover[T check.NumberType](bounds rl.RectangleInt32, pixelRateX float64) *GridMover[T] {
	mv := &GridMover[T]{}
	var _ model.Parent = mv
	var _ Actor = mv
	// var _ Inputer = mv
	mv.Bounds = bounds
	mv.PixelRateX = pixelRateX
	mv.Record = model.NewRecord("cellsmover", int32(categories.NumberMoveri8), &mv.GridMoverItem, model.JSON)
	return mv
}

func (cm *GridMover[T]) GetDrawer() Drawer        { return cm.drawer }
func (cm *GridMover[T]) GetRecord() *model.Record { return cm.Record }
func (cm *GridMover[T]) GetItem() any             { return &cm.GridMoverItem }

func (cm *GridMover[T]) LinkChildren(recs ...*model.Record) {
	// log.Println("MAKELINK", len(recs))
	err := MakeLink(cm.AddDrawer, 1, 1, recs...)
	if err != nil {
		log.Fatal(err)
	}
}

func (cm *GridMover[T]) Children() []model.Recorder {
	return []model.Recorder{cm.drawer}
}

func (cm *GridMover[T]) AddDrawer(dr *NumberGrid[T]) {
	cm.drawer = dr
}

func doMod(a, b int32) int32 {
	return (b + a) % b
}

func (cs *GridMover[T]) CountNeighbors(cellX, cellY int32) int {
	count := 0
	dr := cs.drawer
	cells := dr.GetCells()

	//-1..1
	for y := int32(-1); y < 2; y++ {
		//-1..1
		for x := int32(-1); x < 2; x++ {
			nx, ny := doMod(cellX+x, dr.Cols), doMod(cellY+y, dr.Rows)
			neigh := cells[nx][ny]
			if neigh.Get(Alive) != T(0) {
				count++
			}
		}
	}

	if cells[cellX][cellY].Get(Alive) != T(0) {
		count--
	}

	return count
}

func (cm *GridMover[T]) Refresh(now float64, rect rl.RectangleInt32) {
	cm.Bounds = rect
	dr := cm.drawer
	if dr == nil {
		log.Fatalln("nil drawer")
	}
	cm.drawer.Refresh(rect, false)
}

func (cm *GridMover[T]) Act(can_move bool, now float64) {
	// if can_move {
	// 	cm.Update()
	// }
	cm.drawer.Draw(rl.Vector3{X: float32(cm.Bounds.X),
		Y: float32(cm.Bounds.Y),
		Z: 0})
}

func (cm *GridMover[T]) Update() {
	dr := cm.drawer
	cells := dr.GetCells()
	for i := int32(0); i <= dr.Cols; i++ {
		for j := int32(0); j <= dr.Rows; j++ {
			neighbors := cm.CountNeighbors(i, j)
			cell := cells[i][j]
			curr := cell.Get(Alive)
			if curr == T(0) {
				if neighbors < 2 {
					cell.Set(Next, 0)
				} else if neighbors > 3 {
					cell.Set(Next, 0)
				} else {
					cell.Set(Next, 1)
				}
			} else if neighbors == 3 {
				cell.Set(Next, 1)
				cell.Set(Visited, 1)
			}
		}
	}

	for i := int32(0); i < dr.Cols; i++ {
		for j := int32(0); j < dr.Rows; j++ {
			cell := cells[i][j]
			cell.Set(Alive, cell.Get(Next))
		}
	}
}

// INPUT
func (cm *GridMover[T]) Input() {
	// control
	if rl.IsKeyPressed(rl.KeyR) {
		cm.drawer.Refresh(cm.Bounds, false)
	}
	if rl.IsKeyPressed(rl.KeyC) {
		cm.drawer.Refresh(cm.Bounds, true)
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

func (cm *GridMover[T]) Click(clickX, clickY int32) {
	// dr := cm.drawer
	// cells := dr.GetCells()
	// cell := dr.PositionToCell(clickX, clickY)
	for y := int32(-1); y < 2; y++ {

		for x := int32(-1); x < 2; x++ {

			// cell := cells[x][y]

			// if int32(cell.X) < clickX && int32(cell.X)+dr.CellWidth > clickX &&
			// 	int32(cell.Y) < clickY && int32(cell.Y)+dr.CellHeight > clickY {

			// 	cells[x][y].Alive = !cells[x][y].Alive
			// 	cells[x][y].Next = cells[x][y].Alive
			// }
		}
	}
}
