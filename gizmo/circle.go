package gizmo

import (
	"image/color"

	"github.com/centretown/xray/gizmo/categories"
	"github.com/centretown/xray/model"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var _ Drawer = (*Circle)(nil)
var _ model.Recorder = (*Circle)(nil)

type CircleItem struct {
	Radius int32
	Color  color.RGBA
}

type Circle struct {
	CircleItem
	Record *model.Record
}

func NewCircle(radius int32, colr color.Color) *Circle {
	r, g, b, a := colr.RGBA()
	var rgba = color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)}
	circle := &Circle{}
	circle.Radius = radius
	circle.Color = rgba
	circle.Record = model.NewRecord("circle",
		int32(categories.Circle), &circle.CircleItem, model.JSON)
	return circle
}

func (b *Circle) GetRecord() *model.Record {
	return b.Record
}

func (b *Circle) GetItem() any {
	return &b.CircleItem
}

func (b *Circle) Refresh(float64, rl.RectangleInt32, ...func(any)) {
}

func (b *Circle) Bounds() rl.RectangleInt32 {
	width := b.Radius << 1
	return rl.RectangleInt32{X: 0, Y: 0, Width: width, Height: width}
}

func (b *Circle) Draw(v rl.Vector3) {
	rl.DrawCircle(int32(v.X), int32(v.Y), float32(b.Radius), b.Color)
}
