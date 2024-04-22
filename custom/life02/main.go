package main

import (
	"image/color"

	rl "github.com/centretown/raylib-go/raylib"
	"github.com/centretown/xray/builder"
	"github.com/centretown/xray/gizzmo"
)

func main() {
	game, install, err := builder.Build(build_life02)
	if err == nil && !install {
		game.Run()
	}
}

func build_life02(game *gizzmo.Game) {
	vf := game.SetViewPort(900, 800)
	vp := rl.Vector4{X: vf.Width, Y: vf.Height}

	game.Content.Title = "Life"
	game.Content.Author = "Dave"
	game.Content.Description = "Life is a testing game. It implements NumberGrid, NumberGridMover"
	game.Content.Instructions = `Staying Alive
	- Press R to reset.
	- Press C to clear.`

	grid := gizzmo.NewGrid[int8](vp, 40, 40,
		color.RGBA{R: 255, G: 56, B: 12, A: 255},  //horizontal
		color.RGBA{R: 255, G: 255, B: 12, A: 255}, //vertical

		color.RGBA{R: 0, G: 240, B: 0, A: 255},   //Alive
		color.RGBA{R: 255, G: 255, B: 0, A: 255}, //Visited
		color.RGBA{R: 0, G: 0, B: 255, A: 255},   //Next
	)
	game.AddColors(grid.Content.HorizontalColor, grid.Content.VerticalColor)
	game.AddColors(grid.Content.StateColors...)

	grid_mv := gizzmo.NewGridMover[int8](vp, 5)
	grid_mv.AddDrawer(grid)

	game.AddActor(grid_mv, 1)
	game.Content.FrameRate = 20
	game.Content.BackGround = color.RGBA{R: 31, G: 31, B: 31, A: 255}
}
