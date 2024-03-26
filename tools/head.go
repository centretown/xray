package tools

import rl "github.com/gen2brain/raylib-go/raylib"

var _ CanDraw = (*HeadText)(nil)

type HeadText struct {
	rl.Texture2D
}

func NewHeadText(txt rl.Texture2D) *HeadText {
	text := &HeadText{
		Texture2D: txt,
	}
	return text
}

func (hd *HeadText) Draw(x, y int32) {
	rl.DrawTexture(hd.Texture2D, x, y, rl.White)
}

func (hd *HeadText) MinSize() (w, h int32) {
	return hd.Texture2D.Width, hd.Texture2D.Width
}
func (hd *HeadText) Width() int32 {
	return hd.Texture2D.Width
}
func (hd *HeadText) Height() int32 {
	return hd.Texture2D.Width
}

// func (hd *HeadText) Resize(boundsX, boundsY int32) {
// 	rl.UpdateTextureRec(hd.texture, Rectangle{0,0,boundsX,boundsY}, const void *pixels);
// }
