package gizzmo

import (
	"log"

	rl "github.com/centretown/raylib-go/raylib"
	"github.com/centretown/xray/capture"
)

// func (gs *Game) CanCapture() bool {
// 	item := &gs.Content
// 	canCapture := item.Current >= item.previousCapture+item.CaptureInterval
// 	moveFloat := check.As[float64](canCapture)
// 	item.previousCapture = moveFloat*item.CaptureInterval + moveFloat*item.Current
// 	return canCapture
// }

func (gs *Game) BeginCapture(mode string) {

	item := &gs.Content
	if item.capturing {
		log.Println("already capturing...")
		return
	}

	if mode == "mp4" {
		fps := gs.Content.FrameRate
		item.captureCount = item.captureFrames
		item.capturing = true
		log.Println("Capturing mp4...", fps)
		go capture.CaptureVideo(item.captureStop, item.captureSource,
			int32(gs.Content.Width), int32(gs.Content.Height), int32(fps))
	}
}

func (gs *Game) captureTexture() {
	content := &gs.Content
	if !content.capturing {
		log.Println("not supposed to capture")
		return
	}

	content.captureImage = rl.LoadImageFromTexture(content.captureTexture.Texture)
	content.captureSource <- gs.Content.captureImage
	content.captureCount--
	if content.captureCount < 0 {
		gs.EndCapture()
	}
}

func (gs *Game) EndCapture() {
	item := &gs.Content
	if !item.capturing {
		log.Println("nothing to end. not capturing!")
		return
	}
	log.Println("EndCapture")
	item.captureCount = -1
	item.capturing = false
	item.captureStop <- 1
}
