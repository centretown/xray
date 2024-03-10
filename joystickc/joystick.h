#ifndef __STICK_H__
#define __STICK_H__
#include <stdbool.h>
#include <sys/types.h>

bool IsJoystickAvailable(int Joystick);                  

const char *GetJoystickName(int Joystick);               

bool IsJoystickButtonPressed(int Joystick, int button);  

bool IsJoystickButtonDown(int Joystick, int button);     

bool IsJoystickButtonReleased(int Joystick, int button); 

bool IsJoystickButtonUp(int Joystick, int button);       

int GetJoystickButtonPressed(void);                      

int GetJoystickButtonCount(int Joystick);

int GetJoystickAxisCount(int Joystick);                  

int16_t GetJoystickAxisValue(int Joystick, int axis);   

float GetJoystickAxisMovement(int Joystick, int axis);   

int SetJoystickMappings(const char *mappings);           

// Update the state each cycle
void BeginJoystick(void);                                

// Get the button mapping name?
const char *GetButtonName(int Joystick, int button);

// #define STICK_EXTRA
// Dump state for analysis
void DumpJoystick(void);

#endif // __STICK_H__