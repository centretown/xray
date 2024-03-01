package tools

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var colors = []rl.Color{
	rl.White,
	rl.Blue,
	rl.Yellow,
	rl.Red,
	rl.White,
	rl.Red,
	rl.White,
	rl.Yellow,
	rl.Lime,
	rl.DarkGreen,
}

var _ Drawable = (*Ball)(nil)

type Ball struct {
	Radius int32
	Colors []rl.Color
	draw   func(x, y int32)
}

func NewBall(radius int32, colors []rl.Color) *Ball {
	b := &Ball{
		Radius: radius,
		Colors: colors,
	}

	if len(b.Colors) == 0 {
		b.Colors = append(b.Colors, rl.Lime)
	}

	if len(b.Colors) < 2 {
		b.draw = b.drawSolid
	} else {
		b.draw = b.drawGradient
	}
	return b
}

func (b *Ball) Resize(radius int32) {
	if radius <= 0 {
		panic("ball zero radius")
	}
	b.Radius = radius
}

func (b *Ball) Draw(x, y int32) {
	b.draw(x, y)
}

func (b *Ball) drawGradient(x, y int32) {
	rl.DrawCircleGradient(x, y, float32(b.Radius),
		b.Colors[0], b.Colors[1])
}

func (b *Ball) drawSolid(x, y int32) {
	rl.DrawCircle(x, y, float32(b.Radius), b.Colors[0])
}
