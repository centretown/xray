package main

import (
	"image/color"

	"github.com/centretown/xray/builder"
	"github.com/centretown/xray/gizmo"
)

func main() {
	game, install, err := builder.Build(build_life02)
	if err == nil && !install {
		game.Run()
	}
}

func build_life02(game *gizmo.Game, resourcePath string) {
	viewPort := game.GetViewPort()
	cells := gizmo.NewCells(viewPort.Width, viewPort.Height, 12)
	cells_mv := gizmo.NewCellsMover(viewPort, 5)
	cells_mv.AddDrawer(cells)

	game.AddActor(cells_mv, 1)
	game.FrameRate = 20
	game.FixedSize = true
	cells.Colors = []color.RGBA{gizmo.Red, gizmo.Blue, gizmo.Yellow}
	game.AddColors(cells.Colors)
}
