package tools

import (
	"fmt"
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const min_ball_radius = 5

var _ CanDraw = (*Ball)(nil)

type Ball struct {
	Radius int32
	Color  color.RGBA
	draw   func(x, y int32)
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

	ball.draw = ball.drawSolid
	return ball
}

func (b *Ball) Width() int32 {
	return b.Radius << 1
}

func (b *Ball) Height() int32 {
	return b.Radius << 1
}

func (b *Ball) MinSize() (int32, int32) {
	return min_ball_radius, min_ball_radius
}

func (b *Ball) Resize(width, height int32) {
	radius := min(width, height)
	if radius <= min_ball_radius {
		panic(fmt.Sprintf("%d less than minimum radius", radius))
	}
	b.Radius = radius
}

func (b *Ball) Draw(x, y int32) {
	b.draw(x, y)
}

// func (b *Ball) drawGradient(x, y int32) {
// 	rl.DrawCircleGradient(x, y, float32(b.Radius),
// 		b.Colors[0], b.Colors[1])
// }

func (b *Ball) drawSolid(x, y int32) {
	rl.DrawCircle(x, y, float32(b.Radius), b.Color)
}
