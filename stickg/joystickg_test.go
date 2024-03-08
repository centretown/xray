package stickg

import (
	"fmt"
	"testing"
	"time"

	"github.com/holoplot/go-evdev"
)

func TestJoy(t *testing.T) {
	j := NewJoyStickG()
	j.BeginJoystick()

	if !j.IsJoystickAvailable(0) {
		t.Log("NOTHING TO TEST!")
		return
	}

	fmt.Println("start and stop after 5 seconds")
	// j.Sticks[0].Start()
	j.Sticks[0].Dump()
	ch := make(chan int)
	go dumpEvent(j.Sticks[0].Device, ch)
	time.Sleep(45 * time.Second)
	ch <- 1
	time.Sleep(time.Millisecond * 3)
	j.Sticks[0].Close()
}

func dumpEvent(device *evdev.InputDevice, ch chan int) {
	//[EV_KEY], code: 0x122 [BTN_THUMB
	var (
		eventType evdev.EvType   = evdev.EV_KEY
		prev      evdev.StateMap = make(evdev.StateMap)
		state     evdev.StateMap
		err       error
	)

	state, err = device.State(eventType)
	if err != nil {
		fmt.Println(err)
		return
	}

	for code, value := range state {
		prev[code] = value
	}

	stateChanged := func(s evdev.StateMap) (change bool) {
		for code, value := range s {
			v, ok := prev[code]
			if ok && v != value {
				prev[code] = value
				change = true
			}
		}
		return
	}

	for {
		state, err = device.State(eventType)
		if err != nil {
			fmt.Println(err)
		} else if stateChanged(state) {
			for code, value := range state {
				fmt.Printf("    Event code %d (%s) state %v\n",
					code,
					evdev.CodeName(eventType, code), value)
			}
		}

		select {
		case <-ch:
			return
		default:
			time.Sleep(time.Millisecond)
		}
	}
}
