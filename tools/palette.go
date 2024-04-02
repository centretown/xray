package tools

import (
	"image/color"

	"github.com/centretown/xray/capture"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func CreatePaletteFromTextures(BG rl.Color, fixedPalette color.Palette, actors ...Moveable) (color.Palette, map[color.Color]uint8) {

	rl.BeginDrawing()

	rl.ClearBackground(BG)
	x := int32(0)

	for _, actor := range actors {
		t, ok := actor.Drawer().(*Picture)
		if ok {
			t.DrawSimple(x, 0)
			x += t.Rect().Width + 120
		}
	}
	rl.EndDrawing()

	pic := rl.LoadImageFromScreen().ToImage()
	return capture.ExtendPalette(fixedPalette, pic, 256)
}
