package gizmo

import (
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var _ = rand.NewSource(time.Now().UnixMicro())

func (gs *Game) Run() {
	content := &gs.Content
	content.FixedPalette = append(content.FixedPalette, content.BackGround)
	gs.createPalette()
	rl.InitWindow(content.Width, content.Height, gs.Record.Class)
	for _, txt := range gs.listTextures() {
		txt.Load()
	}
	rl.SetTraceLogLevel(rl.LogWarning)

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

	if !content.FixedSize {
		rl.SetWindowState(rl.FlagWindowResizable)
	}
	rl.SetTargetFPS(content.FrameRate)
	content.Current = rl.GetTime()
	gs.Refresh(content.Current)

	for !rl.WindowShouldClose() {

		content.Current = rl.GetTime()

		if rl.IsWindowResized() {
			gs.Refresh(content.Current)
		}

		rl.BeginDrawing()

		rl.ClearBackground(content.BackGround)

		for _, dr := range gs.Drawers() {
			dr.Draw(rl.Vector4{X: 0, Y: 0, Z: 0})
		}

		for _, actor := range gs.Actors() {
			actor.Move(!content.Paused, content.Current)
		}

		gs.DrawStatus()

		rl.EndDrawing()

		gs.ProcessInput()

		if content.Capturing {
			gs.gifCapture()
		}
	}
}
