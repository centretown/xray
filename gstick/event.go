package gstick

import "github.com/holoplot/go-evdev"

// button mappings

var AxisEvents = []evdev.EvCode{
	evdev.ABS_X,
	evdev.ABS_Y,
	evdev.ABS_Z,
	evdev.ABS_RX,
	evdev.ABS_RY,
	evdev.ABS_RZ,
	evdev.ABS_HAT0X,
	evdev.ABS_HAT0Y,
}

type AxisInfoMap map[evdev.EvCode]evdev.AbsInfo
type AxisStateMap map[evdev.EvCode]int32
type ButtonStateMap map[evdev.EvCode]bool

var ButtonCount = len(GameButtons)

const BTN_GAMEPAD = evdev.BTN_GAMEPAD

var GameButtons = []evdev.EvCode{
	evdev.BTN_WEST,
	evdev.BTN_SOUTH,
	evdev.BTN_EAST,
	evdev.BTN_NORTH,
	evdev.BTN_TL,
	evdev.BTN_TR,
	evdev.BTN_TL2,
	evdev.BTN_TR2,
	evdev.BTN_SELECT,
	evdev.BTN_START,
	evdev.BTN_THUMBL,
	evdev.BTN_THUMBR,
	evdev.BTN_MODE,
}

const BTN_JOYSTICK = evdev.BTN_JOYSTICK

var JoyButtons = []evdev.EvCode{
	evdev.BTN_JOYSTICK,
	evdev.BTN_THUMB,
	evdev.BTN_THUMB2,
	evdev.BTN_TOP,
	evdev.BTN_TOP2,
	evdev.BTN_PINKIE,
	evdev.BTN_BASE,
	evdev.BTN_BASE2,
	evdev.BTN_BASE3,
	evdev.BTN_BASE4,
	evdev.BTN_BASE5,
	evdev.BTN_BASE6,
	evdev.KEY_RESERVED, // == 0 OR not available
}
