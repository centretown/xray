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

	rl.InitWindow(int32(content.Width), int32(content.Height), gs.Content.Title)
	rl.SetTraceLogLevel(rl.LogWarning)

	for _, txt := range gs.Content.textureList {
		txt.Load()
	}
	content.screen = rl.LoadRenderTexture(int32(content.Width), int32(content.Height))

	defer func() {
		for _, txt := range gs.Content.textureList {
			txt.Unload()
		}
		rl.UnloadRenderTexture(gs.Content.screen)
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
			rl.UnloadRenderTexture(gs.Content.screen)
			content.screen = rl.LoadRenderTexture(int32(rl.GetRenderWidth()),
				int32(rl.GetRenderHeight()))
			gs.Refresh(content.Current)
		}

		rl.BeginDrawing()
		rl.BeginTextureMode(gs.Content.screen)
		rl.ClearBackground(content.BackGround)

		// iterate from deepest to shallowest
		depthList = gs.SortDepthList()
		for i := len(depthList) - 1; i >= 0; i-- {
			drawer = depthList[i].Drawer
			mover, isMover = drawer.(Mover)
			if isMover {
				mover.Move(!content.Paused, content.Current)
			} else {
				drawer.Draw(rl.Vector4{X: 0, Y: 0,
					Z: depthList[i].GetDepth()})
			}
		}

		gs.DrawStatus()

		rl.EndTextureMode()
		//I'm pretty sure you can give the
		// source rect or the dest rect a negative height and it will flip it.

		rl.DrawTexturePro(gs.Content.screen.Texture,

			rl.Rectangle{X: 0, Y: 0,
				Width:  gs.Content.Width,
				Height: -gs.Content.Height,
			},
			rl.Rectangle{X: 0, Y: 0,
				Width:  gs.Content.Width,
				Height: gs.Content.Height,
			},
			rl.Vector2{
				X: 0,
				Y: 0,
				// X: gs.Content.Width / 2,
				// Y: gs.Content.Height / 2,
			},
			1, White)
		rl.EndDrawing()

		gs.ProcessInput()

		if content.Capturing {
			gs.screenCapture()
		}
	}
}
