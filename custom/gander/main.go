package main

import (
	"path/filepath"

	"github.com/centretown/xray/builder"
	"github.com/centretown/xray/gizmo"
)

func main() {
	game, install, err := builder.Build(build_gander)
	if err == nil && !install {
		game.Run()
	}
}

func build_gander(game *gizmo.Game) {
	viewPort := game.SetViewPort(1600, 800)
	vp := viewPort.ToFloat32()

	hole := gizmo.NewTexture("polar.png")
	hole_mv := gizmo.NewTracker(vp, 5, 5, 5)
	hole_mv.AddDrawer(hole)
	game.AddActor(hole_mv, 6)

	moon := gizmo.NewTexture("moon-solo-300.png")
	moon_mv := gizmo.NewTracker(vp, 10, 10, 5)
	moon_mv.AddDrawer(moon)
	game.AddActor(moon_mv, 6)

	ball := gizmo.NewEllipse(gizmo.Red, 20, 15)
	ball_mv := gizmo.NewTracker(vp, 200, 100, 0)
	ball_mv.AddDrawer(ball)
	game.AddActor(ball_mv, 5)

	head := gizmo.NewTexture("head_300.png")
	head_mv := gizmo.NewTracker(vp, 70, 140, 1.75)
	head_mv.AddDrawer(head)
	game.AddActor(head_mv, 8)

	gander := gizmo.NewTexture(filepath.Join("gander.png"))
	gander_mv := gizmo.NewTracker(vp, 300, 300, 0.5)
	gander_mv.AddDrawer(gander)
	game.AddActor(gander_mv, 4)

	door := gizmo.NewTexture(filepath.Join("doorstop_01.png"))
	door_mv := gizmo.NewTracker(vp, 100, 100, .5)
	door_mv.AddDrawer(door)
	game.AddActor(door_mv, 10)

	game.Content.FrameRate = 30
}
