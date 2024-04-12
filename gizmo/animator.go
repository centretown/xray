package gizmo

import (
	"github.com/centretown/xray/model"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Drawer interface {
	model.Recorder
	Bounds() rl.RectangleInt32
	Draw(rl.Vector3)
	Refresh(rect rl.RectangleInt32, options ...bool)
}

type Inputer interface {
	Input()
}

type Actor interface {
	model.Recorder
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
