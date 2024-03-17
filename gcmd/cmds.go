package gcmd

import (
	"fmt"
	"xray/b2"
	"xray/gpads"
)

func LastButtonPressed(cmd *GCmd) {
	button := js.GetPadButtonPressed()
	fmt.Printf("[%2d]\r", button)
}

func IsButtonUp(cmd *GCmd) {
	up := js.IsPadButtonUp(cmd.Pad, cmd.Button)
	fmt.Printf("[%d:%d]\r", cmd.Button, b2.ToInt(up))
}

func IsButtonDown(cmd *GCmd) {
	down := js.IsPadButtonDown(cmd.Pad, cmd.Button)
	fmt.Printf("[%d:%d]\r", cmd.Button, b2.ToInt(down))
}

func IsButtonReleased(cmd *GCmd) {
	released := js.IsPadButtonReleased(cmd.Pad, cmd.Button)
	fmt.Printf("[%d:%d]\r", cmd.Button, b2.ToInt(released))
}

func IsButtonPressed(cmd *GCmd) {
	pressed := js.IsPadButtonPressed(cmd.Pad, cmd.Button)
	fmt.Printf("[%d:%d]\r", cmd.Button, b2.ToInt(pressed))
}

func GetAxisValues(cmd *GCmd) {
	count := js.GetPadAxisCount(cmd.Pad)
	fmt.Print("axes:  ")
	for i := range count {
		value := js.GetPadAxisValue(cmd.Pad, i)
		fmt.Printf("[%d:%6d] ", i, value)
	}
	fmt.Print("\r")
}

func GetAxisMovement(cmd *GCmd) {
	fmt.Print("axes:  ")
	count := js.GetPadAxisCount(cmd.Pad)
	for i := range count {
		move := js.GetPadAxisMovement(cmd.Pad, i)
		fmt.Printf("[%d:%6.0f] ", i, move)
	}
	fmt.Print("\r")
}

func DumpPad(cmd *GCmd) {
	js.DumpPad()
}

const MAX_BUTTONS = 64

func TestKeys(cmd *GCmd) {
	count := gpads.RL_RightThumb - gpads.RL_Unknown + 1
	for i := range count {
		down := js.IsPadButtonDown(cmd.Pad, i)
		fmt.Printf("[%x:%2d]", i, b2.ToInt(down))
	}
	fmt.Print("\r")
}

func TestAxes(cmd *GCmd) {
	count := js.GetPadAxisCount(cmd.Pad)
	for i := range count {
		val := js.GetPadAxisValue(cmd.Pad, i)
		mov := js.GetPadAxisMovement(cmd.Pad, i)
		fmt.Printf("[%x:%03x:%3.0f]", i, val, mov)
	}
	fmt.Print("  \r")
}
