package gizmo

import (
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var _ = rand.NewSource(time.Now().UnixMicro())

func (gs *Game) Run() {
	gs.FixedPalette = append(gs.FixedPalette, gs.BackGround)

	gs.createPalette()

	rl.SetTraceLogLevel(rl.LogWarning)
	rl.InitWindow(gs.Width, gs.Height, gs.Record.Title)

	for _, txt := range gs.listTextures() {
		txt.Load()
	}

	defer func() {
		for _, actor := range gs.Actors() {
			t, ok := actor.GetDrawer().(*Texture)
			if ok {
				t.Unload()
			}
		}
		for _, dr := range gs.Drawers() {
			t, ok := dr.(*Texture)
			if ok {
				t.Unload()
			}
		}
		rl.CloseWindow()
	}()

	if !gs.FixedSize {
		rl.SetWindowState(rl.FlagWindowResizable)
	}
	rl.SetTargetFPS(gs.FrameRate)
	gs.Current = rl.GetTime()
	gs.Refresh(gs.Current)

	for !rl.WindowShouldClose() {

		gs.Current = rl.GetTime()

		if rl.IsWindowResized() {
			gs.Refresh(gs.Current)
		}

		rl.BeginDrawing()

		rl.ClearBackground(gs.BackGround)

		for _, dr := range gs.Drawers() {
			dr.Draw(rl.Vector3{X: 0, Y: 0, Z: 0})
		}

		for _, actor := range gs.Actors() {
			actor.Move(!gs.Paused, gs.Current)
		}

		gs.DrawStatus()

		rl.EndDrawing()

		gs.ProcessInput()

		if gs.Capturing {
			gs.gifCapture()
		}
	}
}
