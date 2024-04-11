package main

import (
	"image/color"
	"path/filepath"

	"github.com/centretown/xray/builder"
	"github.com/centretown/xray/gizmo"
)

func main() {
	game, err := builder.Build(build_gander)
	if err == nil {
		game.Run()
	}
}

func build_gander(game *gizmo.Game, resourcePath string) {
	var fixedPalette = []color.RGBA{
		gizmo.White,
		gizmo.Black,
		gizmo.Red,
		gizmo.Yellow,
		gizmo.Green,
		gizmo.Cyan,
		gizmo.Blue,
		gizmo.Magenta,
	}

	viewPort := game.GetViewPort()

	hole := gizmo.NewTexture(filepath.Join(resourcePath, "polar.png"))
	hole_mv := gizmo.NewMover(viewPort, 5, 5, 5)
	hole_mv.AddDrawer(hole)
	game.AddActor(hole_mv, 6)

	moon := gizmo.NewTexture(filepath.Join(resourcePath, "moon-solo-300.png"))
	moon_mv := gizmo.NewMover(viewPort, 10, 10, 5)
	moon_mv.AddDrawer(moon)
	game.AddActor(moon_mv, 6)

	ball := gizmo.NewCircle(20, gizmo.Red)
	ball_mv := gizmo.NewMover(viewPort, 200, 100, 0)
	ball_mv.AddDrawer(ball)
	game.AddActor(ball_mv, 7)

	head := gizmo.NewTexture(filepath.Join(resourcePath, "head_300.png"))
	head_mv := gizmo.NewMover(viewPort, 70, 140, 1.75)
	head_mv.AddDrawer(head)
	game.AddActor(head_mv, 8)

	gander := gizmo.NewTexture(filepath.Join(resourcePath, "gander.png"))
	gander_mv := gizmo.NewMover(viewPort, 300, 300, 0.5)
	gander_mv.AddDrawer(gander)
	game.AddActor(gander_mv, 4)

	door := gizmo.NewTexture(filepath.Join(resourcePath, "doorstop.png"))
	door_mv := gizmo.NewMover(viewPort, 100, 100, .5)
	door_mv.AddDrawer(door)
	game.AddActor(door_mv, 10)

	game.FixedPalette = fixedPalette
	game.Width = 1280
	game.Height = 720
	game.FrameRate = 50
}
