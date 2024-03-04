package main

import (
	"image/color"
	"xray/tools"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var colors = []color.RGBA{
	rl.White,
	rl.Blue,
	rl.Yellow,
	rl.Red,
	rl.White,
	rl.Red,
	rl.White,
	rl.Yellow,
	rl.Lime,
	rl.DarkGreen,
}

func main() {
	runr := tools.NewRunner(1280, 720, 60)
	viewPort := runr.GetViewPort()

	runr.Add(tools.NewBall(60, colors), tools.NewBouncer(viewPort, 60, 60), 0)
	runr.Add(tools.NewBall(40, colors[6:]), tools.NewBouncer(viewPort, 40, 40), 1)
	runr.Add(tools.NewBall(30, colors[2:]), tools.NewBouncer(viewPort, 30, 30), 2)
	runr.Add(tools.NewBall(20, colors[4:]), tools.NewBouncer(viewPort, 20, 20), 3)
	runr.Run2d()
}
