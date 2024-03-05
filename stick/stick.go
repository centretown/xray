package stick

/*
#include "stick.h"
*/
import "C"

func BeginJoystick() {
	C.BeginJoystick()
}

func IsJoystickAvailable(Joystick int) bool {
	return bool(C.IsJoystickAvailable(C.int(Joystick)))
}

func GetJoystickName(Joystick int) string {
	return C.GoString(C.GetJoystickName(C.int(Joystick)))
}

func IsJoystickButtonPressed(Joystick int, button int) bool {
	return bool(C.IsJoystickButtonPressed(C.int(Joystick), C.int(button)))
}

func IsJoystickButtonDown(Joystick int, button int) bool {
	return bool(C.IsJoystickButtonDown(C.int(Joystick), C.int(button)))
}

func IsJoystickButtonReleased(Joystick int, button int) bool {
	return bool(C.IsJoystickButtonReleased(C.int(Joystick), C.int(button)))
}

func IsJoystickButtonUp(Joystick int, button int) bool {
	return bool(C.IsJoystickButtonUp(C.int(Joystick), C.int(button)))
}

func GetJoystickButtonPressed() int {
	return int(C.GetJoystickButtonPressed())
}

func GetJoystickAxisCount(Joystick int) int {
	return int(C.GetJoystickAxisCount(C.int(Joystick)))
}

func GetJoystickAxisMovement(Joystick int, axis int) float32 {
	return float32(C.GetJoystickAxisMovement(C.int(Joystick), C.int(axis)))
}

func SetJoystickMappings(mappings string) int {
	return 0
}

func Dump() {}

func GetButtonName(Joystick int, button int) string {
	return C.GoString(C.GetButtonName(C.int(Joystick), C.int(button)))
}
