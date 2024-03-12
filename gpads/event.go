package gpads

import "github.com/holoplot/go-evdev"

// button mappings

var ButtonCount = len(GameButtons)

const BTN_GAMEPAD = evdev.BTN_GAMEPAD

const (
	BTN_WEST int = iota
	BTN_SOUTH
	BTN_EAST
	BTN_NORTH
	BTN_TL
	BTN_TR
	BTN_TL2
	BTN_TR2
	BTN_SELECT
	BTN_START
	BTN_THUMBL
	BTN_THUMBR
	BTN_MODE
)

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
}

var JoyToGame = map[evdev.EvCode]evdev.EvCode{
	evdev.BTN_JOYSTICK: evdev.BTN_WEST,
	evdev.BTN_THUMB:    evdev.BTN_SOUTH,
	evdev.BTN_THUMB2:   evdev.BTN_EAST,
	evdev.BTN_TOP:      evdev.BTN_NORTH,
	evdev.BTN_TOP2:     evdev.BTN_TL,
	evdev.BTN_PINKIE:   evdev.BTN_TR,
	evdev.BTN_BASE:     evdev.BTN_TL2,
	evdev.BTN_BASE2:    evdev.BTN_TR2,
	evdev.BTN_BASE3:    evdev.BTN_SELECT,
	evdev.BTN_BASE4:    evdev.BTN_START,
	evdev.BTN_BASE5:    evdev.BTN_THUMBL,
	evdev.BTN_BASE6:    evdev.BTN_THUMBR,
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
