package gizzmo

import (
	"github.com/centretown/gpads/gpads"
	"github.com/centretown/xray/notes"
	"github.com/centretown/xray/numbers"
)

func (gs *Game) ProcessInput() {
	content := &gs.Content
	content.gamepad.BeginPad()

	if content.currentTime > content.nextInput {
		content.nextInput = content.currentTime + content.InputInterval
		for i := range content.gamepad.GetPadCount() {
			gs.CheckPad(i)
		}
	}

	if content.commandState {
		return
	}

	for _, child := range gs.Children() {
		t, ok := child.(Inputer)
		if ok {
			t.Input()
		}
	}
}

func (gs *Game) CheckPad(padNum int32) {
	content := &gs.Content
	var more_down bool

	if content.gamepad.IsGamepadButtonDown(padNum, gpads.RL_MiddleRight) {
		content.commandState = !content.commandState
	}

	if !content.commandState {
		return
	}

	for command := range notes.COMMANDS {
		switch command {
		case notes.MORE:
			more_down = content.gamepad.IsGamepadButtonDown(padNum, gpads.RL_RightTrigger1)

		case notes.NEXT_NOTE:
			if content.gamepad.IsGamepadButtonDown(padNum, gpads.RL_LeftFaceDown) {
				gs.UpdateState(command)
			}
		case notes.PREV_NOTE:
			if content.gamepad.IsGamepadButtonDown(padNum, gpads.RL_LeftFaceUp) {
				gs.UpdateState(command)
			}
		case notes.INC:
			if content.gamepad.IsGamepadButtonDown(padNum, gpads.RL_LeftFaceRight) {
				inc := numbers.AsOr(more_down, notes.INC_MORE, notes.INC)
				gs.UpdateState(inc)
			}
		case notes.DEC:
			if content.gamepad.IsGamepadButtonDown(padNum, gpads.RL_LeftFaceLeft) {
				dec := numbers.AsOr(more_down, notes.DEC_MORE, notes.DEC)
				gs.UpdateState(dec)
			}
		case notes.CAPTURE:
			if content.gamepad.IsGamepadButtonDown(padNum, gpads.RL_MiddleLeft) {
				gs.UpdateState(command)
			}
		case notes.PAUSE_PLAY:
			if content.gamepad.IsGamepadButtonDown(padNum, gpads.RL_RightFaceLeft) {
				gs.UpdateState(command)
			}
		}
	}
}
