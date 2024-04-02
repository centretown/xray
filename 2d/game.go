package main

import (
	"fmt"
	"image"
	"image/color"

	"github.com/centretown/xray/capture"
	"github.com/centretown/xray/tools"
	"github.com/centretown/xray/try"

	"github.com/centretown/gpads/gpads"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	nextInput    float64
	capturing    bool
	paused       bool
	captureDelay int
	captureStart int
	captureCount int

	previousCapture float64
	captureInterval float64

	stopChan chan int
	scrChan  chan image.Image

	current float64
	pads    *gpads.GPads
	Actors  []*tools.Actor

	pal      color.Palette
	colorMap map[color.Color]uint8
	fps      int32
}

func NewGameState(fps int32) *Game {
	gs := &Game{
		stopChan:        make(chan int),
		scrChan:         make(chan image.Image),
		current:         rl.GetTime(),
		pads:            gpads.NewGPads(),
		captureStart:    250,
		captureDelay:    4,
		captureInterval: float64(rl.GetFrameTime()) * 2,
		fps:             fps,
		Actors:          make([]*tools.Actor, 0),
	}
	return gs
}

const (
	TIMES_TEN = iota
	FPS_INC
	FPS_DEC
	CAPTURE_COUNT_INC
	CAPTURE_COUNT_DEC
	CAPTURE_GIF
	CAPTURE_PNG
	PAUSED
	PAD_STATES
)

func (gs *Game) AddActor(d tools.Drawable, a tools.Moveable, after float64) {
	gs.Actors = append(gs.Actors, tools.NewActor(d, a, after))
}

func (gs *Game) CanCapture() bool {
	canCapture := gs.current >= gs.previousCapture+gs.captureInterval
	moveFloat := try.As[float64](canCapture)
	gs.previousCapture = moveFloat*gs.captureInterval + moveFloat*gs.current
	return canCapture
}

func mul10[T int | int32](m bool) T {
	return 1 + 9*try.As[T](m)
}

func (gs *Game) ProcessInput() {
	gs.pads.BeginPad()
	if gs.current > gs.nextInput {
		gs.nextInput = gs.current + .2
		for i := range gs.pads.GetStickCount() {
			gs.CheckPad(i)
		}
	}
}

func (gs *Game) CheckPad(i int) {
	var mul, down bool
	for b := range PAD_STATES {
		switch b {
		case TIMES_TEN:
			mul = gs.pads.IsPadButtonDown(i, rl.GamepadButtonLeftTrigger1)
		case FPS_INC:
			if gs.pads.IsPadButtonDown(i, rl.GamepadButtonLeftFaceUp) {
				gs.fps += mul10[int32](mul)
				rl.SetTargetFPS(gs.fps)
			}
		case FPS_DEC:
			if gs.pads.IsPadButtonDown(i, rl.GamepadButtonLeftFaceDown) {
				gs.fps -= mul10[int32](mul)
				if gs.fps < 5 {
					gs.fps = 5
				}
				rl.SetTargetFPS(gs.fps)
			}
		case CAPTURE_COUNT_INC:
			if gs.pads.IsPadButtonDown(i, rl.GamepadButtonRightFaceUp) {
				gs.captureStart += mul10[int](mul)
			}
		case CAPTURE_COUNT_DEC:
			if gs.pads.IsPadButtonDown(i, rl.GamepadButtonRightFaceDown) {
				gs.captureStart -= mul10[int](mul)
				if gs.captureStart < 1 {
					gs.captureStart = 1
				}
			}
		case CAPTURE_GIF:
			down = gs.pads.IsPadButtonDown(i, rl.GamepadButtonMiddleLeft)
			if down && gs.capturing {
				gs.EndGIFCapture()
			} else if down {
				gs.BeginGIFCapture()
			}
		case CAPTURE_PNG:
			if gs.pads.IsPadButtonDown(i, rl.GamepadButtonMiddleRight) {
				capture.CapturePNG(rl.LoadImageFromScreen().ToImage())
			}
		case PAUSED:
			if gs.pads.IsPadButtonDown(i, rl.GamepadButtonRightFaceLeft) {
				gs.paused = !gs.paused
				if !gs.paused {
					gs.Refresh(gs.current)
				}
			}

		}
	}
}

func (gs *Game) BeginGIFCapture() {
	if gs.capturing {
		fmt.Println("already capturing...")
		return
	}
	gs.captureCount = gs.captureStart
	gs.capturing = true

	fps := rl.GetFPS()
	if fps >= 50 {
		rl.SetTargetFPS(50)
		gs.captureDelay = 2
	} else {
		rl.SetTargetFPS(25)
		gs.captureDelay = 4
	}

	go capture.CaptureGIF(gs.stopChan, gs.scrChan, gs.pal,
		gs.captureDelay, gs.colorMap)
}

func (gs *Game) GIFCapture() {
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

func (gs *Game) EndGIFCapture() {
	if !gs.capturing {
		fmt.Println("nothing to end. not capturing!")
		return
	}
	fmt.Println("end capturing!")
	gs.captureCount = -1
	gs.capturing = false
	gs.stopChan <- 1
}

func (gs *Game) DrawStatus() {
	mb := gs.GetMessageBox()
	rl.DrawLine(mb.X, mb.Y, mb.Width, mb.Y, rl.Red)

	monitor := rl.GetCurrentMonitor()

	text := fmt.Sprintf("FPS:%3d, Monitor:%1d (%4d/%4d %3d), View: %4dx%4d, Capture Count:%4d",
		rl.GetFPS(),
		monitor, rl.GetMonitorWidth(monitor),
		rl.GetMonitorHeight(monitor), rl.GetMonitorRefreshRate(monitor),
		rl.GetScreenWidth(), rl.GetScreenHeight(),
		gs.captureStart)
	rl.DrawText(text, mb.X, mb.Y+mb.Height-22, 16, rl.Green)

	if gs.capturing {
		rl.DrawText(fmt.Sprintf("Capturing... %4d", gs.captureCount),
			mb.X, mb.Y+32, 16, rl.Green)
	}
}

func (gs *Game) Refresh(current float64) {
	viewPort := gs.GetViewPort()
	for _, run := range gs.Actors {
		run.Resize(viewPort, current)
	}
}

const (
	msg_height = 80
	min_width  = 200
	min_height = 280
)

func (gs *Game) GetViewPort() rl.RectangleInt32 {
	rw := rl.GetRenderWidth()
	rh := rl.GetRenderHeight()

	if rw >= min_width && rh >= min_height {
		return rl.RectangleInt32{
			X:      0,
			Y:      0,
			Width:  int32(rw),
			Height: int32(rh - msg_height),
		}
	}

	return rl.RectangleInt32{
		X:      0,
		Y:      0,
		Width:  min_width,
		Height: min_height - msg_height,
	}
}

func (gs *Game) GetMessageBox() (rect rl.RectangleInt32) {
	rw := int32(rl.GetRenderWidth())
	rh := int32(rl.GetRenderHeight())
	rect.X = 0
	rect.Width = rw
	rect.Y = rh - msg_height
	rect.Height = msg_height
	return
}

func (gs *Game) Dump() {
}
