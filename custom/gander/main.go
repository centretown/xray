package main

import (
	"path/filepath"

	"github.com/centretown/xray/builder"
	"github.com/centretown/xray/gizzmo"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	game, install, err := builder.Build(build_gander)
	if err == nil && !install {
		game.Run()
	}
}

func build_gander(game *gizzmo.Game) {
	viewPort := game.SetViewPort(1600, 800)
	vf := viewPort.ToFloat32()
	vp := rl.Vector4{X: vf.Width, Y: vf.Height, Z: float32(gizzmo.Deepest)}

	game.Content.Title = "Gander"
	game.Content.Author = "Dave"
	game.Content.Description = "Gander is a testing game. It implements Ellipse, Tracker, Texture"
	game.Content.Instructions = "Nothing to do right now."

	hole := gizzmo.NewTexture("polar.png", 6)
	hole_mv := gizzmo.NewTracker(rl.Vector4{X: vp.X, Y: vp.Y, Z: 100},
		// rates
		rl.Vector4{X: 3, Y: 3, Z: 0, W: 5},
		// minimums
		rl.Vector3{X: 0, Y: 0, Z: 90},
		// maximums
		rl.Vector3{X: vp.X, Y: vp.Y, Z: 100})
	hole_mv.AddDrawer(hole)
	game.AddActor(hole_mv, hole.GetDepth())

	moon := gizzmo.NewTexture("moon-solo-300.png", 5)
	moon_mv := gizzmo.NewTracker(rl.Vector4{X: vp.X, Y: vp.Y, Z: 50},
		// rates
		rl.Vector4{X: 30, Y: 30, Z: 0, W: 5},
		// minimums
		rl.Vector3{X: 0, Y: 0, Z: 30},
		// maximums
		rl.Vector3{X: vp.X, Y: vp.Y, Z: 50})
	moon_mv.AddDrawer(moon)
	game.AddActor(moon_mv, moon.GetDepth())

	ball := gizzmo.NewEllipse(gizzmo.Red, 80, 35, 4)
	ball_mv := gizzmo.NewTracker(rl.Vector4{X: vp.X, Y: vp.Y, Z: 28},
		// rates
		rl.Vector4{X: 200, Y: 100, Z: 0, W: 0},
		// minimums
		rl.Vector3{X: 0, Y: 0, Z: 28},
		// maximums
		rl.Vector3{X: vp.X, Y: vp.Y, Z: 58})
	ball_mv.AddDrawer(ball)
	game.AddActor(ball_mv, ball.GetDepth())

	head := gizzmo.NewTexture("head_300.png", 1)
	head_mv := gizzmo.NewTracker(rl.Vector4{X: vp.X, Y: vp.Y, Z: 28},
		// rates
		rl.Vector4{X: 70, Y: 140, Z: 20, W: 1.75},
		// minimums
		rl.Vector3{X: 0, Y: 0, Z: 1},
		// maximums
		rl.Vector3{X: vp.X, Y: vp.Y, Z: 80})
	head_mv.AddDrawer(head)
	game.AddActor(head_mv, head.GetDepth())

	gander := gizzmo.NewTexture(filepath.Join("gander.png"), 2)
	gander_mv := gizzmo.NewTracker(rl.Vector4{X: vp.X, Y: vp.Y, Z: 28},
		// rates
		rl.Vector4{X: 300, Y: 300, Z: 0, W: 0.5},
		// minimums
		rl.Vector3{X: 0, Y: 0, Z: 28},
		// maximums
		rl.Vector3{X: vp.X, Y: vp.Y, Z: 58})
	gander_mv.AddDrawer(gander)
	game.AddActor(gander_mv, gander.GetDepth())

	door := gizzmo.NewTexture(filepath.Join("doorstop_01.png"), 3)
	door_mv := gizzmo.NewTracker(rl.Vector4{X: vp.X, Y: vp.Y, Z: 28},
		// rates
		rl.Vector4{X: 100, Y: 100, Z: 0, W: 0.5},
		// minimums
		rl.Vector3{X: 0, Y: 0, Z: 28},
		// maximums
		rl.Vector3{X: vp.X, Y: vp.Y, Z: 58})
	door_mv.AddDrawer(door)
	game.AddActor(door_mv, door.GetDepth())

	game.Content.FrameRate = 30
}
