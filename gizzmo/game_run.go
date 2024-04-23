package gizzmo

import (
	"math/rand"
	"time"

	rl "github.com/centretown/raylib-go/raylib"
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

	rl.InitWindow(int32(content.Width), int32(content.Height), gs.Content.Title)
	rl.SetTraceLogLevel(rl.LogWarning)

	for _, txt := range gs.Content.textureList {
		txt.Load()
	}
	content.captureTexture = rl.LoadRenderTexture(int32(content.Width), int32(content.Height))

	defer gs.unload()

	if !content.FixedSize {
		rl.SetWindowState(rl.FlagWindowResizable)
	}
	rl.SetTargetFPS(int32(content.FrameRate))
	content.Current = rl.GetTime()
	gs.Refresh(content.Current)

	for !rl.WindowShouldClose() {

		content.Current = rl.GetTime()

		if rl.IsWindowResized() {
			rl.UnloadRenderTexture(gs.Content.captureTexture)
			// var height = int32(float32(rl.GetRenderWidth()) / content.aspectRatio)
			content.captureTexture = rl.LoadRenderTexture(
				int32(rl.GetRenderWidth()),
				int32(rl.GetRenderHeight()))
			gs.Refresh(content.Current)
		}

		rl.BeginDrawing()
		rl.BeginTextureMode(content.captureTexture)
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

		rl.EndTextureMode()
		//I'm pretty sure you can give the
		// source rect or the dest rect a negative height and it will flip it.
		tex := gs.Content.captureTexture.Texture
		rl.DrawTexturePro(tex,

			rl.Rectangle{X: 0, Y: 0,
				Width:  float32(tex.Width),
				Height: float32(tex.Height),
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

		gs.DrawStatus()

		rl.EndDrawing()

		gs.ProcessInput()

		if content.capturing {
			gs.captureTexture()
		}
	}
}

func (gs *Game) unload() {
	for _, txt := range gs.Content.textureList {
		txt.Unload()
	}
	rl.UnloadRenderTexture(gs.Content.captureTexture)
	gs.data.Close()
	rl.CloseWindow()

}
