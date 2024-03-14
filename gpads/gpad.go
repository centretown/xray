package gpads

import (
	"fmt"
	"xray/b2i"

	"github.com/holoplot/go-evdev"
)

type GPad struct {
	PadType          PadType
	Device           *evdev.InputDevice
	InputID          evdev.InputID
	Version          [3]int
	Name             string
	PhysicalLocation string
	UniqueID         string
	AxesInfo         AxisInfoMap
	Properties       []evdev.EvProp

	// ButtonBase  evdev.EvCode
	ButtonCodes []evdev.EvCode

	preAxisState AxisStateMap
	curAxisState AxisStateMap
	axisAdjust   map[evdev.EvCode]int32

	curButtonState evdev.StateMap
	PressedOnce    uint64       // 1<<(evdev.EvCode - ButtonBase)
	ReleasedOnce   uint64       // 1<<(evdev.EvCode - ButtonBase)
	LastPressed    evdev.EvCode // Button number

	intialized bool
}

func newPadG(device *evdev.InputDevice) (gpad *GPad) {
	gpad = &GPad{
		Device:         device,
		preAxisState:   make(AxisStateMap),
		curAxisState:   make(AxisStateMap),
		axisAdjust:     make(map[evdev.EvCode]int32),
		curButtonState: make(evdev.StateMap),
	}
	return gpad
}

func OpenGPad(path string) (*GPad, error) {
	var (
		state  evdev.StateMap
		err    error
		device *evdev.InputDevice
	)

	device, err = evdev.Open(path)
	if err != nil {
		fmt.Println("OpenStickG", path, err)
		return nil, err
	}

	gpad := newPadG(device)
	gpad.Version[0], gpad.Version[1], gpad.Version[2] = device.DriverVersion()
	gpad.InputID, _ = device.InputID()
	gpad.Name, _ = device.Name()
	gpad.PhysicalLocation, _ = device.PhysicalLocation()
	gpad.UniqueID, _ = device.UniqueID()

	state, err = device.State(evdev.EV_KEY)
	if err != nil {
		fmt.Println("EV_KEY.State: ", err)
	} else {
		_, isJoy := state[BTN_JOYSTICK]
		if isJoy {
			gpad.PadType = PAD_JOYSTICK
			gpad.ButtonCodes = JoyButtons
		} else if gpad.InputID.Vendor == 0x45e && gpad.InputID.Product == 0x28e &&
			gpad.InputID.Version == 0x110 {
			// TODO: Improve above condition
			gpad.PadType = PAD_XBOX
			gpad.ButtonCodes = XBoxButtons
		} else {
			gpad.PadType = PAD_PS3
			gpad.ButtonCodes = PS3Buttons
		}

		for _, code := range gpad.ButtonCodes {
			gpad.curButtonState[code] = state[code]
		}
	}

	gpad.AxesInfo, err = device.AbsInfos()
	if err != nil {
		fmt.Println("device.AbsInfos(): ", err)
	} else {
		for code, info := range gpad.AxesInfo {
			var adj int32 = 1
			diff := info.Maximum - info.Minimum
			if diff > 4 {
				adj = diff / 256
				if adj == 0 {
					adj = 1
				}
			}
			gpad.axisAdjust[code] = adj
		}
	}

	gpad.Properties = device.Properties()

	return gpad, nil
}

func (gpad *GPad) ReadState() {
	state, err := gpad.Device.State(evdev.EV_KEY)
	if err == nil {
		var nowDown, wasDown bool

		absInfos, err := gpad.Device.AbsInfos()
		if err == nil {
			for k, v := range absInfos {
				gpad.preAxisState[k] = gpad.curAxisState[k]
				gpad.curAxisState[k] = v.Value / gpad.axisAdjust[k]
			}
		}

		for i, code := range gpad.ButtonCodes {
			wasDown = gpad.curButtonState[code]
			switch i {
			case 0:
			case BTN_DPAD_UP:
				nowDown = gpad.curAxisState[evdev.ABS_HAT0Y] < 0
			case BTN_DPAD_RIGHT:
				nowDown = gpad.curAxisState[evdev.ABS_HAT0X] > 0
			case BTN_DPAD_DOWN:
				nowDown = gpad.curAxisState[evdev.ABS_HAT0Y] > 0
			case BTN_DPAD_LEFT:
				nowDown = gpad.curAxisState[evdev.ABS_HAT0X] < 0
			default:
				nowDown = state[code]
			}
			gpad.PressedOnce |= b2i.Bool2uint64(!wasDown && nowDown) << i
			gpad.ReleasedOnce |= b2i.Bool2uint64(wasDown && !nowDown) << i
			gpad.curButtonState[code] = nowDown
		}
	}

}

func (gpad *GPad) AxisValue(axis int) int32 {
	k := AxisEvents[axis]
	return gpad.curAxisState[k]
}

func (gpad *GPad) AxisMove(axis int) float32 {
	k := AxisEvents[axis]
	return float32(gpad.curAxisState[k] - gpad.preAxisState[k])
}

func (gpad *GPad) ButtonDown(button int) bool {
	code, ok := gpad.curButtonState[gpad.ButtonCodes[button]]
	return ok && code
}
func (gpad *GPad) ButtonPressed(button int) bool {
	return gpad.PressedOnce&(1<<button) != 0
}

func (gpad *GPad) ButtonReleased(button int) bool {
	return gpad.ReleasedOnce&(1<<button) != 0
}

func CheckOne[T comparable](pre, cur map[evdev.EvCode]T, k evdev.EvCode) {
	v := cur[k]
	if pre[k] != v {
		fmt.Printf("[%v:%v]\n", k, v)
	}
}

func CheckAll[T comparable](pre, cur map[evdev.EvCode]T) {
	for k, v := range cur {
		if pre[k] != v {
			fmt.Printf("[%v:%v]\n", k, v)
		}
	}
}

func (gpad *GPad) DumpState() {
	// CheckAll(gpad.preButtonState, gpad.curButtonState)
	CheckAll(gpad.preAxisState, gpad.curAxisState)
}

func (gpad *GPad) Dump() {
	fmt.Printf("Pad Type: %s\n", gpad.PadType)

	fmt.Printf("Input driver version is %d.%d.%d\n",
		gpad.Version[0],
		gpad.Version[1],
		gpad.Version[2],
	)

	fmt.Printf("Input device ID: bus 0x%x vendor 0x%x product 0x%x version 0x%x\n",
		gpad.InputID.BusType, gpad.InputID.Vendor, gpad.InputID.Product, gpad.InputID.Version)

	fmt.Printf("Input device name: \"%s\"\n", gpad.Name)
	fmt.Printf("Input device physical location: %s\n", gpad.PhysicalLocation)
	fmt.Printf("Input device unique ID: %s\n", gpad.UniqueID)

	fmt.Println("Axes:")
	for code, absInfo := range gpad.AxesInfo {
		fmt.Printf("    Event code %d (%s) Value: %d Min: %d Max: %d Fuzz: %d Flat: %d Resolution: %d\n",
			code, evdev.CodeName(evdev.EV_ABS, code), absInfo.Value, absInfo.Minimum, absInfo.Maximum,
			absInfo.Fuzz, absInfo.Flat, absInfo.Resolution)
	}

	fmt.Println("Buttons:")
	for code, value := range gpad.curButtonState {
		fmt.Printf("    Event code %d (%s) state %v\n",
			code, evdev.CodeName(evdev.EV_KEY, code), value)
	}

	fmt.Println("Properties:")
	props := gpad.Properties
	if len(props) < 1 {
		fmt.Println("    none")
	}
	for _, p := range props {
		fmt.Printf("   Property type %d (%s)\n", p, evdev.PropName(p))
	}
}

func (gpad *GPad) Close() {
	gpad.Device.Close()
	gpad.intialized = false
}
