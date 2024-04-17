package gizmo

import (
	"fmt"
	"image/color"

	"github.com/centretown/xray/gizmo/class"
	"github.com/centretown/xray/gizmodb/model"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var _ Drawer = (*Texture)(nil)
var _ model.Recorder = (*Texture)(nil)

type TextureItem struct {
	Resource  model.Resource
	texture2D rl.Texture2D
}

type Texture struct {
	Shape[TextureItem]
}

func NewTextureFromRecord(record *model.Record) (tex *Texture) {
	tex = &Texture{}
	ShapeFromRecord(&tex.Shape, record)
	model.InitResource(&tex.Content.Custom.Resource,
		tex.Content.Custom.Resource.Path, int32(class.Texture))
	return tex
}

func NewTexture(path string) *Texture {
	tex := &Texture{}
	InitShape[TextureItem](&tex.Shape, class.Texture.String(), int32(class.Texture),
		color.RGBA{}, 0, 0)
	model.InitResource(&tex.Content.Custom.Resource, path, int32(class.Texture))
	var _ Drawer = tex

	return tex
}

func (tex *Texture) Load() *Texture {
	fmt.Println("LOAD TEXTURE")
	res := &tex.Content.Custom.Resource

	// if !tex.initialized {
	// 	model.InitResource(res,
	// 		res.Path, int32(categories.Texture))
	// 	tex.initialized = true
	// }

	if res.Err == nil {
		tex.Content.Custom.texture2D = rl.LoadTexture(tex.Content.Custom.Resource.Path)
		fmt.Println("TEXTURE LOADed")
	}
	return tex
}

func (tex *Texture) Unload() { rl.UnloadTexture(tex.Content.Custom.texture2D) }

func (tex *Texture) Draw(v rl.Vector4) {
	x, y, rotation := v.X, v.Y, v.W
	width, height := float32(tex.Content.Custom.texture2D.Width),
		float32(tex.Content.Custom.texture2D.Height)
	srcRec := rl.Rectangle{X: 0, Y: 0, Width: width, Height: height}
	destRec := rl.Rectangle{X: x, Y: y, Width: width, Height: height}
	origin := rl.Vector2{X: width / 2, Y: height / 2}

	rl.DrawTexturePro(tex.Content.Custom.texture2D, srcRec, destRec, origin,
		rotation, White)
}
