package gpads

import (
	"fmt"
	"xray/b2i"

	"github.com/holoplot/go-evdev"
)

type GPad struct {
	Device           *evdev.InputDevice
	InputID          evdev.InputID
	Version          [3]int
	Name             string
	PhysicalLocation string
	UniqueID         string
	AxesInfo         AxisInfoMap
	Properties       []evdev.EvProp

	ButtonBase  evdev.EvCode
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

func newPadG(device *evdev.InputDevice) (stg *GPad) {
	stg = &GPad{
		Device:         device,
		preAxisState:   make(AxisStateMap),
		curAxisState:   make(AxisStateMap),
		axisAdjust:     make(map[evdev.EvCode]int32),
		curButtonState: make(evdev.StateMap),
	}
	return stg
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

	stg := newPadG(device)
	stg.Version[0], stg.Version[1], stg.Version[2] = device.DriverVersion()
	stg.InputID, _ = device.InputID()
	stg.Name, _ = device.Name()
	stg.PhysicalLocation, _ = device.PhysicalLocation()
	stg.UniqueID, _ = device.UniqueID()

	state, err = device.State(evdev.EV_KEY)
	if err != nil {
		fmt.Println("EV_KEY.State: ", err)
	} else {
		for code, nowDown := range state {
			stg.curButtonState[code] = nowDown
		}
	}

	_, ok := stg.curButtonState[BTN_JOYSTICK]
	if ok {
		stg.ButtonBase = BTN_JOYSTICK
		stg.ButtonCodes = JoyButtons
	} else {
		stg.ButtonBase = BTN_GAMEPAD
		stg.ButtonCodes = GameButtons
	}

	stg.AxesInfo, err = device.AbsInfos()
	if err != nil {
		fmt.Println("device.AbsInfos(): ", err)
	} else {
		for code, info := range stg.AxesInfo {
			var adj int32 = 1
			diff := info.Maximum - info.Minimum
			if diff > 255 {
				adj = diff / 256
			}
			stg.axisAdjust[code] = adj
		}
	}

	stg.Properties = device.Properties()

	return stg, nil
}

func (stg *GPad) ReadState() {
	state, err := stg.Device.State(evdev.EV_KEY)
	if err == nil {
		var (
			wasDown bool
			offset  uint64
		)

		for code, nowDown := range state {
			wasDown = stg.curButtonState[code]
			offset = uint64(code - stg.ButtonBase)
			stg.PressedOnce |= b2i.Bool2uint64(!wasDown && nowDown) << offset
			stg.ReleasedOnce |= b2i.Bool2uint64(wasDown && !nowDown) << offset
			stg.curButtonState[code] = nowDown
		}
	}

	absInfos, err := stg.Device.AbsInfos()
	if err == nil {
		for k, v := range absInfos {
			stg.preAxisState[k] = stg.curAxisState[k]
			stg.curAxisState[k] = v.Value / stg.axisAdjust[k]
		}
	}
}

func (stg *GPad) AxisValue(axis int) int32 {
	k := AxisEvents[axis]
	return stg.curAxisState[k]
}

func (stg *GPad) AxisMove(axis int) float32 {
	k := AxisEvents[axis]
	return float32(stg.curAxisState[k] - stg.preAxisState[k])
}

func (stg *GPad) ButtonDown(button int) bool {
	code, ok := stg.curButtonState[stg.ButtonCodes[button]]
	return ok && code
}
func (stg *GPad) ButtonPressed(button int) bool {
	offset := uint64(stg.ButtonCodes[button] - stg.ButtonBase)
	return stg.PressedOnce&(1<<offset) != 0
}

func (stg *GPad) ButtonReleased(button int) bool {
	offset := uint64(stg.ButtonCodes[button] - stg.ButtonBase)
	return stg.ReleasedOnce&(1<<offset) != 0
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

func (stg *GPad) DumpState() {
	// CheckAll(stg.preButtonState, stg.curButtonState)
	CheckAll(stg.preAxisState, stg.curAxisState)
}

func (stg *GPad) Dump() {
	fmt.Printf("Input driver version is %d.%d.%d\n",
		stg.Version[0],
		stg.Version[1],
		stg.Version[2],
	)

	fmt.Printf("Input device ID: bus 0x%x vendor 0x%x product 0x%x version 0x%x\n",
		stg.InputID.BusType, stg.InputID.Vendor, stg.InputID.Product, stg.InputID.Version)

	fmt.Printf("Input device name: \"%s\"\n", stg.Name)
	fmt.Printf("Input device physical location: %s\n", stg.PhysicalLocation)
	fmt.Printf("Input device unique ID: %s\n", stg.UniqueID)

	fmt.Println("Axes:")
	for code, absInfo := range stg.AxesInfo {
		fmt.Printf("    Event code %d (%s) Value: %d Min: %d Max: %d Fuzz: %d Flat: %d Resolution: %d\n",
			code, evdev.CodeName(evdev.EV_ABS, code), absInfo.Value, absInfo.Minimum, absInfo.Maximum,
			absInfo.Fuzz, absInfo.Flat, absInfo.Resolution)
	}

	fmt.Println("Buttons:")
	for code, value := range stg.curButtonState {
		fmt.Printf("    Event code %d (%s) state %v\n",
			code, evdev.CodeName(evdev.EV_KEY, code), value)
	}

	fmt.Println("Properties:")
	props := stg.Properties
	if len(props) < 1 {
		fmt.Println("    none")
	}
	for _, p := range props {
		fmt.Printf("   Property type %d (%s)\n", p, evdev.PropName(p))
	}
}

func (stg *GPad) Close() {
	stg.Device.Close()
	stg.intialized = false
}
