package jstick

type Jstick interface {
	BeginJoystick()
	IsJoystickAvailable(Joystick int) bool
	GetJoystickName(Joystick int) string
	IsJoystickButtonPressed(Joystick int, button int) bool
	IsJoystickButtonDown(Joystick int, button int) bool
	IsJoystickButtonReleased(Joystick int, button int) bool
	IsJoystickButtonUp(Joystick int, button int) bool
	GetJoystickButtonPressed() int
	GetJoystickAxisCount(Joystick int) int
	GetJoystickButtonCount(Joystick int) int
	GetJoystickAxisMovement(Joystick int, axis int) float32
	GetJoystickAxisValue(Joystick int, axis int) int32
	SetJoystickMappings(mappings string) int
	DumpJoystick()
	GetButtonName(Joystick int, button int) string
}
