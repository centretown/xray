package gpads

import (
	"fmt"
	"xray/pad"

	"github.com/holoplot/go-evdev"
)

var _ pad.Pad = NewGPads()

const PAD_MAX = 4

type GPads struct {
	Pads       [PAD_MAX]*GPad
	Err        error
	padCount   int
	intialized bool
}

func NewGPads() *GPads {
	js := &GPads{}
	return js
}

func (js *GPads) intialize() {
	var (
		devicePaths []evdev.InputPath
		err         error
	)

	defer func() { js.intialized = true }()

	devicePaths, err = evdev.ListDevicePaths()
	if err != nil {
		fmt.Println("ListDevicePaths", err)
		devicePaths = make([]evdev.InputPath, 0)
	}

	js.padCount = 0
	for _, p := range devicePaths {
		stg, err := OpenGPad(p.Path)
		if err != nil {
			fmt.Printf("Open %s [%s] %v\n", p.Name, p.Path, err)
			continue
		}
		js.Pads[js.padCount] = stg
		js.padCount++
	}
}

func (js *GPads) BeginPad() {
	if !js.intialized {
		js.intialize()
	}

	for i := range js.padCount {
		js.Pads[i].ReadState()
	}
}

func (js *GPads) IsPadAvailable(pad int) bool {
	return js.padCount > pad
}

const UNDEFINED = "UNDEFINED"

func (js *GPads) GetStickCount() int {
	return js.padCount
}

func (js *GPads) GetPadName(pad int) string {
	if js.padCount <= pad {
		return UNDEFINED
	}
	return js.Pads[pad].Name
}

func (js *GPads) DumpState(pad int) {
	if js.padCount <= pad {
		return
	}
	js.Pads[pad].DumpState()
}

func (js *GPads) GetButtonName(pad int, button int) string {
	if js.padCount <= pad || button >= ButtonCount {
		return UNDEFINED
	}
	return evdev.KEYNames[GameButtons[button]]
}

func (js *GPads) IsPadButtonPressed(pad int, button int) bool {
	if js.padCount <= pad || ButtonCount <= button {
		return false
	}
	return js.Pads[pad].ButtonPressed(button)
}

func (js *GPads) IsPadButtonDown(pad int, button int) bool {
	if js.padCount <= pad || button >= ButtonCount {
		return false
	}
	return js.Pads[pad].ButtonDown(button)
}

func (js *GPads) IsPadButtonReleased(pad int, button int) bool {
	if js.padCount <= pad || button >= ButtonCount {
		return false
	}
	return js.Pads[pad].ButtonReleased(button)
}

func (js *GPads) IsPadButtonUp(pad int, button int) bool {
	if js.padCount <= pad || button >= ButtonCount {
		return false
	}
	return !js.Pads[pad].ButtonDown(button)
}

func (js *GPads) GetPadButtonPressed() int {
	return 0
}

func (js *GPads) GetPadAxisCount(pad int) int {
	if js.padCount <= pad {
		return 0
	}
	return len(js.Pads[pad].curAxisState)
}

func (js *GPads) GetPadButtonCount(pad int) int {
	if js.padCount <= pad {
		return 0
	}
	return len(js.Pads[pad].curButtonState)
}

func (js *GPads) GetPadAxisMovement(pad int, axis int) float32 {
	if js.padCount <= pad {
		return 0
	}
	return js.Pads[pad].AxisMove(axis)
}

func (js *GPads) GetPadAxisValue(pad int, axis int) int32 {
	if js.padCount <= pad {
		return 0
	}
	return js.Pads[pad].AxisValue(axis)
}

func (js *GPads) SetPadMappings(mappings string) int {
	return 0
}

func (js *GPads) DumpPad() {
	for i := 0; i < js.padCount; i++ {
		js.Pads[i].Dump()
		fmt.Println()
	}
}

func (js *GPads) Close() {
	for i := 0; i < js.padCount; i++ {
		js.Pads[i].Close()
	}
}
