package gpads

import "github.com/holoplot/go-evdev"

// button mappings

const BTN_GAMEPAD = evdev.BTN_GAMEPAD
const BTN_JOYSTICK = evdev.BTN_JOYSTICK

type PadType int

const (
	PAD_PS3 PadType = iota
	PAD_XBOX
	PAD_JOYSTICK
)

var PadTypes = []string{
	"PS3",
	"XBOX",
	"JOYSTICK",
}

func (t PadType) String() string {
	if t > PAD_JOYSTICK {
		return UNDEFINED
	}
	return PadTypes[t]
}

// mirror raylib constants

const (
	RL_Unknown int = iota
	RL_LeftFaceUp
	RL_LeftFaceRight
	RL_LeftFaceDown
	RL_LeftFaceLeft
	RL_RightFaceUp
	RL_RightFaceRight
	RL_RightFaceDown
	RL_RightFaceLeft
	RL_LeftTrigger1
	RL_LeftTrigger2
	RL_RightTrigger1
	RL_RightTrigger2
	RL_MiddleLeft
	RL_Middle
	RL_MiddleRight
	RL_LeftThumb
	RL_RightThumb
	RL_BUTTON_COUNT
)

var XBoxButtons = []evdev.EvCode{
	0,                    // Unknown button, just for error checking
	evdev.BTN_DPAD_UP,    // Gamepad left DPAD up button
	evdev.BTN_DPAD_RIGHT, // Gamepad left DPAD right button
	evdev.BTN_DPAD_DOWN,  // Gamepad left DPAD down button
	evdev.BTN_DPAD_LEFT,  // Gamepad left DPAD left button
	evdev.BTN_Y,          // Gamepad right button up (i.e. PS3: Triangle, Xbox: Y)
	evdev.BTN_B,          // Gamepad right button left (i.e. PS3: Circle, Xbox: B)
	evdev.BTN_A,          // Gamepad right button down (i.e. PS3: Cross, Xbox: A)
	evdev.BTN_X,          // Gamepad right button right (i.e. PS3: Square, Xbox: X)
	evdev.BTN_TL,         // Gamepad top/back trigger left (first), it could be a trailing button
	evdev.BTN_TL2,        // Gamepad top/back trigger left (second), it could be a trailing button
	evdev.BTN_TR,         // Gamepad top/back trigger right (one), it could be a trailing button
	evdev.BTN_TR2,        // Gamepad top/back trigger right (second), it could be a trailing button
	evdev.BTN_SELECT,     // Gamepad center buttons, left one (i.e. PS3: Select)
	evdev.BTN_MODE,       // Gamepad center buttons, middle one (i.e. PS3: PS, Xbox: XBOX)
	evdev.BTN_START,      // Gamepad center buttons, right one (i.e. PS3: Start)
	evdev.BTN_THUMBL,     // Gamepad joystick pressed button left
	evdev.BTN_THUMBR,     // Gamepad joystick pressed button right
}

var PS3Buttons = []evdev.EvCode{
	0,                    // Unknown button, just for error checking
	evdev.BTN_DPAD_UP,    // Gamepad left DPAD up button
	evdev.BTN_DPAD_RIGHT, // Gamepad left DPAD right button
	evdev.BTN_DPAD_DOWN,  // Gamepad left DPAD down button
	evdev.BTN_DPAD_LEFT,  // Gamepad left DPAD left button
	evdev.BTN_NORTH,      // Gamepad right button up (i.e. PS3: Triangle, Xbox: Y)
	evdev.BTN_EAST,       // Gamepad right button left (i.e. PS3: Circle, Xbox: B)
	evdev.BTN_SOUTH,      // Gamepad right button down (i.e. PS3: Cross, Xbox: A)
	evdev.BTN_WEST,       // Gamepad right button right (i.e. PS3: Square, Xbox: X)
	evdev.BTN_TL,         // Gamepad top/back trigger left (first), it could be a trailing button
	evdev.BTN_TL2,        // Gamepad top/back trigger left (second), it could be a trailing button
	evdev.BTN_TR,         // Gamepad top/back trigger right (one), it could be a trailing button
	evdev.BTN_TR2,        // Gamepad top/back trigger right (second), it could be a trailing button
	evdev.BTN_SELECT,     // Gamepad center buttons, left one (i.e. PS3: Select)
	evdev.BTN_MODE,       // Gamepad center buttons, middle one (i.e. PS3: PS, Xbox: XBOX)
	evdev.BTN_START,      // Gamepad center buttons, right one (i.e. PS3: Start)
	evdev.BTN_THUMBL,     // Gamepad joystick pressed button left
	evdev.BTN_THUMBR,     // Gamepad joystick pressed button right
}

var JoyButtons = []evdev.EvCode{
	0,                    // Unknown button, just for error checking
	evdev.BTN_DPAD_UP,    // Gamepad left DPAD up button
	evdev.BTN_DPAD_RIGHT, // Gamepad left DPAD right button
	evdev.BTN_DPAD_DOWN,  // Gamepad left DPAD down button
	evdev.BTN_DPAD_LEFT,  // Gamepad left DPAD left button
	evdev.BTN_TOP,        // Gamepad right button up (i.e. PS3: Triangle, Xbox: Y)
	evdev.BTN_THUMB2,     // Gamepad right button left (i.e. PS3: Circle, Xbox: B)
	evdev.BTN_THUMB,      // Gamepad right button down (i.e. PS3: Cross, Xbox: A)
	evdev.BTN_JOYSTICK,   // Gamepad right button right (i.e. PS3: Square, Xbox: X)
	evdev.BTN_TOP2,       // Gamepad top/back trigger left (first), it could be a trailing button
	evdev.BTN_BASE,       // Gamepad top/back trigger left (second), it could be a trailing button
	evdev.BTN_PINKIE,     // Gamepad top/back trigger right (one), it could be a trailing button
	evdev.BTN_BASE2,      // Gamepad top/back trigger right (second), it could be a trailing button
	evdev.BTN_BASE3,      // Gamepad center buttons, left one (i.e. PS3: Select)
	evdev.BTN_DEAD,       // Gamepad center buttons, middle one (i.e. PS3: PS, Xbox: XBOX)
	evdev.BTN_BASE4,      // Gamepad center buttons, right one (i.e. PS3: Start)
	evdev.BTN_BASE5,      // Gamepad joystick pressed button left
	evdev.BTN_BASE6,      // Gamepad joystick pressed button right
}

// mirror raylib constants

const (
	RL_AxisLeftX int = iota
	RL_AxisLeftY
	RL_AxisRightX
	RL_AxisRightY
	RL_LeftTrigger
	RL_RightTrigger
	RL_HatX
	RL_HatY
	RL_AXIS_COUNT
)

var GameAxes = []evdev.EvCode{
	evdev.ABS_X,
	evdev.ABS_Y,
	evdev.ABS_RX,
	evdev.ABS_RY,
	evdev.ABS_Z,
	evdev.ABS_RZ,
	evdev.ABS_HAT0X,
	evdev.ABS_HAT0Y,
}

var JoyAxes = []evdev.EvCode{
	evdev.ABS_X,
	evdev.ABS_Y,
	evdev.ABS_Z,
	evdev.ABS_RZ,
	evdev.ABS_RX, //absent
	evdev.ABS_RY, //absent
	evdev.ABS_HAT0X,
	evdev.ABS_HAT0Y,
}

type AxisInfoMap map[evdev.EvCode]evdev.AbsInfo
type AxisStateMap map[evdev.EvCode]int32
type ButtonStateMap map[evdev.EvCode]bool
