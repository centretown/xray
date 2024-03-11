# joystickc

My gamepad didn't work with raylib, so I naively built a joystick controller with c and go support. Linux only.

## [joystickc.go](joystickc.go)

Implements [pad](../pad/README.md) interface.

## [joystick.c](joystick.c)/[joystick.h](joystick.h)

Pad/Gamepad interface for linux using old style joystick.h

`bool IsPadAvailable(int Pad)`

Check if a Pad number is available.

`const char *GetPadName(int Pad)`

Get Pad internal name id.

`bool IsPadButtonPressed(int Pad, int button)`

Check if a Pad button has been pressed once.

`bool IsPadButtonDown(int Pad, int button)`

Check if a Pad button is being down.

`bool IsPadButtonReleased(int Pad, int button)`

Check if a Pad button has been released once.

`bool IsPadButtonUp(int Pad, int button)`

Check if a Pad button is NOT down.

`int GetPadButtonPressed(void)`

Get the last Pad button pressed.

`int GetPadAxisCount(int Pad)`

Get Pad axis count for a Pad

`float GetPadAxisMovement(int Pad, int axis)`

Get axis movement value for a Pad axis

`int SetPadMappings(const char *mappings)`

Set internal Pad mappings (SDL_GameControllerDB).

**not yet implemented.**

`void BeginPad(void)`
 
Updates the joystick state. 
This must be called once per frame at the *BeginDrawing* step.

