package gizmo

import (
	"image/color"

	"github.com/centretown/xray/capture"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func CreatePaletteFromTextures(BG rl.Color, fixedPalette color.Palette, game *Game) (color.Palette, map[color.Color]uint8) {
	x := int32(0)

	rl.BeginDrawing()

	rl.ClearBackground(BG)

	for _, obj := range game.actors {
		if t, ok := obj.GetDrawer().(*Texture); ok {
			t.Load()
			t.DrawSimple(x, 0)
			x += t.Bounds().Width + 120
		}
	}

	for _, obj := range game.drawers {
		if t, ok := obj.(*Texture); ok {
			t.Load()
			t.DrawSimple(x, 0)
			x += t.Bounds().Width + 120
		}
	}

	rl.EndDrawing()

	pic := rl.LoadImageFromScreen().ToImage()
	return capture.ExtendPalette(fixedPalette, pic, 256)
}
