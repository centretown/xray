package tools

import rl "github.com/gen2brain/raylib-go/raylib"

type Drawable interface {
	Draw(x, y int32)
}

type Animator interface {
	Animate(can_move int32, dr Drawable)
	Resize(rect rl.RectangleInt32, boundsX, boundsY int32)
}
