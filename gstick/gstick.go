package gstick

import (
	"fmt"

	"github.com/holoplot/go-evdev"
)

// dave@yeller:~$ echo -n "2563:0119:ig" | sudo tee /sys/module/usbcore/parameters/quirks
// 2563:0119:ig
// davdmesg | tail
// [80024.818700] hid-generic 0003:2563:0119.0025: input,hiddev1,hidraw5: USB HID v1.11 Gamepad [shanwan Wired Controller] on usb-0000:00:1d.0-1.5/input3
// [80025.571774] usb 2-1.5: USB disconnect, device number 30
// [80026.024984] usb 2-1.5: new full-speed USB device number 31 using ehci-pci
// [80026.134650] usb 2-1.5: New USB device found, idVendor=045e, idProduct=028e, bcdDevice= 1.10
// [80026.134654] usb 2-1.5: New USB device strings: Mfr=1, Product=2, SerialNumber=3
// [80026.134656] usb 2-1.5: Product: Xbox360 For Windows
// [80026.134657] usb 2-1.5: Manufacturer: shanwan
// [80026.134657] usb 2-1.5: SerialNumber: Shanwan202107142050
// [80026.175608] input: Microsoft X-Box 360 pad as /devices/pci0000:00/0000:00:1d.0/usb2/2-1/2-1.5/2-1.5:1.0/input/input63
// [80026.175748] usbcore: registered new interface driver xpad
// dave@yeller:~$

type GStick struct {
	Device           *evdev.InputDevice
	InputID          evdev.InputID
	Version          [3]int
	Name             string
	PhysicalLocation string
	UniqueID         string
	AxesInfo         AxisInfoMap
	Properties       []evdev.EvProp

	preAxisState   AxisStateMap
	curAxisState   AxisStateMap
	preButtonState evdev.StateMap
	curButtonState evdev.StateMap

	//joy or game
	ButtonEvent ButtonEvent

	Buttons []Button
	Axes    []Axis

	spinning bool
}

func newStickG(device *evdev.InputDevice) (stg *GStick) {
	stg = &GStick{
		Device:         device,
		preAxisState:   make(AxisStateMap),
		curAxisState:   make(AxisStateMap),
		preButtonState: make(evdev.StateMap),
		curButtonState: make(evdev.StateMap),
	}
	return stg
}

func OpenStickG(path string) (*GStick, error) {
	device, err := evdev.Open(path)
	if err != nil {
		fmt.Println("OpenStickG", path, err)
		return nil, err
	}
	// defer device.NonBlock()

	stg := newStickG(device)

	stg.Version[0], stg.Version[1], stg.Version[2] = device.DriverVersion()
	stg.InputID, err = device.InputID()
	if err != nil {
		fmt.Println("InputID", err)
	}
	stg.Name, err = device.Name()
	if err != nil {
		fmt.Println("Name", err)
	}
	stg.PhysicalLocation, err = device.PhysicalLocation()
	if err != nil {
		fmt.Println("PhysicalLocation", err)
	}
	stg.UniqueID, err = device.UniqueID()
	if err != nil {
		fmt.Println("UniqueID", err)
	}

	for _, t := range device.CapableTypes() {
		var (
			state evdev.StateMap
			err   error
		)

		if t == evdev.EV_KEY {
			state, err = device.State(t)
			if err != nil {
				fmt.Println("device.State: ", err, t)
			} else {
				for k, v := range state {
					stg.preButtonState[k] = false
					stg.curButtonState[k] = v
				}
			}
		}

		if t == evdev.EV_ABS {
			stg.AxesInfo, err = device.AbsInfos()
			if err != nil {
				fmt.Println("device.AbsInfos(): ", err)
			}
		}
	}
	stg.Properties = device.Properties()
	return stg, nil
}

func (stg *GStick) ReadState() {
	state, err := stg.Device.State(evdev.EV_KEY)
	if err == nil {
		for k, v := range state {
			stg.preButtonState[k] = stg.curButtonState[k]
			stg.curButtonState[k] = v
		}
	}

	absInfos, err := stg.Device.AbsInfos()
	if err == nil {
		for k, v := range absInfos {
			stg.preAxisState[k] = stg.curAxisState[k]
			stg.curAxisState[k] = v.Value
		}
	}

}

func (stg *GStick) Dump() {
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

	fmt.Println("Axes")
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
	for _, p := range props {
		fmt.Printf("  Property type %d (%s)\n", p, evdev.PropName(p))
	}
}

func (stg *GStick) Close() {
	stg.spinning = false
	stg.Device.Close()
	fmt.Println("stg.Stop closed")
}
