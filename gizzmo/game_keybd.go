package gizzmo

import (
	rl "github.com/centretown/raylib-go/raylib"
	"github.com/centretown/xray/notes"
	"github.com/centretown/xray/numbers"
)

func (gs *Game) CheckKeys() notes.COMMAND {
	var more_down bool

	for command := range notes.COMMANDS {
		switch command {
		case notes.OPTIONS:
			if rl.IsKeyDown(rl.KeyF1) {
				return command
			}

		case notes.MORE:
			more_down = rl.IsKeyDown(rl.KeyLeftControl) ||
				rl.IsKeyDown(rl.KeyRightControl)

		case notes.NEXT:
			if rl.IsKeyDown(rl.KeyDown) {
				return command
			}
		case notes.PREVIOUS:
			if rl.IsKeyDown(rl.KeyUp) {
				return command
			}
		case notes.INCREMENT:
			if rl.IsKeyDown(rl.KeyRight) {
				inc := numbers.AsOr(more_down, notes.INCREMENT_MORE, notes.INCREMENT)
				return inc
			}
		case notes.DECREMENT:
			if rl.IsKeyDown(rl.KeyLeft) {
				dec := numbers.AsOr(more_down, notes.DECREMENT_MORE, notes.DECREMENT)
				return dec
			}
		case notes.SHARE:
			if rl.IsKeyDown(rl.KeyF5) {
				return command
			}
		case notes.PAUSE_PLAY:
			if rl.IsKeyDown(rl.KeyF7) {
				return command
			}
		case notes.ACTION:
			if rl.IsKeyDown(rl.KeyInsert) {
				return command
			}
		case notes.BACK:
			if rl.IsKeyDown(rl.KeyBackspace) {
				return command
			}
		case notes.CANCEL:
			if rl.IsKeyDown(rl.KeyDelete) {
				return command
			}
		case notes.OUT:
			if rl.IsKeyDown(rl.KeyHome) {
				return command
			}

		}
	}
	return notes.NONE

}
