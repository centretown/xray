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

func NewBall(radius int32, col color.RGBA) *Ball {
	b := &Ball{
		Radius: radius,
		Color:  col,
	}

	b.draw = b.drawSolid
	return b
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
