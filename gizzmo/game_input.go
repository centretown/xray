package gizzmo

import (
	"github.com/centretown/gpads/gpads"
	"github.com/centretown/xray/capture"
	"github.com/centretown/xray/check"
	rl "github.com/gen2brain/raylib-go/raylib"
)

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

func (gs *Game) ProcessInput() {
	item := &gs.Content
	item.gamepad.BeginPad()
	if item.Current > item.nextInput {
		item.nextInput = item.Current + item.InputInterval
		for i := range item.gamepad.GetPadCount() {
			gs.CheckPad(i)
		}
	}

	for _, ch := range gs.Children() {
		t, ok := ch.(Inputer)
		if ok {
			t.Input()
		}
	}

	item.FramesCounter++
}

func (gs *Game) CheckPad(i int32) {
	item := &gs.Content
	var multiply_by_ten, down bool
	// gs.gamepad.GetPadButtonPressed()

	for b := range PAD_STATES {
		switch b {
		case TIMES_TEN:
			multiply_by_ten = item.gamepad.IsGamepadButtonDown(i, gpads.RL_LeftTrigger1)
			// rl.GamepadButtonLeftTrigger1)

		case FPS_INC:
			if item.gamepad.IsGamepadButtonDown(i, gpads.RL_LeftFaceUp) {
				item.FrameRate += check.AsOr[int32](multiply_by_ten, 10, 1)
				rl.SetTargetFPS(item.FrameRate)
			}
		case FPS_DEC:
			if item.gamepad.IsGamepadButtonDown(i, gpads.RL_LeftFaceDown) {
				item.FrameRate -= check.AsOr[int32](multiply_by_ten, 10, 1)
				if item.FrameRate < 5 {
					item.FrameRate = 5
				}
				rl.SetTargetFPS(item.FrameRate)
			}

		case CAPTURE_COUNT_INC:
			if item.gamepad.IsGamepadButtonDown(i, gpads.RL_RightFaceUp) {
				item.CaptureStart += check.AsOr(multiply_by_ten, 10, 1)
			}
		case CAPTURE_COUNT_DEC:
			if item.gamepad.IsGamepadButtonDown(i, gpads.RL_RightFaceDown) {
				item.CaptureStart -= check.AsOr(multiply_by_ten, 10, 1)
				if item.CaptureStart < 1 {
					item.CaptureStart = 1
				}
			}

		case CAPTURE_GIF:
			down = item.gamepad.IsGamepadButtonDown(i, gpads.RL_MiddleLeft)
			if down && item.Capturing {
				gs.EndGIFCapture()
			} else if down {
				gs.BeginGIFCapture()
			}

		case CAPTURE_PNG:
			if item.gamepad.IsGamepadButtonDown(i, gpads.RL_MiddleRight) {
				capture.CapturePNG("", rl.LoadImageFromScreen().ToImage())
			}
		case PAUSED:
			if item.gamepad.IsGamepadButtonDown(i, gpads.RL_RightFaceLeft) {
				item.Paused = !item.Paused
				if !item.Paused {
					gs.Refresh(item.Current)
				}
			}

		}
	}
}
