package main

import (
	"path/filepath"

	"github.com/centretown/xray/builder"
	"github.com/centretown/xray/gizzmo"
)

func main() {
	game, install, err := builder.Build(build_gander)
	if err == nil && !install {
		game.Run()
	}
}

func build_gander(game *gizzmo.Game) {
	viewPort := game.SetViewPort(1600, 800)
	vp := viewPort.ToFloat32()

	game.Content.Title = "Gander"
	game.Content.Author = "Dave"
	game.Content.Description = "Gander is a testing game. It implements Ellipse, Tracker, Texture"
	game.Content.Instructions = "Nothing to do right now."

	hole := gizzmo.NewTexture("polar.png")
	hole_mv := gizzmo.NewTracker(vp, 5, 5, 5)
	hole_mv.AddDrawer(hole)
	game.AddActor(hole_mv, 6)

	moon := gizzmo.NewTexture("moon-solo-300.png")
	moon_mv := gizzmo.NewTracker(vp, 10, 10, 5)
	moon_mv.AddDrawer(moon)
	game.AddActor(moon_mv, 6)

	ball := gizzmo.NewEllipse(gizzmo.Red, 20, 15)
	ball_mv := gizzmo.NewTracker(vp, 200, 100, 0)
	ball_mv.AddDrawer(ball)
	game.AddActor(ball_mv, 5)

	head := gizzmo.NewTexture("head_300.png")
	head_mv := gizzmo.NewTracker(vp, 70, 140, 1.75)
	head_mv.AddDrawer(head)
	game.AddActor(head_mv, 8)

	gander := gizzmo.NewTexture(filepath.Join("gander.png"))
	gander_mv := gizzmo.NewTracker(vp, 300, 300, 0.5)
	gander_mv.AddDrawer(gander)
	game.AddActor(gander_mv, 4)

	door := gizzmo.NewTexture(filepath.Join("doorstop_01.png"))
	door_mv := gizzmo.NewTracker(vp, 100, 100, .5)
	door_mv.AddDrawer(door)
	game.AddActor(door_mv, 10)

	game.Content.FrameRate = 30
}
