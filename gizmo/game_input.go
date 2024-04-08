package gizmo

import (
	"github.com/centretown/gpads/gpads"
	"github.com/centretown/xray/capture"
	"github.com/centretown/xray/try"
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
	gs.gamepad.BeginPad()
	if gs.Current > gs.nextInput {
		gs.nextInput = gs.Current + gs.InputInterval
		for i := range gs.gamepad.GetPadCount() {
			gs.CheckPad(i)
		}
	}
}

func (gs *Game) CheckPad(i int32) {
	var is_multiply, down bool
	// gs.gamepad.GetPadButtonPressed()

	for b := range PAD_STATES {
		switch b {
		case TIMES_TEN:
			is_multiply = gs.gamepad.IsGamepadButtonDown(i, gpads.RL_LeftTrigger1)
			// rl.GamepadButtonLeftTrigger1)

		case FPS_INC:
			if gs.gamepad.IsGamepadButtonDown(i, gpads.RL_LeftFaceUp) {
				gs.FrameRate += try.Or[int32](is_multiply, 1, 10)
				rl.SetTargetFPS(gs.FrameRate)
			}
		case FPS_DEC:
			if gs.gamepad.IsGamepadButtonDown(i, gpads.RL_LeftFaceDown) {
				gs.FrameRate -= try.Or[int32](is_multiply, 1, 10)
				if gs.FrameRate < 5 {
					gs.FrameRate = 5
				}
				rl.SetTargetFPS(gs.FrameRate)
			}

		case CAPTURE_COUNT_INC:
			if gs.gamepad.IsGamepadButtonDown(i, gpads.RL_RightFaceUp) {
				gs.captureStart += try.Or(is_multiply, 1, 10)
			}
		case CAPTURE_COUNT_DEC:
			if gs.gamepad.IsGamepadButtonDown(i, gpads.RL_RightFaceDown) {
				gs.captureStart -= try.Or(is_multiply, 1, 10)
				if gs.captureStart < 1 {
					gs.captureStart = 1
				}
			}

		case CAPTURE_GIF:
			down = gs.gamepad.IsGamepadButtonDown(i, gpads.RL_MiddleLeft)
			if down && gs.Capturing {
				gs.EndGIFCapture()
			} else if down {
				gs.BeginGIFCapture()
			}

		case CAPTURE_PNG:
			if gs.gamepad.IsGamepadButtonDown(i, gpads.RL_MiddleRight) {
				capture.CapturePNG(gs.path, rl.LoadImageFromScreen().ToImage())
			}
		case PAUSED:
			if gs.gamepad.IsGamepadButtonDown(i, gpads.RL_RightFaceLeft) {
				gs.Paused = !gs.Paused
				if !gs.Paused {
					gs.Refresh(gs.Current)
				}
			}

		}
	}
}
