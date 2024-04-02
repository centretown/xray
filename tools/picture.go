package tools

import (
	"image/color"

	"github.com/centretown/xray/model"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var _ Drawable = (*Picture)(nil)

type Picture struct {
	*model.Resource
	texture2D rl.Texture2D
}

func NewPicture(path string) (pic *Picture) {
	res := model.NewFileResource(path, model.Picture)
	pic = &Picture{
		Resource: res,
	}

	if res.Err == nil && pic.Item.Scheme == model.File {
		pic.texture2D = rl.LoadTexture(pic.Item.Path)
	}

	return pic
}

func (pic *Picture) Unload() {
	rl.UnloadTexture(pic.texture2D)
}

func (pic *Picture) DrawSimple(x, l int32) {
	rl.DrawTexture(pic.texture2D, x, 0, rl.White)
}

func (pic *Picture) Draw(v rl.Vector3) {
	x, y, rotation := v.X, v.Y, v.Z
	width, height := float32(pic.texture2D.Width), float32(pic.texture2D.Height)
	srcRec := rl.Rectangle{X: 0, Y: 0, Width: width, Height: height}
	destRec := rl.Rectangle{X: x, Y: y, Width: width, Height: height}
	origin := rl.Vector2{X: width / 2, Y: height / 2}

	rl.DrawTexturePro(pic.texture2D, srcRec, destRec, origin,
		rotation, color.RGBA{255, 255, 255, 255})
}

func (pic *Picture) Rect() rl.RectangleInt32 {
	return rl.RectangleInt32{X: 0, Y: 0,
		Width: pic.texture2D.Width, Height: pic.texture2D.Width}
}
