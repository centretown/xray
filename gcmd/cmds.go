package gcmd

import (
	"fmt"
	"xray/b2i"
	"xray/gpads"
)

func LastButtonPressed(cmd *GCmd) {
	button := js.GetPadButtonPressed()
	up := js.IsPadButtonDown(cmd.Pad, button)
	fmt.Printf("[%4d:%4d]\r", button, b2i.Bool2int(up))
}

func IsButtonUp(cmd *GCmd) {
	up := js.IsPadButtonUp(cmd.Pad, cmd.Button)
	fmt.Printf("[%d:%d]\r", cmd.Button, b2i.Bool2int(up))
}

func IsButtonDown(cmd *GCmd) {
	down := js.IsPadButtonDown(cmd.Pad, cmd.Button)
	fmt.Printf("[%d:%d]\r", cmd.Button, b2i.Bool2int(down))
}

func IsButtonReleased(cmd *GCmd) {
	released := js.IsPadButtonReleased(cmd.Pad, cmd.Button)
	fmt.Printf("[%d:%d]\r", cmd.Button, b2i.Bool2int(released))
}

func IsButtonPressed(cmd *GCmd) {
	pressed := js.IsPadButtonPressed(cmd.Pad, cmd.Button)
	fmt.Printf("[%d:%d]\r", cmd.Button, b2i.Bool2int(pressed))
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
	count := gpads.BTN_THUMBR - gpads.BTN_RESERVED + 1
	for i := range count {
		down := js.IsPadButtonDown(cmd.Pad, i)
		fmt.Printf("[%x:%2d]", i, b2i.Bool2int(down))
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
