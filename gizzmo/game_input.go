package gizzmo

import (
	"time"

	"github.com/centretown/gpads/gpads"
	rl "github.com/centretown/raylib-go/raylib"
	"github.com/centretown/xray/notes"
	"github.com/centretown/xray/numbers"
)

func (gs *Game) ProcessInput(repeatinterval float64,
	repeatCh <-chan float64, stop <-chan int) {

	var (
		content = &gs.Content

		thisTime     float64 = rl.GetTime()
		nextTime     float64 = thisTime + repeatinterval
		baseInterval         = repeatinterval
		interval             = baseInterval

		command, currentCommand,
		previousCommand int
		repeatCount int
	)

	for {
		time.Sleep(time.Millisecond)
		select {
		case <-stop:
			return
		case baseInterval = <-repeatCh:
		default:
			thisTime = content.currentTime

			command = gs.CheckKeys()
			if command == notes.NONE {
				for i := range content.gamepad.GetPadCount() {
					command = gs.CheckPad(i)
					if command != notes.NONE {
						break
					}
				}
			}

			if command == notes.NONE {
				if thisTime > nextTime {
					repeatCount = 0
					interval = baseInterval
					previousCommand = notes.NONE
				}
				continue
			}

			currentCommand = command
			if currentCommand == previousCommand && thisTime < nextTime {
				continue
			}

			gs.UpdateState(currentCommand)

			if currentCommand != previousCommand {
				repeatCount = 0
				previousCommand = currentCommand
				nextTime = thisTime + interval
				interval = baseInterval
				continue
			}
			if repeatCount > 0 {
				interval = baseInterval / float64(repeatCount)
			}
			nextTime = thisTime + interval
			repeatCount++
		}
	}

}

func (gs *Game) CheckPad(padNum int32) int {
	content := &gs.Content
	var more_down bool
	content.gamepad.BeginPad()

	for command := range notes.COMMANDS {
		switch command {
		case notes.HELP:
			if content.gamepad.IsGamepadButtonDown(padNum, gpads.RL_MiddleRight) {
				return command
			}
		case notes.MORE:
			more_down = content.gamepad.IsGamepadButtonDown(padNum, gpads.RL_RightTrigger2)

		case notes.NEXT_NOTE:
			if content.gamepad.IsGamepadButtonDown(padNum, gpads.RL_LeftFaceDown) {
				return command
			}
		case notes.PREV_NOTE:
			if content.gamepad.IsGamepadButtonDown(padNum, gpads.RL_LeftFaceUp) {
				return command
			}
		case notes.INC:
			if content.gamepad.IsGamepadButtonDown(padNum, gpads.RL_LeftFaceRight) {
				inc := numbers.AsOr(more_down, notes.INC_MORE, notes.INC)
				return inc
			}
		case notes.DEC:
			if content.gamepad.IsGamepadButtonDown(padNum, gpads.RL_LeftFaceLeft) {
				dec := numbers.AsOr(more_down, notes.DEC_MORE, notes.DEC)
				return dec
			}
		case notes.CAPTURE:
			if content.gamepad.IsGamepadButtonDown(padNum, gpads.RL_MiddleLeft) {
				return command
			}
		case notes.PAUSE_PLAY:
			if content.gamepad.IsGamepadButtonDown(padNum, gpads.RL_RightFaceLeft) {
				return command
			}
		}
	}
	return notes.NONE
}

func (gs *Game) CheckKeys() int {
	var more_down bool

	for command := range notes.COMMANDS {
		switch command {
		case notes.HELP:
			if rl.IsKeyDown(rl.KeyF1) {
				return command
			}

		case notes.MORE:
			more_down = rl.IsKeyDown(rl.KeyLeftControl) ||
				rl.IsKeyDown(rl.KeyRightControl)

		case notes.NEXT_NOTE:
			if rl.IsKeyDown(rl.KeyDown) {
				return command
			}
		case notes.PREV_NOTE:
			if rl.IsKeyDown(rl.KeyUp) {
				return command
			}
		case notes.INC:
			if rl.IsKeyDown(rl.KeyRight) {
				inc := numbers.AsOr(more_down, notes.INC_MORE, notes.INC)
				return inc
			}
		case notes.DEC:
			if rl.IsKeyDown(rl.KeyLeft) {
				dec := numbers.AsOr(more_down, notes.DEC_MORE, notes.DEC)
				return dec
			}
		case notes.CAPTURE:
			if rl.IsKeyDown(rl.KeyF5) {
				return command
			}
		case notes.PAUSE_PLAY:
			if rl.IsKeyDown(rl.KeyF7) {
				return command
			}
		}
	}
	return notes.NONE

}
