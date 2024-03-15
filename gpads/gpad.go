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

	ButtonCodes []evdev.EvCode
	AxisCodes   []evdev.EvCode

	AxisStatePrev [RL_AXIS_COUNT]int32
	AxisState     [RL_AXIS_COUNT]int32
	// TODO: include jitter detection
	axisAdjust [RL_AXIS_COUNT]int32

	ButtonState  [RL_BUTTON_COUNT]bool
	PressedOnce  uint64       // 1<<(evdev.EvCode - ButtonBase)
	ReleasedOnce uint64       // 1<<(evdev.EvCode - ButtonBase)
	LastPressed  evdev.EvCode // Button number?

	intialized bool
}

func newPadG(device *evdev.InputDevice) (gpad *GPad) {
	gpad = &GPad{
		Device: device,
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
		gpad.AxisCodes = GameAxes
		_, isJoy := state[BTN_JOYSTICK]
		if isJoy {
			gpad.PadType = PAD_JOYSTICK
			gpad.ButtonCodes = JoyButtons
			gpad.AxisCodes = JoyAxes
		} else if gpad.InputID.Vendor == 0x45e && gpad.InputID.Product == 0x28e &&
			gpad.InputID.Version == 0x110 {
			// TODO: Improve above condition
			// above derived from p4 clone posing as xbox
			gpad.PadType = PAD_XBOX
			gpad.ButtonCodes = XBoxButtons
		} else {
			gpad.PadType = PAD_PS3
			gpad.ButtonCodes = PS3Buttons
		}

		for button, code := range gpad.ButtonCodes {
			gpad.ButtonState[button] = state[code]
		}
	}

	gpad.AxesInfo, err = device.AbsInfos()
	if err != nil {
		fmt.Println("device.AbsInfos(): ", err)
	} else {
		for axis, code := range gpad.AxisCodes {
			var adj int32 = 1
			info, ok := gpad.AxesInfo[code]
			if ok {
				diff := info.Maximum - info.Minimum
				if diff > 256 {
					adj = diff / 256
					if adj == 0 {
						adj = 1
					}
				}
				gpad.AxisState[axis] = info.Value
			}
			gpad.axisAdjust[axis] = adj
		}

	}

	gpad.Properties = device.Properties()

	return gpad, nil
}

func (gpad *GPad) ReadState() {
	var (
		isDown, wasDown, ok bool
		info                evdev.AbsInfo
		button              int
		code                evdev.EvCode
	)

	absInfos, err := gpad.Device.AbsInfos()
	if err == nil {
		for axis, code := range gpad.AxisCodes {
			info, ok = absInfos[code]
			if ok {
				gpad.AxisStatePrev[axis] = gpad.AxisState[axis]
				gpad.AxisState[axis] = info.Value / gpad.axisAdjust[axis]
			}
		}
	}

	state, err := gpad.Device.State(evdev.EV_KEY)
	if err == nil {
		for button, code = range gpad.ButtonCodes {
			wasDown = gpad.ButtonState[button]
			switch button {
			case 0:
			case RL_LeftFaceUp:
				isDown = gpad.AxisState[RL_HatY] < 0
			case RL_LeftFaceRight:
				isDown = gpad.AxisState[RL_HatX] > 0
			case RL_LeftFaceDown:
				isDown = gpad.AxisState[RL_HatY] > 0
			case RL_LeftFaceLeft:
				isDown = gpad.AxisState[RL_HatX] < 0
			case RL_LeftTrigger2:
				isDown = (gpad.PadType == PAD_XBOX && gpad.AxisState[RL_LeftTrigger] > 0) ||
					state[code]
			case RL_RightTrigger2:
				isDown = (gpad.PadType == PAD_XBOX && gpad.AxisState[RL_RightTrigger] > 0) ||
					state[code]
			default:
				isDown = state[code]
			}

			gpad.PressedOnce |= b2i.Bool2uint64(!wasDown && isDown) << button
			gpad.ReleasedOnce |= b2i.Bool2uint64(wasDown && !isDown) << button
			gpad.ButtonState[button] = isDown
		}
	}

}

func (gpad *GPad) AxisValue(axis int) int32 {
	return gpad.AxisState[axis]
}

func (gpad *GPad) AxisMove(axis int) float32 {
	return float32(gpad.AxisState[axis] - gpad.AxisStatePrev[axis])
}

func (gpad *GPad) ButtonDown(button int) bool {
	return gpad.ButtonState[button]
}
func (gpad *GPad) ButtonPressed(button int) bool {
	return gpad.PressedOnce&(1<<button) != 0
}

func (gpad *GPad) ButtonReleased(button int) bool {
	return gpad.ReleasedOnce&(1<<button) != 0
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
	for _, code := range gpad.AxisCodes {
		info, ok := gpad.AxesInfo[code]
		if ok {
			fmt.Printf("    Event code %d (%s) Value: %d Min: %d Max: %d Fuzz: %d Flat: %d Resolution: %d\n",
				code, evdev.CodeName(evdev.EV_ABS, code), info.Value, info.Minimum, info.Maximum,
				info.Fuzz, info.Flat, info.Resolution)
		}
	}

	fmt.Println("Buttons:")
	for index, value := range gpad.ButtonState {
		code := gpad.ButtonCodes[index]
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
