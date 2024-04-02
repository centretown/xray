package tools

import (
	"image/color"

	"github.com/centretown/xray/rayl"
)

// const min_ball_radius = 5?

var _ Drawable = (*Circle)(nil)

type Circle struct {
	Radius int32
	Color  color.RGBA
}

func NewBall(radius int32, col color.Color) *Circle {
	r, g, b, a := col.RGBA()
	var c = color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(a),
	}

	ball := &Circle{
		Radius: radius,
		Color:  c,
	}
	return ball
}

func (b *Circle) Rect() rayl.RectangleInt32 {
	width := b.Radius << 1
	return rayl.RectangleInt32{
		X:      0,
		Y:      0,
		Width:  width,
		Height: width,
	}
}

func (b *Circle) Draw(v rayl.Vector3) {
	// x, y := int32(v.X), int32(v.Y)
	// rl.DrawCircle(x, y, float32(b.Radius), b.Color)
}
