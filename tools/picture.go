package tools

import (
	"image/color"

	"github.com/centretown/xray/model"
	"github.com/centretown/xray/rayl"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var _ Drawable = (*Picture)(nil)

type Picture struct {
	*model.Resource
	Comment   string
	texture2D rl.Texture2D
}

func NewPicture(path string) (pic *Picture) {
	pic = &Picture{Comment: "game element"}
	res := model.NewFileResource(path, model.Picture, &pic.Comment)
	pic.Resource = res
	return pic
}

func (pic *Picture) Load() *Picture {
	pic.texture2D = rl.LoadTexture(pic.Item.Path)
	return pic
}

func (pic *Picture) Unload() {
	rl.UnloadTexture(pic.texture2D)
}

func (pic *Picture) DrawSimple(x, y int32) {
	rl.DrawTexture(pic.texture2D, x, y, rl.White)
}

func (pic *Picture) Draw(v rayl.Vector3) {
	x, y, rotation := v.X, v.Y, v.Z
	width, height := float32(pic.texture2D.Width), float32(pic.texture2D.Height)
	srcRec := rl.Rectangle{X: 0, Y: 0, Width: width, Height: height}
	destRec := rl.Rectangle{X: x, Y: y, Width: width, Height: height}
	origin := rl.Vector2{X: width / 2, Y: height / 2}

	rl.DrawTexturePro(pic.texture2D, srcRec, destRec, origin,
		rotation, color.RGBA{255, 255, 255, 255})
}

func (pic *Picture) Rect() rayl.RectangleInt32 {
	return rayl.RectangleInt32{X: 0, Y: 0,
		Width: pic.texture2D.Width, Height: pic.texture2D.Width}
}
