package gizmo

import (
	"log"

	"github.com/centretown/xray/capture"
	"github.com/centretown/xray/try"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (gs *Game) CanCapture() bool {
	canCapture := gs.Current >= gs.previousCapture+gs.CaptureInterval
	moveFloat := try.As[float64](canCapture)
	gs.previousCapture = moveFloat*gs.CaptureInterval + moveFloat*gs.Current
	return canCapture
}

func (gs *Game) BeginGIFCapture() {
	if gs.Capturing {
		log.Println("already capturing...")
		return
	}
	gs.CaptureCount = gs.captureStart
	gs.Capturing = true

	fps := rl.GetFPS()
	if fps >= 50 {
		rl.SetTargetFPS(50)
		gs.captureDelay = 2
	} else {
		rl.SetTargetFPS(25)
		gs.captureDelay = 4
	}

	go capture.CaptureGIF(gs.path, gs.stopChan, gs.scrChan, gs.palette,
		gs.captureDelay, gs.colorMap)
}

func (gs *Game) gifCapture() {
	if !gs.Capturing {
		log.Println("not supposed to capture")
		return
	}

	gs.scrChan <- rl.LoadImageFromScreen().ToImage()
	gs.CaptureCount--
	if gs.CaptureCount < 0 {
		gs.EndGIFCapture()
	}
}

func (gs *Game) EndGIFCapture() {
	if !gs.Capturing {
		log.Println("nothing to end. not capturing!")
		return
	}
	log.Println("end capturing!")
	gs.CaptureCount = -1
	gs.Capturing = false
	gs.stopChan <- 1
}
