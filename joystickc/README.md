# joystickc

My gamepad didn't work with raylib, so I naively built a joystick controller with c and go support. Linux only.

## [joystickc.go](joystickc.go)

Implements [jstick](../jstick/README.md) interface.

## [joystick.c](joystick.c)/[joystick.h](joystick.h)

Joystick/Gamepad interface for linux using old style joystick.h

`bool IsJoystickAvailable(int Joystick)`

Check if a Joystick number is available.

`const char *GetJoystickName(int Joystick)`

Get Joystick internal name id.

`bool IsJoystickButtonPressed(int Joystick, int button)`

Check if a Joystick button has been pressed once.

`bool IsJoystickButtonDown(int Joystick, int button)`

Check if a Joystick button is being down.

`bool IsJoystickButtonReleased(int Joystick, int button)`

Check if a Joystick button has been released once.

`bool IsJoystickButtonUp(int Joystick, int button)`

Check if a Joystick button is NOT down.

`int GetJoystickButtonPressed(void)`

Get the last Joystick button pressed.

`int GetJoystickAxisCount(int Joystick)`

Get Joystick axis count for a Joystick

`float GetJoystickAxisMovement(int Joystick, int axis)`

Get axis movement value for a Joystick axis

`int SetJoystickMappings(const char *mappings)`

Set internal Joystick mappings (SDL_GameControllerDB).

**not yet implemented.**

`void BeginJoystick(void)`
 
Updates the joystick state. 
This must be called once per frame at the *BeginDrawing* step.

