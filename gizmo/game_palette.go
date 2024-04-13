package gizmo

import (
	"image"
	"image/color"
	"io"
	"log"
	"os"

	"github.com/centretown/xray/capture"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (game *Game) listTextures() (txts []*Texture) {
	txts = make([]*Texture, 0)
	for _, obj := range game.actors {
		if t, ok := obj.GetDrawer().(*Texture); ok {
			txts = append(txts, t)
		}
	}

	for _, obj := range game.drawers {
		if t, ok := obj.(*Texture); ok {
			txts = append(txts, t)
		}
	}
	return
}

func (game *Game) CreatePaletteFromTextures(BG rl.Color, fixedPalette color.Palette) (color.Palette, map[color.Color]uint8) {
	var (
		err  error
		img  image.Image
		imgs []image.Image = make([]image.Image, 0)
		txts []*Texture
		txt  *Texture
	)

	txts = game.listTextures()

	for _, txt = range txts {
		img, err = loadImage(txt.Resource.Path)
		if err == nil {
			imgs = append(imgs, img)
		} else {
			log.Fatal(err)
		}
	}

	pal, colorMap := capture.ExtendPalette(fixedPalette, imgs, 256)
	return pal, colorMap
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
