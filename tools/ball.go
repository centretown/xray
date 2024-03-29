package tools

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const min_ball_radius = 5

var _ Drawable = (*Ball)(nil)

type Ball struct {
	Radius int32
	Color  color.RGBA
}

func NewBall(radius int32, col color.Color) *Ball {
	r, g, b, a := col.RGBA()
	var c = color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(a),
	}

	ball := &Ball{
		Radius: radius,
		Color:  c,
	}
	return ball
}

func (b *Ball) Rect() rl.RectangleInt32 {
	width := b.Radius << 1
	return rl.RectangleInt32{
		X:      0,
		Y:      0,
		Width:  width,
		Height: width,
	}
}

func (b *Ball) Draw(x, y int32) {
	rl.DrawCircle(x, y, float32(b.Radius), b.Color)
}
