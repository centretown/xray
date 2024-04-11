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

func build_life01(game *gizmo.Game, resourcePath string) {
	viewPort := game.GetViewPort()
	cells := gizmo.NewCellsOrg(viewPort.Width, viewPort.Height, 12)
	game.AddDrawer(cells)
	game.FrameRate = 20
	game.FixedSize = true
	cells.Colors = []color.RGBA{gizmo.White, gizmo.Blue, gizmo.Green}
	game.AddColors(cells.Colors)
}
