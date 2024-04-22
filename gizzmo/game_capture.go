package gizzmo

import (
	"image"
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
	if item.Capturing {
		log.Println("already capturing...")
		return
	}

	// if mode == "gif" {
	// 	if fps >= 50 {
	// 		rl.SetTargetFPS(50)
	// 		item.CaptureDelay = 2
	// 	} else {
	// 		rl.SetTargetFPS(25)
	// 		item.CaptureDelay = 4
	// 	}

	// 	// go capture.CaptureGIF(item.stopChan, item.scrChan, item.palette,
	// 	// 	item.CaptureDelay, item.colorMap)
	if mode == "mp4" {
		fps := rl.GetFPS()
		item.CaptureCount = item.CaptureStart
		item.Capturing = true
		log.Println("Capturing mp4...", fps)
		go capture.CaptureVideo(item.stopChan, item.scrChan,
			int32(gs.Content.Width), int32(gs.Content.Height), fps)
	}
}

func (gs *Game) screenCapture() {
	content := &gs.Content
	if !content.Capturing {
		log.Println("not supposed to capture")
		return
	}

	tex := content.screen.Texture
	content.captureImage = image.NewRGBA(image.Rectangle{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: int(content.screen.Texture.Width),
			Y: int(content.screen.Texture.Height)}})

	scr := rl.LoadImageFromTexture(tex)
	scr.ToImageEx(content.captureImage)
	// rl.UnloadTexture(tex)

	content.scrChan <- gs.Content.captureImage
	content.CaptureCount--
	if content.CaptureCount < 0 {
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
	close(item.scrChan)
	item.scrChan = make(chan *image.RGBA)
}
