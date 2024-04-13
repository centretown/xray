package gizmo

import (
	"github.com/centretown/xray/gizmo/categories"
	"github.com/centretown/xray/model"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var _ Drawer = (*Texture)(nil)
var _ model.Recorder = (*Texture)(nil)

type TextureItem struct {
	Resource  *model.Resource
	texture2D rl.Texture2D
}

type Texture struct {
	TextureItem
	Record *model.Record
}

func NewTexture(path string) (tex *Texture) {
	tex = &Texture{}
	tex.Resource = model.NewFileResource(path, int32(categories.Texture), &tex.TextureItem)
	tex.Record = model.NewRecord("texture",
		int32(categories.Texture), &tex.TextureItem, model.JSON)
	return tex
}

func (tex *Texture) GetRecord() *model.Record { return tex.Record }
func (tex *Texture) GetItem() any             { return &tex.TextureItem }

func (tex *Texture) Load() *Texture {
	if tex.Resource != nil && tex.Resource.Err == nil {
		tex.texture2D = rl.LoadTexture(tex.Resource.Path)
	}
	return tex
}

// func (tex *Texture) LoadImage() (img image.Image, err error) {
// 	var (
// 		title = "LoadImage"
// 		rdr   io.ReadCloser
// 		enc   string
// 	)

// 	rdr, err = os.Open(tex.Resource.Path)
// 	if err != nil {
// 		log.Println()
// 	}
// 	defer rdr.Close()

// 	img, enc, err = image.Decode(rdr)
// 	if err != nil {
// 		log.Printf("%s %s file %s\n", title, tex.Resource.Path, err)
// 	} else {
// 		log.Printf("%s %s file %s success\n", title, enc, tex.Resource.Path)
// 	}
// 	return
// }

func (tex *Texture) Unload() { rl.UnloadTexture(tex.texture2D) }

func (tex *Texture) DrawSimple(x, y int32) {
	rl.DrawTexture(tex.texture2D, x, y, White)
}

func (tex *Texture) Draw(v rl.Vector3) {
	x, y, rotation := v.X, v.Y, v.Z
	width, height := float32(tex.texture2D.Width), float32(tex.texture2D.Height)
	srcRec := rl.Rectangle{X: 0, Y: 0, Width: width, Height: height}
	destRec := rl.Rectangle{X: x, Y: y, Width: width, Height: height}
	origin := rl.Vector2{X: width / 2, Y: height / 2}

	rl.DrawTexturePro(tex.texture2D, srcRec, destRec, origin,
		rotation, White)
}
func (tex *Texture) Refresh(rect rl.RectangleInt32, options ...bool) {}

func (tex *Texture) Bounds() rl.RectangleInt32 {
	return rl.RectangleInt32{X: 0, Y: 0,
		Width: tex.Resource.Width, Height: tex.Resource.Height}
}
