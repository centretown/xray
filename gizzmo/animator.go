package gizzmo

import (
	rl "github.com/centretown/raylib-go/raylib"
	"github.com/centretown/xray/gizzmodb/model"
)

type HasDepth interface {
	GetDepth() float32
}

type Drawer interface {
	model.Recorder
	Bounds() rl.Vector4
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
	GetDepth() float32
}

type Motor interface {
	Position() int32
	Direction() int32
	Move(current, rate float64) (position int32)
	Refresh(now float64, max int32)
}
