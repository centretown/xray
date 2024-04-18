package gizzmo

import (
	"image"
	"image/color"
	"io"
	"log"
	"os"

	"github.com/centretown/xray/capture"
)

var (
	Black   = color.RGBA{R: 0, G: 0, B: 0, A: 255}
	White   = color.RGBA{R: 255, G: 255, B: 255, A: 255}
	Red     = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	Yellow  = color.RGBA{R: 255, G: 255, B: 0, A: 255}
	Green   = color.RGBA{R: 0, G: 255, B: 0, A: 255}
	Cyan    = color.RGBA{R: 0, G: 255, B: 255, A: 255}
	Blue    = color.RGBA{R: 0, G: 0, B: 255, A: 255}
	Magenta = color.RGBA{R: 255, G: 0, B: 255, A: 255}
)

// AddColors to the FixedPalette
func (gs *Game) AddColors(clrs ...color.RGBA) {
	gs.Content.FixedPalette = append(gs.Content.FixedPalette, clrs...)
}

// createPalette generates a 256 color palette from RGBA colors
// found in the game textures
func (gs *Game) createPalette() {

	gs.Content.FixedPalette = append(gs.Content.FixedPalette,
		color.RGBA{R: 255, G: 255, B: 255, A: 0},
		Black,
		White,
		Red,
		Yellow,
		Green,
		Cyan,
		Blue,
		Magenta,
	)

	var (
		err  error
		img  image.Image
		imgs []image.Image = make([]image.Image, 0)
		txt  *Texture
	)

	for _, txt = range gs.Content.textureList {
		img, err = loadImage(txt.Content.Custom.Resource.Path)
		if err == nil {
			imgs = append(imgs, img)
		} else {
			log.Fatal(err)
		}
	}

	gs.Content.palette, gs.Content.colorMap =
		capture.ExtendPalette(gs.Content.FixedPalette, imgs, 256)
}

func loadImage(path string) (img image.Image, err error) {
	var (
		rdr io.ReadCloser
	)

	rdr, err = os.Open(path)
	if err != nil {
		return
	}
	defer rdr.Close()

	img, _, err = image.Decode(rdr)
	return
}
