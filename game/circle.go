package game

import (
	"encoding/json"
	"fmt"
	"image/color"

	"github.com/centretown/xray/game/categories"
	"github.com/centretown/xray/model"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// const min_ball_radius = 5?

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

func NewCircle(radius int32, col color.Color) *Circle {
	r, g, b, a := col.RGBA()
	var c = color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(a),
	}

	circle := &Circle{}
	circle.Radius = radius
	circle.Color = c
	circle.Record = model.NewRecord("circle",
		int32(categories.Circle), &circle.CircleItem, model.JSON)

	return circle
}

func (b *Circle) GetRecord() *model.Record { return b.Record }
func (b *Circle) GetItem() any             { return &b.CircleItem }

func (b *Circle) Decode(rec *model.Record) (err error) {
	b.Record = rec
	cat := categories.Category(rec.Category)
	if cat == categories.Circle {
		err = json.Unmarshal([]byte(rec.Content), &b.CircleItem)
		if err != nil {
			panic(err)
		}
	} else {
		err = fmt.Errorf("wrong category want %s have %s",
			categories.Circle, cat)
	}
	return
}

func (b *Circle) Bounds() rl.RectangleInt32 {
	width := b.Radius << 1
	return rl.RectangleInt32{
		X:      0,
		Y:      0,
		Width:  width,
		Height: width,
	}
}

func (b *Circle) Draw(v rl.Vector3) {
	x, y := int32(v.X), int32(v.Y)
	rl.DrawCircle(x, y, float32(b.Radius), b.Color)
}
