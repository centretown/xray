package stickg

import (
	"fmt"

	"github.com/holoplot/go-evdev"
)

const JOYSTICK_AXIS_MAX = 16

type StickTypes struct {
	EvType   evdev.EvType
	AbsInfos map[evdev.EvCode]evdev.AbsInfo
}

type StickG struct {
	Device           *evdev.InputDevice
	InputID          evdev.InputID
	Version          [3]int
	Name             string
	PhysicalLocation string
	UniqueID         string
	StickTypes       []StickTypes
	Properties       []evdev.EvProp

	spinning bool
}

func OpenStickG(path string) (stg *StickG, err error) {
	var device *evdev.InputDevice
	device, err = evdev.Open(path)
	if err != nil {
		return
	}
	defer device.NonBlock()

	stg = &StickG{
		Device:     device,
		StickTypes: make([]StickTypes, 0),
	}

	stg.Version[0], stg.Version[1], stg.Version[2] = device.DriverVersion()
	stg.InputID, err = device.InputID()
	if err != nil {
		return
	}
	stg.Name, err = device.Name()
	if err != nil {
		return
	}
	stg.PhysicalLocation, err = device.PhysicalLocation()
	if err != nil {
		return
	}
	stg.UniqueID, err = device.UniqueID()
	if err != nil {
		return
	}

	for _, t := range device.CapableTypes() {
		st := StickTypes{}
		st.EvType = t
		var (
			absInfos = make(map[evdev.EvCode]evdev.AbsInfo)
			err      error
		)
		if t == evdev.EV_ABS {
			absInfos, err = device.AbsInfos()
			if err == nil {
				st.AbsInfos = absInfos
			}
		}
		stg.StickTypes = append(stg.StickTypes, st)
	}
	stg.Properties = device.Properties()
	return
}

func (stg *StickG) Dump() {
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
	fmt.Printf("Supported events:\n")
	for _, stgType := range stg.StickTypes {
		t := stgType.EvType
		fmt.Printf("  Event type %d (%s)\n", t, evdev.TypeName(t))

		// state, err := stg.State(t)
		// if err == nil {
		// 	for code, value := range state {
		// 		fmt.Printf("    Event code %d (%s) state %v\n", code, evdev.CodeName(t, code), value)
		// 	}
		// }

		if t != evdev.EV_ABS {
			continue
		}

		absInfos := stgType.AbsInfos

		for code, absInfo := range absInfos {
			fmt.Printf("    Event code %d (%s)\n", code, evdev.CodeName(t, code))
			fmt.Printf("      Value: %d\n", absInfo.Value)
			fmt.Printf("      Min: %d\n", absInfo.Minimum)
			fmt.Printf("      Max: %d\n", absInfo.Maximum)

			if absInfo.Fuzz != 0 {
				fmt.Printf("      Fuzz: %d\n", absInfo.Fuzz)
			}
			if absInfo.Flat != 0 {
				fmt.Printf("      Flat: %d\n", absInfo.Flat)
			}
			if absInfo.Resolution != 0 {
				fmt.Printf("      Resolution: %d\n", absInfo.Resolution)
			}
		}
	}

	fmt.Printf("Properties:\n")
	props := stg.Properties
	for _, p := range props {
		fmt.Printf("  Property type %d (%s)\n", p, evdev.PropName(p))
	}

}

func (stg *StickG) Close() {
	stg.spinning = false
	stg.Device.Close()
	fmt.Println("stg.Stop closed")
}

// Event code 299 (BTN_BASE6) state false
// Event code 288 (BTN_JOYSTICK/BTN_TRIGGER) state false
// Event code 292 (BTN_TOP2) state false
// Event code 293 (BTN_PINKIE) state false
// Event code 295 (BTN_BASE2) state false
// Event code 296 (BTN_BASE3) state false
// Event code 297 (BTN_BASE4) state false
// Event code 298 (BTN_BASE5) state false
// Event code 289 (BTN_THUMB) state true
// Event code 290 (BTN_THUMB2) state false
// Event code 291 (BTN_TOP) state false
// Event code 294 (BTN_BASE) state false

// BTN_JOYSTICK = 0x120
// BTN_TRIGGER  = 0x120
// BTN_THUMB    = 0x121
// BTN_THUMB2   = 0x122
// BTN_TOP      = 0x123
// BTN_TOP2     = 0x124
// BTN_PINKIE   = 0x125
// BTN_BASE     = 0x126
// BTN_BASE2    = 0x127
// BTN_BASE3    = 0x128
// BTN_BASE4    = 0x129
// BTN_BASE5    = 0x12a
// BTN_BASE6    = 0x12b
// BTN_DEAD     = 0x12f

// BTN_GAMEPAD = 0x130
// BTN_SOUTH   = 0x130
// BTN_A       = BTN_SOUTH
// BTN_EAST    = 0x131
// BTN_B       = BTN_EAST
// BTN_C       = 0x132
// BTN_NORTH   = 0x133
// BTN_X       = BTN_NORTH
// BTN_WEST    = 0x134
// BTN_Y       = BTN_WEST
// BTN_Z       = 0x135
// BTN_TL      = 0x136
// BTN_TR      = 0x137
// BTN_TL2     = 0x138
// BTN_TR2     = 0x139
// BTN_SELECT  = 0x13a
// BTN_START   = 0x13b
// BTN_MODE    = 0x13c
// BTN_THUMBL  = 0x13d
// BTN_THUMBR  = 0x13e

// func (stg *StickG) Start() {
// 	if !stg.spinning {
// 		go stg.Spin()
// 		stg.spinning = true
// 	}
// }

// func (stg *StickG) Spin() {
// 	stg.InputDevice.NonBlock()
// 	for {

// 		e, err := stg.ReadOne()
// 		if err != nil {
// 			fmt.Printf("Error reading from device: %v\n", err)
// 			return
// 		}

// 		ts := fmt.Sprintf("Event: time %d.%06d", e.Time.Sec, e.Time.Usec)

// 		switch e.Type {
// 		case evdev.EV_SYN:
// 			switch e.Code {
// 			case evdev.SYN_MT_REPORT:
// 				fmt.Printf("%s, ++++++++++++++ %s ++++++++++++\n", ts, e.CodeName())
// 			case evdev.SYN_DROPPED:
// 				fmt.Printf("%s, >>>>>>>>>>>>>> %s <<<<<<<<<<<<\n", ts, e.CodeName())
// 			default:
// 				fmt.Printf("%s, -------------- %s ------------\n", ts, e.CodeName())
// 			}
// 		default:
// 			fmt.Printf("%s, %s\n", ts, e.String())
// 		}

// 		time.Sleep(time.Millisecond)
// 	}

// }
