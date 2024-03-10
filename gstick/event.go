package gstick

import "github.com/holoplot/go-evdev"

// button mappings

type Button int

const UNKNOWN = "unknown"

const (
	BTN_WEST Button = iota
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
	BTN_COUNT
)

var ButtonLabels = []string{
	"WEST",
	"SOUTH",
	"EAST",
	"NORTH",
	"TL",
	"TR",
	"TL2",
	"TR2",
	"SELECT",
	"START",
	"THUMBL",
	"THUMBR",
	"MODE",
}

func (b Button) String() string {
	if b < BTN_COUNT {
		return ButtonLabels[b]
	}
	return UNKNOWN
}

const JOYSTICK_BUTTON = evdev.BTN_JOYSTICK

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

const GAMEPAD_BUTTON = evdev.BTN_GAMEPAD

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

type Axis int

const (
	ABS_RX Axis = iota
	ABS_RY
	ABS_RZ
	ABS_HAT0X
	ABS_HAT0Y
	ABS_X
	ABS_Y
	ABS_Z
	AXIS_COUNT
)

var AxisLabels = []string{
	"ABS_X",
	"ABS_Y",
	"ABS_Z",
	"ABS_RX",
	"ABS_RY",
	"ABS_RZ",
	"ABS_HAT0X",
	"ABS_HAT0Y",
}

func (a Axis) String() string {
	if a < AXIS_COUNT {
		return AxisLabels[a]
	}
	return UNKNOWN
}

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

type ButtonEvent interface {
	Event(evdev.EvCode) evdev.EvCode
}

var _ ButtonEvent = &GameButtonEvent{}

type GameButtonEvent struct{}

func (gbe *GameButtonEvent) Event(code evdev.EvCode) evdev.EvCode { return code }

var JoyButtonMap = map[evdev.EvCode]evdev.EvCode{
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
	evdev.KEY_RESERVED: evdev.BTN_MODE,
}

var _ ButtonEvent = &JoyButtonEvent{}

type JoyButtonEvent struct{}

func (jbe *JoyButtonEvent) Event(code evdev.EvCode) evdev.EvCode { return JoyButtonMap[code] }
