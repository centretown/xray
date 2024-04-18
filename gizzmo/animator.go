package gizzmo

import (
	"github.com/centretown/xray/gizzmodb/model"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type HasDepth interface {
	GetDepth() int32
}

type Drawer interface {
	model.Recorder
	Bounds() rl.Rectangle
	Refresh(now float64, rect rl.Vector4, funcs ...func(any))
	Draw(rl.Vector4)
}

type Inputer interface {
	Input()
}

type Mover interface {
	model.Recorder
	Drawer
	Move(can_move bool, current float64)
	GetDrawer() Drawer
}

type Motor interface {
	Position() int32
	Direction() int32
	Move(current, rate float64) (position int32)
	Refresh(now float64, max int32)
}
