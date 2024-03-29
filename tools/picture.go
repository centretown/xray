package tools

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var _ Drawable = (*Picture)(nil)

type Picture struct {
	rl.Texture2D
	// Rotation int
	// lastX    int32
}

func NewPicture(texture rl.Texture2D, rotation int) *Picture {
	text := &Picture{
		Texture2D: texture,
		// Rotation:  rotation,
	}
	return text
}

func (td *Picture) Draw(v rl.Vector3) {
	x, y, rotation := v.X, v.Y, v.Z
	width, height := float32(td.Texture2D.Width), float32(td.Texture2D.Height)
	srcRec := rl.Rectangle{X: 0, Y: 0, Width: width, Height: height}
	destRec := rl.Rectangle{X: x, Y: y, Width: width, Height: height}
	origin := rl.Vector2{X: width / 2, Y: height / 2}

	rl.DrawTexturePro(td.Texture2D, srcRec, destRec, origin,
		rotation, color.RGBA{255, 255, 255, 255})
}

func (td *Picture) Rect() rl.RectangleInt32 {
	return rl.RectangleInt32{X: 0, Y: 0,
		Width: td.Texture2D.Width, Height: td.Texture2D.Width}
}
