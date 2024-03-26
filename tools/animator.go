package tools

import rl "github.com/gen2brain/raylib-go/raylib"

type CanResize interface {
	Resize(w, h int32)
	MinSize() (w, h int32)
	Width() int32
	Height() int32
}

type CanDraw interface {
	Draw(x, y int32)
}

type CanMove interface {
	Animate(can_move int32, dr CanDraw)
	Position() (int32, int32)
	Resize(rect rl.RectangleInt32, boundsX, boundsY int32)
}
