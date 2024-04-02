package main

import (
	"image/color"

	"github.com/centretown/xray/capture"
	"github.com/centretown/xray/tools"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var BG = rl.Black

const (
	TRANSPARENT = iota
	WHITE
	BLACK
	RED
	YELLOW
	GREEN
	CYAN
	BLUE
	MAGENTA
)

var fixedPal = color.Palette{
	color.Transparent,
	rl.White,
	rl.Black,
	rl.Red,
	rl.Yellow,
	rl.Green,
	color.RGBA{R: 0, G: 255, B: 255, A: 255},
	rl.Blue,
	rl.Magenta,
}

func createPaletteFromTextures(pal color.Palette, actors ...*tools.Actor) (color.Palette, map[color.Color]uint8) {

	rl.BeginDrawing()

	rl.ClearBackground(BG)
	x := int32(0)

	for _, actor := range actors {
		t, ok := actor.Character.(*tools.Picture)
		if ok {
			t.DrawSimple(x, 0)
			x += t.Rect().Width + 120
		}
	}
	rl.EndDrawing()

	pic := rl.LoadImageFromScreen().ToImage()
	return capture.ExtendPalette(pal, pic, 256)
}
