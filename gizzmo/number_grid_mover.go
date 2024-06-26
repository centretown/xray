package gizzmo

import (
	"fmt"
	"log"
	"math/rand"

	rl "github.com/centretown/raylib-go/raylib"
	"github.com/centretown/xray/class"
	"github.com/centretown/xray/gizzmodb/model"
	"github.com/centretown/xray/numbers"
)

const (
	Alive int32 = iota
	Visited
	Next
	GRIDMOVER_COUNT
)

type GridMoverItem[T numbers.NumberType] struct {
	Bounds     rl.Vector4
	PixelRateX float64
	Playing    bool
	drawer     *NumberGrid[T]
}

type GridMover[T numbers.NumberType] struct {
	model.RecorderClass[GridMoverItem[T]]
}

func NewGridMover[T numbers.NumberType](bounds rl.Vector4, pixelRateX float64) *GridMover[T] {
	mv := &GridMover[T]{}
	var _ model.Parent = mv
	var _ Mover = mv
	var _ Inputer = mv

	mv.Content.Bounds = bounds
	mv.Content.PixelRateX = pixelRateX
	model.SetupRecorder[GridMover[T]](mv, class.LifeMover.String(),
		int32(class.LifeMover))
	// mv.init(true)
	return mv
}
func (cm *GridMover[T]) GetDepth() float32 {
	return 1
}

func (cm *GridMover[T]) GetDrawer() Drawer  { return cm.Content.drawer }
func (cm *GridMover[T]) Bounds() rl.Vector4 { return cm.Content.Bounds }
func (cm *GridMover[T]) Draw(v rl.Vector4)  { cm.Content.drawer.Draw(v) }

func (cm *GridMover[T]) LinkChild(recorder model.Recorder) {
	dr, ok := recorder.(*NumberGrid[T])
	if ok {
		cm.AddDrawer(dr)
	} else {
		log.Fatal(fmt.Errorf("GridMoverLinkChildren: not a NumberGrid"))
	}
}

func (cm *GridMover[T]) Children() []model.Recorder {
	return []model.Recorder{cm.Content.drawer}
}

func (cm *GridMover[T]) AddDrawer(dr *NumberGrid[T]) {
	fmt.Println("GRIDMOVER ADDED DRAWER")
	cm.Content.drawer = dr
}

func doMod(a, b int32) int32 {
	return (b + a) % b
}

func (cs *GridMover[T]) CountNeighbors(cellX, cellY int32) int {
	count := 0
	dr := cs.Content.drawer
	cells := dr.GetCells()

	//-1..1
	for y := int32(-1); y < 2; y++ {
		//-1..1
		for x := int32(-1); x < 2; x++ {
			nx, ny := doMod(cellX+x, dr.Content.Cols), doMod(cellY+y, dr.Content.Rows)
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

func (cm *GridMover[T]) Refresh(now float64, v rl.Vector4, f ...func(any)) {
	cm.Content.Bounds = v
	dr := cm.Content.drawer
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

	cm.Content.drawer.Refresh(0, rl.Vector4{
		X: cm.Content.Bounds.X,
		Y: cm.Content.Bounds.Y}, f)
}

func (cm *GridMover[T]) Move(can_move bool, now float64) {
	if can_move {
		cm.Update()
	}
	cm.Content.drawer.Draw(rl.Vector4{X: float32(cm.Content.Bounds.X),
		Y: float32(cm.Content.Bounds.Y),
		Z: 0})
}

func (cm *GridMover[T]) Update() {
	dr := cm.Content.drawer
	cells := dr.GetCells()

	for i := int32(0); i < dr.Content.Cols; i++ {
		for j := int32(0); j < dr.Content.Rows; j++ {
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
	for i := int32(0); i <= dr.Content.Cols; i++ {
		for j := int32(0); j < dr.Content.Rows; j++ {
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
	if rl.IsKeyDown(rl.KeyRight) && !cm.Content.Playing {
		fmt.Println("KeyRight pressed", cm.Content.Playing)
		cm.Update()
	}
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		cm.Click(rl.GetMouseX(), rl.GetMouseY())
	}
	if rl.IsKeyPressed(rl.KeySpace) {
		cm.Content.Playing = !cm.Content.Playing
		fmt.Println("cm.Playing", cm.Content.Playing)
	}

}

func (cm *GridMover[T]) Click(clickX, clickY int32) {
	var (
		value, x, y, cx, cy int32
		dr                  = cm.Content.drawer
		cells               = dr.GetCells()
		ix, iy              = dr.PositionToCell(clickX, clickY)
	)

	for cy = iy - 1; cy < iy+2; cy++ {

		for cx = ix - 1; cx < ix+2; cx++ {

			x = (dr.Content.Cols + cx) % dr.Content.Cols
			y = (dr.Content.Rows + cy) % dr.Content.Rows

			cell := cells[x][y]
			value = int32(cell.Get(Alive))
			value = numbers.As[int32](value == 0)
			cell.Set(Alive, T(value))
			cell.Set(Next, T(value))
		}
	}
}
