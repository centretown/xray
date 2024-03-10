# JStick

My intent was to make it usable and familiar with Raylib Gamepad interface.

[jstick.go](jstick.go)
go interface implemented by GameStick and JoyStickC

`BeginJoystick()`

Updates the Joystick/Gamepad state

`IsJoystickAvailable(Joystick int) bool`

Check if a Joystick number is available.

`GetJoystickName(Joystick int) string`

Get Joystick internal name id.

`IsJoystickButtonPressed(Joystick int, button int) bool`

Check if a Joystick button has been pressed once.

`IsJoystickButtonDown(Joystick int, button int) bool`

Check if a Joystick button is being down.

`IsJoystickButtonReleased(Joystick int, button int) bool`

Check if a Joystick button has been released once.

`IsJoystickButtonUp(Joystick int, button int) bool`

Check if a Joystick button is NOT down.

`GetJoystickButtonPressed() int`

Get the last Joystick button pressed.

`GetJoystickAxisCount(Joystick int) int`

Get Joystick axis count for a Joystick

`GetJoystickButtonCount(Joystick int) int`

Get Joystick button count for a Joystick

`GetJoystickAxisMovement(Joystick int, axis int) float32`

Get axis movement change value for a Joystick axis

`GetJoystickAxisValue(Joystick int, axis int) int32`

Get axis movement value for a Joystick axis

`SetJoystickMappings(mappings string) int`

**not yet implemented.**

`DumpJoystick()`

