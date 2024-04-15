package gizmo

import (
	"fmt"
	"log"
	"math/rand"

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
	Rectangle  rl.RectangleInt32
	PixelRateX float64
	Playing    bool
	drawer     *NumberGrid[T]
}

type GridMover[T check.NumberType] struct {
	GridMoverItem[T]
	Record *model.Record
}

func NewGridMover[T check.NumberType](bounds rl.RectangleInt32, pixelRateX float64) *GridMover[T] {
	mv := &GridMover[T]{}
	var _ model.Parent = mv
	var _ Mover = mv
	var _ Inputer = mv

	mv.Rectangle = bounds
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

func (cm *GridMover[T]) Refresh(now float64, rect rl.RectangleInt32, f ...func(any)) {
	cm.Rectangle = rect
	dr := cm.drawer
	if dr == nil {
		log.Fatalln("nil drawer")
	}
	cm.init(false)
}

func (cm *GridMover[T]) init(clear bool) {

	f := func(t any) {
		cell, ok := t.(*NumberCell[T])
		if ok {
			cell.Clear()
			if rand.Float64() < .1 && !clear {
				cell.Set(Alive, 1)
			}
		}
	}

	cm.drawer.Refresh(0, cm.Rectangle, f)
}

func (cm *GridMover[T]) Move(can_move bool, now float64) {
	if can_move {
		cm.Update()
	}
	cm.drawer.Draw(rl.Vector3{X: float32(cm.Rectangle.X),
		Y: float32(cm.Rectangle.Y),
		Z: 0})
}

func (cm *GridMover[T]) Update() {
	dr := cm.drawer
	cells := dr.GetCells()

	for i := int32(0); i < dr.Cols; i++ {
		for j := int32(0); j < dr.Rows; j++ {
			neighbors := cm.CountNeighbors(i, j)

			cell := cells[i][j]
			// curr :=
			if cell.Get(Alive) == T(1) {
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
			// v := cell.Get(Next)
			// cell.Set(Alive, v)
		}
	}
	for i := int32(0); i <= dr.Cols; i++ {
		for j := int32(0); j < dr.Rows; j++ {
			cell := cells[i][j]
			v := cell.Get(Next)
			cell.Set(Alive, v)
			// cs.cells[i][j].Alive = dr.cells[i][j].Next
		}
	}
}

// INPUT
func (cm *GridMover[T]) Input() {

	if rl.IsKeyPressed(rl.KeyR) {
		fmt.Println("R pressed")
		cm.init(false)
	}
	if rl.IsKeyPressed(rl.KeyC) {
		fmt.Println("R pressed")
		cm.init(true)
	}
	if rl.IsKeyDown(rl.KeyRight) && !cm.Playing {
		fmt.Println("KeyRight pressed", cm.Playing)
		cm.Update()
	}
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		cm.Click(rl.GetMouseX(), rl.GetMouseY())
	}
	if rl.IsKeyPressed(rl.KeySpace) {
		cm.Playing = !cm.Playing
		fmt.Println("cm.Playing", cm.Playing)
	}

}

func (cm *GridMover[T]) Click(clickX, clickY int32) {
	var (
		value, x, y, cx, cy int32
		dr                  = cm.drawer
		cells               = dr.GetCells()
		ix, iy              = dr.PositionToCell(clickX, clickY)
	)

	for cy = iy - 1; cy < iy+2; cy++ {

		for cx = ix - 1; cx < ix+2; cx++ {

			x = (dr.Cols + cx) % dr.Cols
			y = (dr.Rows + cy) % dr.Rows

			cell := cells[x][y]
			value = int32(cell.Get(Alive))
			value = check.As[int32](value == 0)
			cell.Set(Alive, T(value))
			cell.Set(Next, T(value))
		}
	}
}
