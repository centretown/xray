package tools

import (
	"image/color"

	"github.com/centretown/xray/b2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var _ Drawable = (*Picture)(nil)

type Picture struct {
	rl.Texture2D
	Rotation int
	lastX    int32
}

func NewPicture(txt rl.Texture2D, rotation int) *Picture {
	text := &Picture{
		Texture2D: txt,
		Rotation:  rotation,
	}
	return text
}

func (td *Picture) Draw(x, y int32) {
	width, height := float32(td.Texture2D.Width), float32(td.Texture2D.Height)
	srcRec := rl.Rectangle{X: 0, Y: 0, Width: width, Height: height}
	destRec := rl.Rectangle{X: float32(x), Y: float32(y), Width: width, Height: height}
	origin := rl.Vector2{X: width / 2, Y: height / 2}
	rl.DrawTexturePro(td.Texture2D, srcRec, destRec, origin,
		float32(td.Rotation), color.RGBA{255, 255, 255, 255})
	leftToRight := x > td.lastX
	td.lastX = x
	td.Rotation += b2.To[int](leftToRight)<<2 - b2.To[int](!leftToRight)<<2
}

func (td *Picture) Rect() rl.RectangleInt32 {
	return rl.RectangleInt32{X: 0, Y: 0,
		Width: td.Texture2D.Width, Height: td.Texture2D.Width}
}
