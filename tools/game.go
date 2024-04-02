package tools

import (
	"fmt"
	"image"
	"image/color"

	"github.com/centretown/xray/capture"
	"github.com/centretown/xray/rayl"
	"github.com/centretown/xray/try"

	"github.com/centretown/gpads/gpads"
)

type Game struct {
	Actors []Moveable

	Start         float64
	Current       float64
	FPS           int32
	InputInterval float64

	CaptureCount    int
	CaptureInterval float64
	Capturing       bool
	Paused          bool

	backGround color.RGBA
	palette    color.Palette
	colorMap   map[color.Color]uint8

	nextInput       float64
	captureDelay    int
	captureStart    int
	previousCapture float64

	stopChan chan int
	scrChan  chan image.Image

	pads *gpads.GPads
	rl   rayl.RunLib
}

func NewGame(fps int32, rl rayl.RunLib) *Game {
	gs := &Game{
		rl:              rl,
		stopChan:        make(chan int),
		scrChan:         make(chan image.Image),
		Current:         rl.GetTime(),
		pads:            gpads.NewGPads(),
		captureStart:    250,
		captureDelay:    4,
		CaptureInterval: float64(rl.GetFrameTime()) * 2,
		FPS:             fps,
		Actors:          make([]Moveable, 0),
		InputInterval:   .2,
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

func (gs *Game) AddActor(a Moveable, after float64) {
	gs.Actors = append(gs.Actors, a)
}

func (gs *Game) SetColors(BG color.RGBA, pal color.Palette, m map[color.Color]uint8) {
	gs.backGround = BG
	gs.palette = pal
	gs.colorMap = m
}

func (gs *Game) CanCapture() bool {
	canCapture := gs.Current >= gs.previousCapture+gs.CaptureInterval
	moveFloat := try.As[float64](canCapture)
	gs.previousCapture = moveFloat*gs.CaptureInterval + moveFloat*gs.Current
	return canCapture
}

func (gs *Game) ProcessInput() {
	gs.pads.BeginPad()
	if gs.Current > gs.nextInput {
		gs.nextInput = gs.Current + gs.InputInterval
		for i := range gs.pads.GetStickCount() {
			gs.CheckPad(i)
		}
	}
}

func (gs *Game) CheckPad(i int) {
	var is_multiply, down bool
	for b := range PAD_STATES {
		switch b {
		case TIMES_TEN:
			is_multiply = gs.pads.IsPadButtonDown(i, gpads.RL_LeftTrigger1)
			// rl.GamepadButtonLeftTrigger1)

		case FPS_INC:
			if gs.pads.IsPadButtonDown(i, gpads.RL_LeftFaceUp) {
				gs.FPS += try.Or[int32](is_multiply, 1, 10)
				gs.rl.SetTargetFPS(gs.FPS)
			}
		case FPS_DEC:
			if gs.pads.IsPadButtonDown(i, gpads.RL_LeftFaceDown) {
				gs.FPS -= try.Or[int32](is_multiply, 1, 10)
				if gs.FPS < 5 {
					gs.FPS = 5
				}
				gs.rl.SetTargetFPS(gs.FPS)
			}

		case CAPTURE_COUNT_INC:
			if gs.pads.IsPadButtonDown(i, gpads.RL_RightFaceUp) {
				gs.captureStart += try.Or(is_multiply, 1, 10)
			}
		case CAPTURE_COUNT_DEC:
			if gs.pads.IsPadButtonDown(i, gpads.RL_RightFaceDown) {
				gs.captureStart -= try.Or(is_multiply, 1, 10)
				if gs.captureStart < 1 {
					gs.captureStart = 1
				}
			}

		case CAPTURE_GIF:
			down = gs.pads.IsPadButtonDown(i, gpads.RL_MiddleLeft)
			if down && gs.Capturing {
				gs.EndGIFCapture()
			} else if down {
				gs.BeginGIFCapture()
			}

		case CAPTURE_PNG:
			if gs.pads.IsPadButtonDown(i, gpads.RL_MiddleRight) {
				capture.CapturePNG(gs.rl.LoadImageFromScreen())
			}
		case PAUSED:
			if gs.pads.IsPadButtonDown(i, gpads.RL_RightFaceLeft) {
				gs.Paused = !gs.Paused
				if !gs.Paused {
					gs.Refresh(gs.Current)
				}
			}

		}
	}
}

func (gs *Game) BeginGIFCapture() {
	if gs.Capturing {
		fmt.Println("already capturing...")
		return
	}
	gs.CaptureCount = gs.captureStart
	gs.Capturing = true

	fps := gs.rl.GetFPS()
	if fps >= 50 {
		gs.rl.SetTargetFPS(50)
		gs.captureDelay = 2
	} else {
		gs.rl.SetTargetFPS(25)
		gs.captureDelay = 4
	}

	go capture.CaptureGIF(gs.stopChan, gs.scrChan, gs.palette,
		gs.captureDelay, gs.colorMap)
}

func (gs *Game) GIFCapture() {
	if !gs.Capturing {
		fmt.Println("not supposed to capture")
		return
	}

	gs.scrChan <- gs.rl.LoadImageFromScreen()
	gs.CaptureCount--
	if gs.CaptureCount < 0 {
		gs.EndGIFCapture()
	}
}

func (gs *Game) EndGIFCapture() {
	if !gs.Capturing {
		fmt.Println("nothing to end. not capturing!")
		return
	}
	fmt.Println("end capturing!")
	gs.CaptureCount = -1
	gs.Capturing = false
	gs.stopChan <- 1
}

func (gs *Game) DrawStatus() {
	mb := gs.GetMessageBox()
	gs.rl.DrawLine(mb.X, mb.Y, mb.Width, mb.Y, color.RGBA{255, 0, 0, 255})

	monitor := gs.rl.GetCurrentMonitor()

	text := fmt.Sprintf("FPS:%3d, Monitor:%1d (%4d/%4d %3d), View: %4dx%4d, Capture Count:%4d",
		gs.rl.GetFPS(),
		monitor, gs.rl.GetMonitorWidth(monitor),
		gs.rl.GetMonitorHeight(monitor), gs.rl.GetMonitorRefreshRate(monitor),
		gs.rl.GetScreenWidth(), gs.rl.GetScreenHeight(),
		gs.captureStart)

	yellow := color.RGBA{255, 255, 0, 255}
	gs.rl.DrawText(text, mb.X, mb.Y+mb.Height-22, 16, yellow)

	if gs.Capturing {
		gs.rl.DrawText(fmt.Sprintf("Capturing... %4d", gs.CaptureCount),
			mb.X, mb.Y+32, 16, yellow)
	}
}

func (gs *Game) Refresh(current float64) {
	viewPort := gs.GetViewPort()
	for _, run := range gs.Actors {
		run.Refresh(current, viewPort)
	}
}

const (
	msg_height = 80
	min_width  = 200
	min_height = 280
)

func (gs *Game) GetViewPort() rayl.RectangleInt32 {
	rw := gs.rl.GetRenderWidth()
	rh := gs.rl.GetRenderHeight()

	if rw >= min_width && rh >= min_height {
		return rayl.RectangleInt32{
			X:      0,
			Y:      0,
			Width:  int32(rw),
			Height: int32(rh - msg_height),
		}
	}

	return rayl.RectangleInt32{
		X:      0,
		Y:      0,
		Width:  min_width,
		Height: min_height - msg_height,
	}
}

func (gs *Game) GetMessageBox() (rect rayl.RectangleInt32) {
	rw := int32(gs.rl.GetRenderWidth())
	rh := int32(gs.rl.GetRenderHeight())
	rect.X = 0
	rect.Width = rw
	rect.Y = rh - msg_height
	rect.Height = msg_height
	return
}

func (gs *Game) Dump() {
}
