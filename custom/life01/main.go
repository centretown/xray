package main

import (
	"image/color"

	"github.com/centretown/xray/builder"
	"github.com/centretown/xray/gizmo"
)

func main() {
	game, install, err := builder.Build(build_life01)
	if err == nil && !install {
		game.Run()
	}
}

func build_life01(game *gizmo.Game) {
	viewPort := game.GetViewPort()
	cells := gizmo.NewCellsOrg(viewPort.Width, viewPort.Height, 12)
	game.AddDrawer(cells)
	game.Content.FrameRate = 20
	game.Content.FixedSize = true
	cells.Content.Colors = []color.RGBA{gizmo.White, gizmo.Blue, gizmo.Green}
	game.AddColors(cells.Content.Colors...)
}
