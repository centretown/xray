package tools

/*
#include "joy.h"
*/
import "C"

func IsJoystickAvailable(joystick int) bool {
	return bool(C.IsJoystickAvailable(C.int(joystick)))
}

func GetJoystickName(joystick int) string {
	return C.GoString(C.GetJoystickName(C.int(joystick)))
}

func IsJoystickButtonPressed(joystick, button int) bool {
	return bool(C.IsJoystickButtonPressed(C.int(joystick), C.int(button)))
}

func BeginJoystick() {
	C.BeginJoystick()
}

func GetJoystickButtonPressed() int {
	return int(C.GetJoystickButtonPressed())
}
