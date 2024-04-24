package gizzmo

import (
	"github.com/centretown/gpads/gpads"
	"github.com/centretown/xray/check"
)

const (
	TIMES_TEN = iota
	NEXT_TOKEN
	PREV_TOKEN
	INC_TOKEN
	DEC_TOKEN

	CAPTURE_COUNT_INC
	CAPTURE_COUNT_DEC
	CAPTURE_GIF
	CAPTURE_PNG
	PAUSE_PLAY
	CAPTURE_MP4
	RESIZE_UP
	RESIZE_DOWN
	PAD_STATES
)

func (gs *Game) ProcessInput() {
	item := &gs.Content
	item.gamepad.BeginPad()
	if item.currentTime > item.nextInput {
		item.nextInput = item.currentTime + item.InputInterval
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
}

func (gs *Game) CheckPad(i int32) {
	content := &gs.Content
	var multiply_by_ten, down bool
	// gs.gamepad.GetPadButtonPressed()

	if content.gamepad.IsGamepadButtonDown(i, gpads.RL_MiddleRight) {
		content.commandState = !content.commandState
	}

	if !content.commandState {
		return
	}

	for b := range PAD_STATES {
		switch b {
		case TIMES_TEN:
			multiply_by_ten = content.gamepad.IsGamepadButtonDown(i, gpads.RL_LeftTrigger1)
			// rl.GamepadButtonLeftTrigger1)

		case NEXT_TOKEN:
			// 	content.FrameRate += check.AsOr[int64](multiply_by_ten, 10, 1)
			// 	rl.SetTargetFPS(int32(content.FrameRate))
			length := content.tokens.Length
			if content.gamepad.IsGamepadButtonDown(i, gpads.RL_LeftFaceDown) {
				if content.currentToken+1 < length {
					content.currentToken++
				} else {
					content.currentToken = 0
				}
			}

		case PREV_TOKEN:
			// 	content.FrameRate -= check.AsOr[int64](multiply_by_ten, 10, 1)
			// 	if content.FrameRate < 5 {
			// 		content.FrameRate = 5
			// 	}
			// 	rl.SetTargetFPS(int32(content.FrameRate))
			// }
			length := content.tokens.Length
			if content.gamepad.IsGamepadButtonDown(i, gpads.RL_LeftFaceUp) {
				if content.currentToken-1 >= 0 {
					content.currentToken--
				} else {
					content.currentToken = length - 1
				}
			}

		case INC_TOKEN:
			if content.gamepad.IsGamepadButtonDown(i, gpads.RL_LeftFaceRight) {
			}
		case DEC_TOKEN:
			if content.gamepad.IsGamepadButtonDown(i, gpads.RL_LeftFaceLeft) {
			}

		case CAPTURE_COUNT_INC:
			if content.gamepad.IsGamepadButtonDown(i, gpads.RL_RightFaceUp) {
				content.CaptureDuration += check.AsOr(multiply_by_ten, float64(10), 1)
			}
		case CAPTURE_COUNT_DEC:
			if content.gamepad.IsGamepadButtonDown(i, gpads.RL_RightFaceDown) {
				content.CaptureDuration -= check.AsOr(multiply_by_ten, float64(10), 1)
				if content.CaptureDuration < 1 {
					content.CaptureDuration = 1
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

		case CAPTURE_MP4:
			down = content.gamepad.IsGamepadButtonDown(i, gpads.RL_MiddleLeft)
			if down && content.capturing {
				gs.EndCapture()
			} else if down {
				gs.BeginCapture("mp4")
			}
		case PAUSE_PLAY:
			if content.gamepad.IsGamepadButtonDown(i, gpads.RL_RightFaceLeft) {
				content.paused = !content.paused
				if !content.paused {
					gs.Refresh(content.currentTime)
				}
			}
		case RESIZE_UP:
			if content.gamepad.IsGamepadButtonDown(i, gpads.RL_RightTrigger1) {
				gs.resize(gs.Content.screenstate + 1)
			}
		case RESIZE_DOWN:
			if content.gamepad.IsGamepadButtonDown(i, gpads.RL_RightTrigger2) {
				gs.resize(gs.Content.screenstate - 1)
			}
		}
	}
}
