package joystickc

/*
#include "joystick.h"
*/
import "C"

import "xray/pad"

var _ pad.Pad = NewJoyStickC()

type JoystickC struct {
}

func NewJoyStickC() *JoystickC {
	js := &JoystickC{}
	return js
}

func (js *JoystickC) BeginPad() {
	C.BeginJoystick()
}

func (js *JoystickC) IsPadAvailable(Joystick int) bool {
	return bool(C.IsJoystickAvailable(C.int(Joystick)))
}

func (js *JoystickC) GetPadName(Joystick int) string {
	return C.GoString(C.GetJoystickName(C.int(Joystick)))
}

func (js *JoystickC) IsPadButtonPressed(Joystick int, button int) bool {
	return bool(C.IsJoystickButtonPressed(C.int(Joystick), C.int(button)))
}

func (js *JoystickC) IsPadButtonDown(Joystick int, button int) bool {
	return bool(C.IsJoystickButtonDown(C.int(Joystick), C.int(button)))
}

func (js *JoystickC) IsPadButtonReleased(Joystick int, button int) bool {
	return bool(C.IsJoystickButtonReleased(C.int(Joystick), C.int(button)))
}

func (js *JoystickC) IsPadButtonUp(Joystick int, button int) bool {
	return bool(C.IsJoystickButtonUp(C.int(Joystick), C.int(button)))
}

func (js *JoystickC) GetPadButtonPressed() int {
	return int(C.GetJoystickButtonPressed())
}

func (js *JoystickC) GetPadAxisCount(Joystick int) int {
	return int(C.GetJoystickAxisCount(C.int(Joystick)))
}

func (js *JoystickC) GetPadButtonCount(Joystick int) int {
	return int(C.GetJoystickButtonCount(C.int(Joystick)))
}

func (js *JoystickC) GetPadAxisMovement(Joystick int, axis int) float32 {
	return float32(C.GetJoystickAxisMovement(C.int(Joystick), C.int(axis)))
}

func (js *JoystickC) GetPadAxisValue(Joystick int, axis int) int32 {
	return int32(C.GetJoystickAxisValue(C.int(Joystick), C.int(axis)))
}

func (js *JoystickC) SetPadMappings(mappings string) int {
	return 0
}

func (js *JoystickC) DumpPad() {
	C.DumpJoystick()
}

func (js *JoystickC) GetButtonName(Joystick int, button int) string {
	return C.GoString(C.GetButtonName(C.int(Joystick), C.int(button)))
}
