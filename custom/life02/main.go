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
	grid := gizmo.NewGrid[bool](viewPort, 10, 10)
	grid_mv := gizmo.NewCellsMover(viewPort, 5)
	grid_mv.AddDrawer(grid)

	game.AddActor(grid_mv, 1)
	game.FrameRate = 20
	grid.Colors = []color.RGBA{gizmo.Green, gizmo.Blue, gizmo.Yellow}
	game.AddColors(grid.Colors)
}
