package game

import rl "github.com/gen2brain/raylib-go/raylib"

type Drawer interface {
	Bounds() rl.RectangleInt32
	Draw(rl.Vector3)
}

type Actor interface {
	Act(can_move bool, current float64)
	Refresh(now float64, rect rl.RectangleInt32)
	GetDrawer() Drawer
}

type Motor interface {
	Position() int32
	Direction() int32
	Move(current, rate float64) (position int32)
	Refresh(now float64, max int32)
}
