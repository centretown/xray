package joystickc

/*
#include "joystick.h"
*/
import "C"

import "xray/jstick"

var _ jstick.Jstick = NewJoyStickC()

type JoystickC struct {
}

func NewJoyStickC() *JoystickC {
	js := &JoystickC{}
	return js
}

func (js *JoystickC) BeginJoystick() {
	C.BeginJoystick()
}

func (js *JoystickC) IsJoystickAvailable(Joystick int) bool {
	return bool(C.IsJoystickAvailable(C.int(Joystick)))
}

func (js *JoystickC) GetJoystickName(Joystick int) string {
	return C.GoString(C.GetJoystickName(C.int(Joystick)))
}

func (js *JoystickC) IsJoystickButtonPressed(Joystick int, button int) bool {
	return bool(C.IsJoystickButtonPressed(C.int(Joystick), C.int(button)))
}

func (js *JoystickC) IsJoystickButtonDown(Joystick int, button int) bool {
	return bool(C.IsJoystickButtonDown(C.int(Joystick), C.int(button)))
}

func (js *JoystickC) IsJoystickButtonReleased(Joystick int, button int) bool {
	return bool(C.IsJoystickButtonReleased(C.int(Joystick), C.int(button)))
}

func (js *JoystickC) IsJoystickButtonUp(Joystick int, button int) bool {
	return bool(C.IsJoystickButtonUp(C.int(Joystick), C.int(button)))
}

func (js *JoystickC) GetJoystickButtonPressed() int {
	return int(C.GetJoystickButtonPressed())
}

func (js *JoystickC) GetJoystickAxisCount(Joystick int) int {
	return int(C.GetJoystickAxisCount(C.int(Joystick)))
}

func (js *JoystickC) GetJoystickButtonCount(Joystick int) int {
	return int(C.GetJoystickButtonCount(C.int(Joystick)))
}

func (js *JoystickC) GetJoystickAxisMovement(Joystick int, axis int) float32 {
	return float32(C.GetJoystickAxisMovement(C.int(Joystick), C.int(axis)))
}

func (js *JoystickC) GetJoystickAxisValue(Joystick int, axis int) int32 {
	return int32(C.GetJoystickAxisValue(C.int(Joystick), C.int(axis)))
}

func (js *JoystickC) SetJoystickMappings(mappings string) int {
	return 0
}

func (js *JoystickC) DumpJoystick() {
	C.DumpJoystick()
}

func (js *JoystickC) GetButtonName(Joystick int, button int) string {
	return C.GoString(C.GetButtonName(C.int(Joystick), C.int(button)))
}
