package tools

import (
	"image/color"
	"xray/b2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var _ CanDraw = (*TextureDrawer)(nil)

type TextureDrawer struct {
	rl.Texture2D
	rotation int
	lastX    int32
}

func NewTextureDrawer(txt rl.Texture2D) *TextureDrawer {
	text := &TextureDrawer{
		Texture2D: txt,
		rotation:  1,
	}
	return text
}

func (hd *TextureDrawer) Draw(x, y int32) {

	var w, h = float32(hd.Width()), float32(hd.Height())
	srcRec := rl.Rectangle{X: 0, Y: 0, Width: w, Height: h}
	destRec := rl.Rectangle{X: float32(x), Y: float32(y), Width: w, Height: h}
	origin := rl.Vector2{X: w / 2, Y: h / 2}
	rl.DrawTexturePro(hd.Texture2D, srcRec, destRec, origin, float32(hd.rotation), color.RGBA{255, 255, 255, 255})

	ltr := x > hd.lastX
	hd.lastX = x
	hd.rotation += b2.To[int](ltr)<<2 - b2.To[int](!ltr)<<2
}

func (hd *TextureDrawer) MinSize() (w, h int32) {
	return hd.Texture2D.Width, hd.Texture2D.Width
}
func (hd *TextureDrawer) Width() int32 {
	return hd.Texture2D.Width
}
func (hd *TextureDrawer) Height() int32 {
	return hd.Texture2D.Width
}

func (hd *TextureDrawer) Resize(boundsX, boundsY int32) {
	// rl.UpdateTextureRec(hd.texture, Rectangle{0,0,boundsX,boundsY}, const void *pixels);
}
