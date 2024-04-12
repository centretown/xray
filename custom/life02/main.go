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
	viewPort := game.SetViewPort(600, 480)
	grid := gizmo.NewGrid[bool](viewPort, 40, 40,
		color.RGBA{R: 12, G: 195, B: 56, A: 255},
		color.RGBA{R: 12, G: 56, B: 195, A: 255},
		color.RGBA{R: 128, G: 128, B: 0, A: 255},
		color.RGBA{R: 255, G: 255, B: 0, A: 255},
	)
	grid_mv := gizmo.NewCellsMover(viewPort, 5)
	grid_mv.AddDrawer(grid)

	game.AddActor(grid_mv, 1)
	game.FrameRate = 20
	// game.FixedSize = true
	game.AddColors(grid.Colors)
}
