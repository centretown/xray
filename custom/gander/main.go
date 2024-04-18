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

	hole := gizzmo.NewTexture("polar.png", 6)
	hole_mv := gizzmo.NewTracker(vp, 5, 5, 5)
	hole_mv.AddDrawer(hole)
	game.AddActor(hole_mv, hole.GetDepth())

	moon := gizzmo.NewTexture("moon-solo-300.png", 5)
	moon_mv := gizzmo.NewTracker(vp, 10, 10, 5)
	moon_mv.AddDrawer(moon)
	game.AddActor(moon_mv, moon.GetDepth())

	ball := gizzmo.NewEllipse(gizzmo.Red, 20, 15, 4)
	ball_mv := gizzmo.NewTracker(vp, 200, 100, 0)
	ball_mv.AddDrawer(ball)
	game.AddActor(ball_mv, ball.GetDepth())

	head := gizzmo.NewTexture("head_300.png", 1)
	head_mv := gizzmo.NewTracker(vp, 70, 140, 1.75)
	head_mv.AddDrawer(head)
	game.AddActor(head_mv, head.GetDepth())

	gander := gizzmo.NewTexture(filepath.Join("gander.png"), 2)
	gander_mv := gizzmo.NewTracker(vp, 300, 300, 0.5)
	gander_mv.AddDrawer(gander)
	game.AddActor(gander_mv, gander.GetDepth())

	door := gizzmo.NewTexture(filepath.Join("doorstop_01.png"), 3)
	door_mv := gizzmo.NewTracker(vp, 100, 100, .5)
	door_mv.AddDrawer(door)
	game.AddActor(door_mv, door.GetDepth())

	game.Content.FrameRate = 30
}
