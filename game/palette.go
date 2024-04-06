package game

import (
	"image/color"

	"github.com/centretown/xray/capture"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func CreatePaletteFromTextures(BG rl.Color, fixedPalette color.Palette, actors ...*Mover) (color.Palette, map[color.Color]uint8) {

	rl.BeginDrawing()

	rl.ClearBackground(BG)
	x := int32(0)

	for _, actor := range actors {
		t, ok := actor.GetDrawer().(*Texture)
		if ok {
			t.DrawSimple(x, 0)
			x += t.Bounds().Width + 120
		}
	}
	rl.EndDrawing()

	pic := rl.LoadImageFromScreen().ToImage()
	return capture.ExtendPalette(fixedPalette, pic, 256)
}