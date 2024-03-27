package main

import (
	"fmt"
	"image"
	"image/color"

	"github.com/centretown/xray/capture"
	"github.com/centretown/xray/tools"

	"github.com/centretown/gpads/gpads"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameState struct {
	next         float64
	capturing    bool
	paused       bool
	captureStart int
	captureCount int
	stopChan     chan int
	scrChan      chan image.Image

	current  float64
	previous float64
	interval float64
	base     float64
	can_move int32
	pads     *gpads.GPads
	actors   []*tools.TextureDrawer
	pal      color.Palette
	colorMap map[color.Color]uint8
}

func NewGameState() *GameState {
	gs := &GameState{
		stopChan:     make(chan int),
		scrChan:      make(chan image.Image),
		current:      rl.GetTime(),
		previous:     rl.GetTime(),
		interval:     float64(rl.GetFrameTime()),
		base:         baseInterval,
		can_move:     0,
		pads:         gpads.NewGPads(),
		actors:       make([]*tools.TextureDrawer, 0),
		captureStart: 100,
	}
	return gs
}

func (gs *GameState) ProcessInput() {
	gs.pads.BeginPad()

	if gs.current > gs.next {
		gs.next = gs.current + .25

		for i := range gs.pads.GetStickCount() {

			if gs.pads.IsPadButtonDown(i, rl.GamepadButtonRightFaceLeft) {
				gs.paused = !gs.paused
				return
			}

			if gs.pads.IsPadButtonDown(i, rl.GamepadButtonRightFaceUp) {
				gs.captureStart++
				return
			}

			if gs.pads.IsPadButtonDown(i, rl.GamepadButtonRightFaceDown) {
				if gs.captureStart > 1 {
					gs.captureStart--
				}
				return
			}

			if gs.pads.IsPadButtonDown(i, rl.GamepadButtonMiddleLeft) {
				if gs.capturing {
					gs.EndGIFCapture()
				} else {
					gs.BeginGIFCapture()
				}
				return
			}

			if gs.pads.IsPadButtonDown(i, rl.GamepadButtonMiddleRight) {
				capture.CapturePNG(rl.LoadImageFromScreen().ToImage())
				return
			}
		}
	}
	gs.next = gs.current
}

func (gs *GameState) BeginGIFCapture() {
	if gs.capturing {
		fmt.Println("already capturing...")
		return
	}
	gs.captureCount = gs.captureStart
	gs.capturing = true
	go capture.CaptureGIF(gs.stopChan, gs.scrChan, gs.pal, gs.interval, gs.colorMap)
}

func (gs *GameState) GIFCapture() {
	if !gs.capturing {
		fmt.Println("not supposed to capture")
		return
	}

	gs.scrChan <- rl.LoadImageFromScreen().ToImage()
	gs.captureCount--
	if gs.captureCount < 0 {
		gs.EndGIFCapture()
	}
}

func (gs *GameState) EndGIFCapture() {
	if !gs.capturing {
		fmt.Println("nothing to end. not capturing!")
		return
	}
	fmt.Println("end capturing!")
	gs.captureCount = -1
	gs.capturing = false
	gs.stopChan <- 1
}

func (gs *GameState) DrawStatus(runr *tools.Runner) {
	mb := runr.GetMessageBox()
	text := fmt.Sprintf("FPS:%3d, Capture Count:%4d",
		rl.GetFPS(), gs.captureStart)
	rl.DrawText(text, mb.X, mb.Y+mb.Height-22, 20, rl.Green)
}
