package gizzmo

import (
	"time"

	rl "github.com/centretown/raylib-go/raylib"
	"github.com/centretown/xray/notes"
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
		previousCommand notes.COMMAND
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

			gs.updateState(currentCommand)

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
