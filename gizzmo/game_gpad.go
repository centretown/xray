package gizzmo

import (
	"github.com/centretown/gpads/gpads"
	"github.com/centretown/xray/notes"
	"github.com/centretown/xray/numbers"
)

func (gs *Game) CheckPad(padNum int32) notes.Command {
	content := &gs.Content
	content.gamepad.BeginPad()

	var (
		more_down = content.gamepad.IsGamepadButtonDown(padNum, gpads.RL_RightTrigger1)
		aux_down  = content.gamepad.IsGamepadButtonDown(padNum, gpads.RL_LeftTrigger1)
	)

	for command := range notes.COMMAND_COUNT {
		switch command {
		case notes.OPTIONS:
			if content.gamepad.IsGamepadButtonDown(padNum, gpads.RL_MiddleRight) {
				return command
			}
		case notes.MORE:
		case notes.NEXT:
			if content.gamepad.IsGamepadButtonDown(padNum, gpads.RL_LeftFaceDown) {
				return command
			}
		case notes.PREVIOUS:
			if content.gamepad.IsGamepadButtonDown(padNum, gpads.RL_LeftFaceUp) {
				return command
			}
		case notes.INCREMENT:
			if content.gamepad.IsGamepadButtonDown(padNum, gpads.RL_LeftFaceRight) {
				inc := numbers.AsOr(more_down, notes.INCREMENT_MORE, notes.INCREMENT)
				return inc
			}
		case notes.DECREMENT:
			if content.gamepad.IsGamepadButtonDown(padNum, gpads.RL_LeftFaceLeft) {
				dec := numbers.AsOr(more_down, notes.DECREMENT_MORE, notes.DECREMENT)
				return dec
			}
		case notes.SHARE:
			if content.gamepad.IsGamepadButtonDown(padNum, gpads.RL_MiddleLeft) {
				return command
			}
		case notes.PAUSE_PLAY:
			if !aux_down &&
				content.gamepad.IsGamepadButtonDown(padNum, gpads.RL_RightFaceLeft) {
				return command
			}
		case notes.ACTION:
			if content.gamepad.IsGamepadButtonDown(padNum, gpads.RL_RightFaceDown) {
				return command
			}
		case notes.BACK:
			if content.gamepad.IsGamepadButtonDown(padNum, gpads.RL_RightFaceUp) {
				return command
			}
		case notes.CANCEL:
			if content.gamepad.IsGamepadButtonDown(padNum, gpads.RL_RightFaceRight) {
				return command
			}
		case notes.OUT:
			if aux_down &&
				content.gamepad.IsGamepadButtonDown(padNum, gpads.RL_RightFaceLeft) {
				return command
			}
		}
	}
	return notes.NONE
}
