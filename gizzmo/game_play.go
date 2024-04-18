package gizzmo

import (
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var _ = rand.NewSource(time.Now().UnixMicro())

func (gs *Game) Run() {
	gs.BuildLists()

	var (
		content   = &gs.Content
		depthList = gs.Content.depthList
		drawer    Drawer
		mover     Mover
		isMover   bool
	)
	content.FixedPalette = append(content.FixedPalette, content.BackGround)

	gs.createPalette()

	rl.InitWindow(content.Width, content.Height, gs.Content.Title)
	rl.SetTraceLogLevel(rl.LogWarning)

	for _, txt := range gs.Content.textureList {
		txt.Load()
	}

	defer func() {
		for _, txt := range gs.Content.textureList {
			txt.Unload()
		}
		gs.data.Close()
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

		// iterate from deepest to least shallowest
		depthList = gs.SortDepthList()
		for i := len(depthList) - 1; i >= 0; i-- {
			drawer = depthList[i].Drawer
			mover, isMover = drawer.(Mover)
			if isMover {
				mover.Move(!content.Paused, content.Current)
			} else {
				drawer.Draw(rl.Vector4{X: 0, Y: 0, Z: 0})
			}
		}

		gs.DrawStatus()

		rl.EndDrawing()

		gs.ProcessInput()

		if content.Capturing {
			gs.gifCapture()
		}
	}
}
