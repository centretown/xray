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
	viewPort := game.SetViewPort(900, 800)
	grid := gizmo.NewGrid[int8](viewPort, 40, 40,
		color.RGBA{R: 255, G: 56, B: 12, A: 255},
		color.RGBA{R: 255, G: 255, B: 12, A: 255},

		color.RGBA{R: 0, G: 240, B: 0, A: 255},   //Alive
		color.RGBA{R: 255, G: 255, B: 0, A: 255}, //Visited
		color.RGBA{R: 0, G: 0, B: 255, A: 255},   //Next
		color.RGBA{R: 0, G: 0, B: 0, A: 255},     //off
	)
	grid_mv := gizmo.NewGridMover[int8](viewPort, 5)
	grid_mv.AddDrawer(grid)

	game.AddActor(grid_mv, 1)
	game.FrameRate = 20
	// game.FixedSize = true
	game.BackGround = color.RGBA{R: 31, G: 31, B: 31, A: 255}
	game.AddColors(grid.Colors)
}
