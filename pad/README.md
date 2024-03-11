# Pad

My intent was to make it usable and familiar with Raylib Gamepad interface.

[jstick.go](jstick.go)
go interface implemented by GameStick and JoyStickC

`BeginPad()`

Updates the Gamepad state

`IsPadAvailable(Pad int) bool`

Check if a Pad number is available.

`GetPadName(Pad int) string`

Get Pad internal name id.

`IsPadButtonPressed(Pad int, button int) bool`

Check if a Pad button has been pressed once.

`IsPadButtonDown(Pad int, button int) bool`

Check if a Pad button is being down.

`IsPadButtonReleased(Pad int, button int) bool`

Check if a Pad button has been released once.

`IsPadButtonUp(Pad int, button int) bool`

Check if a Pad button is NOT down.

`GetPadButtonPressed() int`

Get the last Pad button pressed.

`GetPadAxisCount(Pad int) int`

Get Pad axis count for a Pad

`GetPadButtonCount(Pad int) int`

Get Pad button count for a Pad

`GetPadAxisMovement(Pad int, axis int) float32`

Get axis movement change value for a Pad axis

`GetPadAxisValue(Pad int, axis int) int32`

Get axis movement value for a Pad axis

`SetPadMappings(mappings string) int`

**not yet implemented.**

`DumpPad()`

