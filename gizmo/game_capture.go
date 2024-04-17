package gizmo

import (
	"log"

	"github.com/centretown/xray/capture"
	"github.com/centretown/xray/check"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (gs *Game) CanCapture() bool {
	item := &gs.Content
	canCapture := item.Current >= item.previousCapture+item.CaptureInterval
	moveFloat := check.As[float64](canCapture)
	item.previousCapture = moveFloat*item.CaptureInterval + moveFloat*item.Current
	return canCapture
}

func (gs *Game) BeginGIFCapture() {
	item := &gs.Content
	if item.Capturing {
		log.Println("already capturing...")
		return
	}
	item.CaptureCount = item.CaptureStart
	item.Capturing = true

	fps := rl.GetFPS()
	if fps >= 50 {
		rl.SetTargetFPS(50)
		item.CaptureDelay = 2
	} else {
		rl.SetTargetFPS(25)
		item.CaptureDelay = 4
	}

	go capture.CaptureGIF("", item.stopChan, item.scrChan, item.palette,
		item.CaptureDelay, item.colorMap)
}

func (gs *Game) gifCapture() {
	item := &gs.Content
	if !item.Capturing {
		log.Println("not supposed to capture")
		return
	}

	item.scrChan <- rl.LoadImageFromScreen().ToImage()
	item.CaptureCount--
	if item.CaptureCount < 0 {
		gs.EndGIFCapture()
	}
}

func (gs *Game) EndGIFCapture() {
	item := &gs.Content
	if !item.Capturing {
		log.Println("nothing to end. not capturing!")
		return
	}
	log.Println("end capturing!")
	item.CaptureCount = -1
	item.Capturing = false
	item.stopChan <- 1
}
