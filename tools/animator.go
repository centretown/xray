package tools

import rl "github.com/gen2brain/raylib-go/raylib"

type Drawable interface {
	Rect() rl.RectangleInt32
	Draw(rl.Vector3)
}

type Moveable interface {
	Draw(can_move bool, current float64, dr Drawable)
	Refresh(now float64, rect rl.RectangleInt32)
}

type Action interface {
	Position() int32
	Direction() int32
	Next(current, rate float64) (position int32)
	Refresh(now float64, position, max int32)
}
