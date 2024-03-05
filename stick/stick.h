#ifndef __STICK_H__
#define __STICK_H__
#include <stdbool.h>

bool IsJoystickAvailable(int Joystick);                  

const char *GetJoystickName(int Joystick);               

bool IsJoystickButtonPressed(int Joystick, int button);  

bool IsJoystickButtonDown(int Joystick, int button);     

bool IsJoystickButtonReleased(int Joystick, int button); 

bool IsJoystickButtonUp(int Joystick, int button);       

int GetJoystickButtonPressed(void);                      

int GetJoystickAxisCount(int Joystick);                  

float GetJoystickAxisMovement(int Joystick, int axis);   

int SetJoystickMappings(const char *mappings);           

// Update the state each cycle
void BeginJoystick(void);                                

// Get the button mapping name?
const char *GetButtonName(int Joystick, int button);

#define STICK_EXTRA
// Dump state for analysis
void Dump(void);

#endif // __STICK_H__