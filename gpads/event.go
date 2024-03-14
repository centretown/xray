package gpads

import "github.com/holoplot/go-evdev"

// button mappings

var ButtonCount = len(PS3Buttons)

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

const (
	BTN_RESERVED int = iota
	BTN_DPAD_UP
	BTN_DPAD_RIGHT
	BTN_DPAD_DOWN
	BTN_DPAD_LEFT
	BTN_NORTH
	BTN_WEST
	BTN_SOUTH
	BTN_EAST
	BTN_TL
	BTN_TL2
	BTN_TR
	BTN_TR2
	BTN_SELECT
	BTN_MODE
	BTN_START
	BTN_THUMBL
	BTN_THUMBR
)

var XBoxButtons = []evdev.EvCode{
	0,                    // Unknown button, just for error checking
	evdev.BTN_DPAD_UP,    // Gamepad left DPAD up button
	evdev.BTN_DPAD_RIGHT, // Gamepad left DPAD right button
	evdev.BTN_DPAD_DOWN,  // Gamepad left DPAD down button
	evdev.BTN_DPAD_LEFT,  // Gamepad left DPAD left button
	evdev.BTN_Y,          // Gamepad right button up (i.e. PS3: Triangle, Xbox: Y)
	evdev.BTN_X,          // Gamepad right button right (i.e. PS3: Square, Xbox: X)
	evdev.BTN_A,          // Gamepad right button down (i.e. PS3: Cross, Xbox: A)
	evdev.BTN_B,          // Gamepad right button left (i.e. PS3: Circle, Xbox: B)
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
	evdev.BTN_WEST,       // Gamepad right button right (i.e. PS3: Square, Xbox: X)
	evdev.BTN_SOUTH,      // Gamepad right button down (i.e. PS3: Cross, Xbox: A)
	evdev.BTN_EAST,       // Gamepad right button left (i.e. PS3: Circle, Xbox: B)
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
	evdev.BTN_JOYSTICK,   // Gamepad right button right (i.e. PS3: Square, Xbox: X)
	evdev.BTN_THUMB,      // Gamepad right button down (i.e. PS3: Cross, Xbox: A)
	evdev.BTN_THUMB2,     // Gamepad right button left (i.e. PS3: Circle, Xbox: B)
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

var JoyToGame = map[evdev.EvCode]evdev.EvCode{
	0:                    0,
	evdev.BTN_DPAD_UP:    evdev.BTN_DPAD_UP,
	evdev.BTN_DPAD_RIGHT: evdev.BTN_DPAD_RIGHT,
	evdev.BTN_DPAD_DOWN:  evdev.BTN_DPAD_DOWN,
	evdev.BTN_DPAD_LEFT:  evdev.BTN_DPAD_LEFT,
	evdev.BTN_TOP:        evdev.BTN_NORTH,
	evdev.BTN_JOYSTICK:   evdev.BTN_WEST,
	evdev.BTN_THUMB:      evdev.BTN_SOUTH,
	evdev.BTN_THUMB2:     evdev.BTN_EAST,
	evdev.BTN_TOP2:       evdev.BTN_TL,
	evdev.BTN_BASE:       evdev.BTN_TL2,
	evdev.BTN_PINKIE:     evdev.BTN_TR,
	evdev.BTN_BASE2:      evdev.BTN_TR2,
	evdev.BTN_BASE3:      evdev.BTN_SELECT,
	evdev.BTN_DEAD:       evdev.BTN_MODE,
	evdev.BTN_BASE4:      evdev.BTN_START,
	evdev.BTN_BASE5:      evdev.BTN_THUMBL,
	evdev.BTN_BASE6:      evdev.BTN_THUMBR,
}

const (
	ABS_X int = iota
	ABS_Y
	ABS_Z
	ABS_RZ
	ABS_HAT0X
	ABS_HAT0Y
	ABS_RY
	ABS_RX
)

var AxisEvents = []evdev.EvCode{
	evdev.ABS_X,
	evdev.ABS_Y,
	evdev.ABS_Z,
	evdev.ABS_RZ,
	evdev.ABS_HAT0X,
	evdev.ABS_HAT0Y,
	evdev.ABS_RY,
	evdev.ABS_RX,
}

type AxisInfoMap map[evdev.EvCode]evdev.AbsInfo
type AxisStateMap map[evdev.EvCode]int32
type ButtonStateMap map[evdev.EvCode]bool