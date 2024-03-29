package main

import (
	"fmt"
	"image"
	"image/color"

	"github.com/centretown/xray/b2"
	"github.com/centretown/xray/capture"
	"github.com/centretown/xray/tools"

	"github.com/centretown/gpads/gpads"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameState struct {
	next         float64
	capturing    bool
	paused       bool
	captureDelay int
	captureStart int
	captureCount int

	previousCapture float64
	captureInterval float64

	stopChan chan int
	scrChan  chan image.Image

	current  float64
	previous float64
	interval float64
	base     float64
	can_move int32
	pads     *gpads.GPads
	actors   []*tools.Picture
	pal      color.Palette
	colorMap map[color.Color]uint8
	fps      int32
}

func NewGameState(fps int32) *GameState {
	gs := &GameState{
		stopChan:        make(chan int),
		scrChan:         make(chan image.Image),
		current:         rl.GetTime(),
		previous:        rl.GetTime(),
		interval:        float64(rl.GetFrameTime() * 2),
		base:            baseInterval,
		can_move:        0,
		pads:            gpads.NewGPads(),
		actors:          make([]*tools.Picture, 0),
		captureStart:    100,
		captureDelay:    4,
		captureInterval: float64(rl.GetFrameTime()) * 2,
		fps:             fps,
	}
	return gs
}

const (
	PAUSED = iota
	FPS_INC
	FPS_DEC
	CAPTURE_GIF
	CAPTURE_PNG
	CAPTURE_COUNT_INC
	CAPTURE_COUNT_DEC
	TIMES_TEN
	PAD_STATE_COUNT
)

func (gs *GameState) CanCapture() bool {
	canCapture := gs.current >= gs.previousCapture+gs.captureInterval
	moveFloat := b2.To[float64](canCapture)
	gs.previousCapture = moveFloat*gs.captureInterval + moveFloat*gs.current
	return canCapture
}

func (gs *GameState) ProcessInput() {
	gs.pads.BeginPad()

	if gs.current > gs.next {
		gs.next = gs.current + .2
		var state [PAD_STATE_COUNT]bool

		for i := range gs.pads.GetStickCount() {

			state[PAUSED] = gs.pads.IsPadButtonDown(i, rl.GamepadButtonRightFaceLeft)

			state[FPS_INC] = gs.pads.IsPadButtonDown(i, rl.GamepadButtonLeftFaceUp)
			state[FPS_DEC] = gs.pads.IsPadButtonDown(i, rl.GamepadButtonLeftFaceDown)

			state[CAPTURE_GIF] = gs.pads.IsPadButtonDown(i, rl.GamepadButtonMiddleLeft)
			state[CAPTURE_PNG] = gs.pads.IsPadButtonDown(i, rl.GamepadButtonMiddleRight)
			state[CAPTURE_COUNT_INC] = gs.pads.IsPadButtonDown(i, rl.GamepadButtonRightFaceUp)
			state[CAPTURE_COUNT_DEC] = gs.pads.IsPadButtonDown(i, rl.GamepadButtonRightFaceDown)
			state[TIMES_TEN] = gs.pads.IsPadButtonDown(i, rl.GamepadButtonLeftTrigger1)

			if state[PAUSED] {
				gs.paused = !gs.paused
			}

			if state[FPS_INC] {
				gs.fps++
				rl.SetTargetFPS(gs.fps)
			}

			if state[FPS_DEC] {
				gs.fps--
				if gs.fps < 1 {
					gs.fps = 1
				}
				rl.SetTargetFPS(gs.fps)
			}

			if state[CAPTURE_GIF] {
				if gs.capturing {
					gs.EndGIFCapture()
				} else {
					gs.BeginGIFCapture()
				}
			}

			if state[CAPTURE_PNG] {
				capture.CapturePNG(rl.LoadImageFromScreen().ToImage())
			}

			if state[CAPTURE_COUNT_INC] {
				inc := 1 + 9*b2.To[int](state[TIMES_TEN])
				gs.captureStart += inc
			}

			if state[CAPTURE_COUNT_DEC] {
				dec := 1 + 9*b2.To[int](state[TIMES_TEN])
				if gs.captureStart-dec > 1 {
					gs.captureStart -= dec
				} else {
					gs.captureStart = 1
				}
			}
		}

	}
}

func (gs *GameState) BeginGIFCapture() {
	if gs.capturing {
		fmt.Println("already capturing...")
		return
	}
	gs.captureCount = gs.captureStart
	gs.capturing = true
	go capture.CaptureGIF(gs.stopChan, gs.scrChan, gs.pal,
		gs.captureDelay, gs.colorMap)
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
	rl.DrawLine(mb.X, mb.Y, mb.Width, mb.Y, rl.Red)

	text := fmt.Sprintf("FPS:%3d, Capture Count:%4d",
		rl.GetFPS(), gs.captureStart)
	rl.DrawText(text, mb.X, mb.Y+mb.Height-22, 20, rl.Green)

	if gs.capturing {
		rl.DrawText(fmt.Sprintf("Capturing... %4d", gs.captureCount),
			mb.X, mb.Y+32, 20, rl.Green)
	}
}

func (gs *GameState) Dump() {
	fmt.Printf("current=%f previous=%f interval=%f base=%f can_move=%d\n",
		gs.current,
		gs.previous,
		gs.interval,
		gs.base,
		gs.can_move)
}
