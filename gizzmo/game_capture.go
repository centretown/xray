package gizzmo

import (
	"log"

	"github.com/centretown/xray/capture"
	rl "github.com/gen2brain/raylib-go/raylib"
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
	if item.Capturing {
		log.Println("already capturing...")
		return
	}
	item.CaptureCount = item.CaptureStart
	item.Capturing = true
	fps := rl.GetFPS()
	log.Println("Capturing...", fps)

	if mode == "gif" {
		if fps >= 50 {
			rl.SetTargetFPS(50)
			item.CaptureDelay = 2
		} else {
			rl.SetTargetFPS(25)
			item.CaptureDelay = 4
		}

		// go capture.CaptureGIF(item.stopChan, item.scrChan, item.palette,
		// 	item.CaptureDelay, item.colorMap)
	} else if mode == "mp4" {
		log.Println("BeginCapture mp4")
		go capture.CaptureVideo(item.stopChan, item.scrChan,
			int32(gs.Content.Width), int32(gs.Content.Height), fps)
	}
}

func (gs *Game) screenCapture() {
	item := &gs.Content
	if !item.Capturing {
		log.Println("not supposed to capture")
		return
	}

	imag := rl.LoadImageFromTexture(gs.Content.screen.Texture)
	item.scrChan <- imag
	item.CaptureCount--
	if item.CaptureCount < 0 {
		gs.EndCapture()
	}
}

func (gs *Game) EndCapture() {
	item := &gs.Content
	if !item.Capturing {
		log.Println("nothing to end. not capturing!")
		return
	}
	log.Println("EndCapture")
	item.CaptureCount = -1
	item.Capturing = false
	item.stopChan <- 1
	// close(item.scrChan)
	// item.scrChan = make(chan *image.RGBA)
}
