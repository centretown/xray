package gizzmo

import (
	"fmt"
	"image/color"

	"github.com/centretown/xray/gizzmo/class"
	"github.com/centretown/xray/gizzmodb/model"
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

func NewTexture(path string, depth float32) *Texture {
	tex := &Texture{}
	InitShape[TextureItem](&tex.Shape, class.Texture.String(), int32(class.Texture),
		color.RGBA{}, 0, 0, depth)
	model.InitResource(&tex.Content.Custom.Resource, path, int32(class.Texture))
	var _ Drawer = tex

	return tex
}

func (tex *Texture) Load() *Texture {
	fmt.Println("LOAD TEXTURE")
	res := &tex.Content.Custom.Resource
	if res.Err == nil {
		tex.Content.Custom.texture2D = rl.LoadTexture(tex.Content.Custom.Resource.Path)
		fmt.Println("TEXTURE LOADed")
		tex.Content.Dimensions.X = float32(tex.Content.Custom.texture2D.Width)
		tex.Content.Dimensions.Y = float32(tex.Content.Custom.texture2D.Height)
	}
	return tex
}

func (tex *Texture) Unload() { rl.UnloadTexture(tex.Content.Custom.texture2D) }

//   Generate GPU mipmaps for a texture
// void GenTextureMipmaps(Texture2D *texture);

// Set texture scaling filter mode
// void SetTextureFilter(Texture2D texture, int filter);

// void SetTextureWrap(Texture2D texture, int wrap);

func (tex *Texture) Draw(v rl.Vector4) {
	x, y, z, rotation := v.X, v.Y, v.Z, v.W
	// x, y, rotation := v.X, v.Y, v.W
	// fmt.Println("z", z, tex.Content.Dimensions.Z)

	scale := z / tex.Content.Dimensions.Z

	width, height := float32(tex.Content.Custom.texture2D.Width),
		float32(tex.Content.Custom.texture2D.Height)
	source := rl.Rectangle{X: 0, Y: 0, Width: width, Height: height}
	destination := rl.Rectangle{X: x, Y: y, Width: scale * width, Height: scale * height}
	origin := rl.Vector2{X: scale * width / 2, Y: scale * height / 2}
	rl.DrawTexturePro(tex.Content.Custom.texture2D, source, destination, origin,
		rotation, White)

}
