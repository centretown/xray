package tools

import rl "github.com/gen2brain/raylib-go/raylib"

type Drawable interface {
	Draw(x, y int32)
	Rect() rl.RectangleInt32
}

type Moveable interface {
	Draw(can_move bool, current float64, dr Drawable)
	Refresh(current float64, rect rl.RectangleInt32)
	// SetPixelRate(float64, float64)
	// GetPixelRate() (float64, float64)
}
