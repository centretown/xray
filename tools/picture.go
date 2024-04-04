package tools

import (
	"image/color"

	"github.com/centretown/xray/model"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var _ Drawable = (*Picture)(nil)

type PictureItem struct {
	Comment   string
	texture2D rl.Texture2D
}

type Picture struct {
	PictureItem
	Resource *model.Resource
}

func NewPicture(path string) (pic *Picture) {
	pic = &Picture{}
	pic.Comment = "what the hell is"
	pic.Resource = model.NewFileResource(path, model.Picture, &pic.PictureItem)
	return pic
}

func (pic *Picture) GetRecord() *model.Record {
	return pic.Resource.Record
}

func (pic *Picture) Load() *Picture {
	pic.texture2D = rl.LoadTexture(pic.Resource.Path)
	return pic
}

func (pic *Picture) Unload() {
	rl.UnloadTexture(pic.texture2D)
}

func (pic *Picture) DrawSimple(x, y int32) {
	rl.DrawTexture(pic.texture2D, x, y,
		color.RGBA{255, 255, 255, 255})
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
