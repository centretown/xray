#ifndef __JOY_H__
#define __JOY_H__
#include <stdbool.h>

bool IsJoystickAvailable(int Joystick);    // Check if a Joystick is available
const char *GetJoystickName(int Joystick); // Get Joystick internal name id
bool IsJoystickButtonPressed(int Joystick, int button);  // Check if a Joystick button has been pressed once
bool IsJoystickButtonDown(int Joystick, int button);     // Check if a Joystick button is being pressed
bool IsJoystickButtonReleased(int Joystick, int button); // Check if a Joystick button has been released once
bool IsJoystickButtonUp(int Joystick, int button);       // Check if a Joystick button is NOT being pressed
int GetJoystickButtonPressed(void);                      // Get the last Joystick button pressed
int GetJoystickAxisCount(int Joystick);                  // Get Joystick axis count for a Joystick
float GetJoystickAxisMovement(int Joystick, int axis);   // Get axis movement value for a Joystick axis
int SetJoystickMappings(const char *mappings);           // Set internal Joystick mappings (SDL_GameControllerDB)
void BeginJoystick(void);
void Dump(void);
#endif // __JOY_H__