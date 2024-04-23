package gizzmo

import (
	"github.com/centretown/gpads/gpads"
	rl "github.com/centretown/raylib-go/raylib"
	"github.com/centretown/xray/check"
)

const (
	TIMES_TEN = iota
	FPS_INC
	FPS_DEC
	CAPTURE_COUNT_INC
	CAPTURE_COUNT_DEC
	CAPTURE_GIF
	CAPTURE_PNG
	PAUSE_PLAY
	CAPTURE_MP4
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

	for _, child := range gs.Children() {
		t, ok := child.(Inputer)
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
				item.FrameRate += check.AsOr[int64](multiply_by_ten, 10, 1)
				rl.SetTargetFPS(int32(item.FrameRate))
			}
		case FPS_DEC:
			if item.gamepad.IsGamepadButtonDown(i, gpads.RL_LeftFaceDown) {
				item.FrameRate -= check.AsOr[int64](multiply_by_ten, 10, 1)
				if item.FrameRate < 5 {
					item.FrameRate = 5
				}
				rl.SetTargetFPS(int32(item.FrameRate))
			}

		case CAPTURE_COUNT_INC:
			if item.gamepad.IsGamepadButtonDown(i, gpads.RL_RightFaceUp) {
				item.captureFrames += check.AsOr(multiply_by_ten, int64(10), 1)
			}
		case CAPTURE_COUNT_DEC:
			if item.gamepad.IsGamepadButtonDown(i, gpads.RL_RightFaceDown) {
				item.captureFrames -= check.AsOr(multiply_by_ten, int64(10), 1)
				if item.captureFrames < 1 {
					item.captureFrames = 1
				}
			}

		case CAPTURE_GIF:
		// 	down = item.gamepad.IsGamepadButtonDown(i, gpads.RL_MiddleLeft)
		// 	if down && item.Capturing {
		// 		gs.EndCapture()
		// 	} else if down {
		// 		gs.BeginCapture("gif")
		// 	}

		case CAPTURE_PNG:
			// if item.gamepad.IsGamepadButtonDown(i, gpads.RL_MiddleRight) {
			// 	capture.CapturePNG(rl.LoadImageFromScreen().ToImage())
			// }
		case PAUSE_PLAY:
			if item.gamepad.IsGamepadButtonDown(i, gpads.RL_RightFaceLeft) {
				item.Paused = !item.Paused
				if !item.Paused {
					gs.Refresh(item.Current)
				}
			}

		case CAPTURE_MP4:
			down = item.gamepad.IsGamepadButtonDown(i, gpads.RL_MiddleLeft)
			if down && item.capturing {
				gs.EndCapture()
			} else if down {
				gs.BeginCapture("mp4")
			}
		}
	}
}
