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
	rl.InitWindow(int32(content.Width), int32(content.Height), gs.Content.Title)
	rl.SetTraceLogLevel(rl.LogWarning)
	gs.BuildNotes()

	for _, txt := range gs.Content.textureList {
		txt.Load()
	}
	content.renderTexture = rl.LoadRenderTexture(int32(content.Width), int32(content.Height))

	defer gs.unload()

	if !content.FixedSize {
		rl.SetWindowState(rl.FlagWindowResizable | rl.FlagBorderlessWindowedMode)
	}
	rl.SetTargetFPS(int32(content.FrameRate))
	content.currentTime = rl.GetTime()
	gs.Refresh(content.currentTime)

	var (
		stopCh     = make(chan int)
		repeatCh   = make(chan float64)
		repeatRate = float64(.30)
	)

	go gs.ProcessInput(repeatRate, repeatCh, stopCh)

	for !rl.WindowShouldClose() {

		content.currentTime = rl.GetTime()
		if rl.IsWindowResized() {
			rl.UnloadRenderTexture(gs.Content.renderTexture)
			content.Width = float32(rl.GetRenderWidth())
			content.Height = float32(rl.GetRenderHeight())

			content.renderTexture = rl.LoadRenderTexture(
				int32(content.Width),
				int32(content.Height))
			gs.Refresh(content.currentTime)
		}

		rl.BeginDrawing()
		rl.BeginTextureMode(content.renderTexture)
		rl.ClearBackground(content.BackGround)

		depthList = gs.SortDepthList()
		for i := len(depthList) - 1; i >= 0; i-- {
			drawer = depthList[i].Drawer
			mover, isMover = drawer.(Mover)
			if isMover {
				mover.Move(!content.paused, content.currentTime)
			} else {
				drawer.Draw(rl.Vector4{X: 0, Y: 0,
					Z: depthList[i].GetDepth()})
			}
		}
		rl.EndTextureMode()

		tex := content.renderTexture.Texture
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
			}, 1, White)

		gs.DrawStatus()
		rl.EndDrawing()

		if !content.paused {
			gs.Refresh(content.currentTime)
		}

		if content.capturing {
			gs.captureTexture()
		}

		if content.beginCapturing {
			gs.BeginCapture("mp4")
		}

		if content.endCapturing {
			gs.EndCapture()
		}
	}

	stopCh <- 1
}

func (gs *Game) unload() {
	for _, txt := range gs.Content.textureList {
		txt.Unload()
	}
	rl.UnloadRenderTexture(gs.Content.renderTexture)
	gs.data.Close()
	rl.CloseWindow()

}

func (gs *Game) Refresh(current float64) {
	viewPort := gs.SetViewPort(float32(rl.GetRenderWidth()),
		float32(rl.GetRenderHeight()))

	for _, mover := range gs.Content.movers {
		mover.Refresh(current, rl.Vector4{
			X: viewPort.Width,
			Y: viewPort.Height})
	}
	for _, drawer := range gs.Content.drawers {
		drawer.Refresh(current, rl.Vector4{
			X: float32(viewPort.Width),
			Y: float32(viewPort.Height)})
	}
}
