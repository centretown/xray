package tools

import rl "github.com/gen2brain/raylib-go/raylib"

type Drawable interface {
	Draw(rl.Vector3)
	Rect() rl.RectangleInt32
}

type Moveable interface {
	Draw(can_move bool, current float64, dr Drawable)
	Refresh(now float64, rect rl.RectangleInt32)
}

type Action interface {
	// SetPixelRate(float64, float64)
	// GetPixelRate() (float64, float64)
	Refresh(now float64, position, max int32)
	Next(current, rate float64)
	Position() int32
	Direction() int32
}
