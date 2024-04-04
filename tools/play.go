package tools

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (gs *Game) Run() {

	defer func() {
		for _, actor := range gs.Actors() {
			t, ok := actor.Drawer().(*Picture)
			if ok {
				fmt.Println("UnloadTexture")
				t.Unload()
			}
		}
		rl.CloseWindow()
	}()

	for !rl.WindowShouldClose() {

		gs.Current = rl.GetTime()

		if rl.IsWindowResized() {
			gs.Refresh(gs.Current)
		}

		rl.BeginDrawing()

		rl.ClearBackground(gs.backGround)

		for _, actor := range gs.Actors() {
			actor.Move(!gs.Paused, gs.Current)
		}

		gs.DrawStatus()

		rl.EndDrawing()

		gs.ProcessInput()

		if gs.Capturing {
			gs.GIFCapture()
		}
	}
}
