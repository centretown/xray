package main

import (
	"fmt"

	"github.com/centretown/xray/tools"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	baseInterval = .02
	screenWidth  = 1280
	screenHeight = 720
	fps          = 30
	captureFps   = 25
)

func main() {
	gs := setup()
	loop(gs)
}

func setup() *Game {

	rl.SetTraceLogLevel(rl.LogWarning)
	rl.InitWindow(screenWidth, screenHeight, "2d")
	rl.SetWindowState(rl.FlagWindowResizable)

	gs := NewGameState(fps)
	viewPort := gs.GetViewPort()

	hole := tools.NewPicture("polar.png")
	bouncer := tools.NewBouncer(hole.Rect(), viewPort, 10, 10, 10)
	gs.AddActor(hole, bouncer, 6)

	ball := tools.NewBall(20, rl.Green)
	bouncer = tools.NewBouncer(ball.Rect(), viewPort, 200, 100, 0)
	gs.AddActor(ball, bouncer, 1)

	head := tools.NewPicture("head_300.png")
	bouncer = tools.NewBouncer(head.Rect(), viewPort, 70, 140, 1.75)
	gs.AddActor(head, bouncer, 8)

	gander := tools.NewPicture("gander.png")
	bouncer = tools.NewBouncer(gander.Rect(), viewPort, 300, 300, 0.5)
	gs.AddActor(gander, bouncer, 4)

	// generate palette and color map for paletted images
	gs.pal, gs.colorMap = createPaletteFromTextures(fixedPal, gs.Actors...)

	rl.SetTargetFPS(gs.fps)
	gs.Refresh(rl.GetTime())
	gs.Dump()

	return gs
}

func loop(gs *Game) {
	defer func() {
		for _, actor := range gs.Actors {
			t, ok := actor.Character.(*tools.Picture)
			if ok {
				fmt.Println("UnloadTexture")
				t.Unload()
			}
		}
		rl.CloseWindow()
	}()

	for !rl.WindowShouldClose() {

		gs.current = rl.GetTime()

		if rl.IsWindowResized() {
			gs.Refresh(gs.current)
		}

		rl.BeginDrawing()

		rl.ClearBackground(BG)

		for _, actor := range gs.Actors {
			actor.Animate(!gs.paused, gs.current)
		}

		gs.DrawStatus()

		rl.EndDrawing()

		gs.ProcessInput()

		if gs.capturing {
			gs.GIFCapture()
		}
	}
}
