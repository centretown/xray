package main

import (
	"image/color"

	"github.com/centretown/xray/builder"
	"github.com/centretown/xray/gizzmo"
)

func main() {
	game, install, err := builder.Build(build_life01)
	if err == nil && !install {
		game.Run()
	}
}

func build_life01(game *gizzmo.Game) {
	viewPort := game.GetViewPort()
	cells := gizzmo.NewCellsOrg(int32(viewPort.Width), int32(viewPort.Height), 12)
	game.AddDrawer(cells)
	game.Content.FrameRate = 20
	game.Content.FixedSize = true
	cells.Content.Colors = []color.RGBA{gizzmo.White, gizzmo.Blue, gizzmo.Green}
	game.AddColors(cells.Content.Colors...)
}
