package gizzmo

import (
	"log"

	rl "github.com/centretown/raylib-go/raylib"
	"github.com/centretown/xray/capture"
)

func (gs *Game) BeginCapture(mode string) {

	content := &gs.Content
	if content.capturing {
		log.Println("already capturing...")
		return
	}

	if mode == "mp4" {
		content.captureCount = 0
		content.captureTotal = int64(content.CaptureDuration) * content.FrameRate
		if content.captureTotal == 0 {
			log.Fatal("content.captureTotal == 0")
		}
		content.capturing = true
		content.captureEnd = content.currentTime + content.CaptureDuration
		log.Println("Capturing mp4...", content.FrameRate)
		go capture.CaptureVideo(content.captureStop, content.captureSource,
			int32(content.Width), int32(content.Height), int32(content.FrameRate))
	}
}

func (gs *Game) captureTexture() {
	content := &gs.Content
	if !content.capturing {
		log.Fatalln("not supposed to be capturing")
		return
	}

	if content.captureCount >= content.captureTotal {
		gs.EndCapture()
		return
	}

	content.captureImage = rl.LoadImageFromTexture(content.renderTexture.Texture)
	content.captureSource <- gs.Content.captureImage
	content.captureCount++
}

func (gs *Game) EndCapture() {
	item := &gs.Content
	if !item.capturing {
		log.Println("nothing to end. not capturing!")
		return
	}
	log.Println("EndCapture")
	item.captureStop <- 1
	item.capturing = false
}
