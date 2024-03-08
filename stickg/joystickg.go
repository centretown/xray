package stickg

import (
	"fmt"
	"xray/jstick"

	"github.com/holoplot/go-evdev"
)

var _ jstick.Jstick = NewJoyStickG()

const JOYSTICK_MAX = 4

type JoyStickG struct {
	Sticks     [JOYSTICK_MAX]*StickG
	Err        error
	stickCount int
	intialized bool
}

func NewJoyStickG() *JoyStickG {
	js := &JoyStickG{}
	return js
}

func (js *JoyStickG) intialize() {
	fmt.Println("intialize")
	var (
		devicePaths []evdev.InputPath
		err         error
	)

	defer func() { js.intialized = true }()

	devicePaths, err = evdev.ListDevicePaths()
	if err != nil {
		fmt.Println("ListDevicePaths", err)
		return
	}

	fmt.Println("ListDevicePaths", devicePaths)

	js.stickCount = 0
	for i := range len(devicePaths) {
		stg, err := OpenStickG(devicePaths[i].Path)
		if err != nil {
			fmt.Println("Open", err)
			continue
		}
		js.Sticks[js.stickCount] = stg
		js.stickCount++
	}
}

func (js *JoyStickG) BeginJoystick() {
	if !js.intialized {
		js.intialize()
	}
}

func (js *JoyStickG) IsJoystickAvailable(Joystick int) bool {
	return js.stickCount > Joystick
}

const undefined = "undefined"

func (js *JoyStickG) GetJoystickName(Joystick int) string {
	if js.stickCount <= Joystick {
		return undefined
	}
	return js.Sticks[Joystick].Name
}

func (js *JoyStickG) IsJoystickButtonPressed(Joystick int, button int) bool {
	return false
}

func (js *JoyStickG) IsJoystickButtonDown(Joystick int, button int) bool {
	return false
}

func (js *JoyStickG) IsJoystickButtonReleased(Joystick int, button int) bool {
	return false
}

func (js *JoyStickG) IsJoystickButtonUp(Joystick int, button int) bool {
	return false
}

func (js *JoyStickG) GetJoystickButtonPressed() int {
	return 0
}

func (js *JoyStickG) GetJoystickAxisCount(Joystick int) int {
	return 0
}

func (js *JoyStickG) GetJoystickButtonCount(Joystick int) int {
	return 0
}

func (js *JoyStickG) GetJoystickAxisMovement(Joystick int, axis int) float32 {
	return 0
}

func (js *JoyStickG) GetJoystickAxisValue(Joystick int, axis int) int32 {
	return 0
}

func (js *JoyStickG) SetJoystickMappings(mappings string) int {
	return 0
}

// func (js *JoyStickG) dumpDevice(d *evdev.InputDevice) {
// 	defer d.NonBlock()

// 	vMajor, vMinor, vMicro := d.DriverVersion()
// 	fmt.Printf("Input driver version is %d.%d.%d\n", vMajor, vMinor, vMicro)

// 	inputID, err := d.InputID()
// 	if err == nil {
// 		fmt.Printf("Input device ID: bus 0x%x vendor 0x%x product 0x%x version 0x%x\n",
// 			inputID.BusType, inputID.Vendor, inputID.Product, inputID.Version)
// 	}

// 	name, err := d.Name()
// 	if err == nil {
// 		fmt.Printf("Input device name: \"%s\"\n", name)
// 	}

// 	phys, err := d.PhysicalLocation()
// 	if err == nil {
// 		fmt.Printf("Input device physical location: %s\n", phys)
// 	}

// 	uniq, err := d.UniqueID()
// 	if err == nil {
// 		fmt.Printf("Input device unique ID: %s\n", uniq)
// 	}

// 	fmt.Printf("Supported events:\n")

// 	for _, t := range d.CapableTypes() {
// 		fmt.Printf("  Event type %d (%s)\n", t, evdev.TypeName(t))

// 		state, err := d.State(t)
// 		if err == nil {
// 			for code, value := range state {
// 				fmt.Printf("    Event code %d (%s) state %v\n", code, evdev.CodeName(t, code), value)
// 			}
// 		}

// 		if t != evdev.EV_ABS {
// 			continue
// 		}

// 		absInfos, err := d.AbsInfos()
// 		if err != nil {
// 			continue
// 		}

// 		for code, absInfo := range absInfos {
// 			fmt.Printf("    Event code %d (%s)\n", code, evdev.CodeName(t, code))
// 			fmt.Printf("      Value: %d\n", absInfo.Value)
// 			fmt.Printf("      Min: %d\n", absInfo.Minimum)
// 			fmt.Printf("      Max: %d\n", absInfo.Maximum)

// 			if absInfo.Fuzz != 0 {
// 				fmt.Printf("      Fuzz: %d\n", absInfo.Fuzz)
// 			}
// 			if absInfo.Flat != 0 {
// 				fmt.Printf("      Flat: %d\n", absInfo.Flat)
// 			}
// 			if absInfo.Resolution != 0 {
// 				fmt.Printf("      Resolution: %d\n", absInfo.Resolution)
// 			}
// 		}
// 	}

// 	fmt.Printf("Properties:\n")

// 	props := d.Properties()

// 	for _, p := range props {
// 		fmt.Printf("  Property type %d (%s)\n", p, evdev.PropName(p))
// 	}

// }

func (js *JoyStickG) DumpJoystick() {
	for i := 0; i < js.stickCount; i++ {
		js.Sticks[i].Dump()
	}
}

func (js *JoyStickG) GetButtonName(Joystick int, button int) string {
	return ""
}
