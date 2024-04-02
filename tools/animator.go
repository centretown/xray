package tools

import "github.com/centretown/xray/rayl"

type Drawable interface {
	Rect() rayl.RectangleInt32
	Draw(rayl.Vector3)
}

type Moveable interface {
	Move(can_move bool, current float64)
	Refresh(now float64, rect rayl.RectangleInt32)
	Drawer() Drawable
}

type Action interface {
	Position() int32
	Direction() int32
	Next(current, rate float64) (position int32)
	Refresh(now float64, max int32)
}
