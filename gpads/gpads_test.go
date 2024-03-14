package gpads

import (
	"fmt"
	"testing"
	"time"

	"github.com/holoplot/go-evdev"
)

func TestJoy(t *testing.T) {
	j := NewGPads()
	j.BeginPad()

	if !j.IsPadAvailable(0) {
		t.Log("NOTHING TO TEST!")
		return
	}

	j.DumpPad()

	// fmt.Println("start and stop after 5 seconds")
	// // j.Sticks[0].Start()
	// j.Sticks[0].Dump()
	// ch := make(chan int)
	// go dumpEvent(j.Sticks[0].Device, ch)
	// time.Sleep(45 * time.Second)
	// ch <- 1
	// time.Sleep(time.Millisecond * 3)
	// j.Sticks[0].Close()
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
			if !ok {
				prev[code] = value
			} else if v != value {
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

// func TestKeyChange(t *testing.T) {
// 	j := NewGPads()
// 	j.BeginPad()
// 	fmt.Println("Count", j.GetStickCount())
// 	fmt.Printf("Selecting %s\n", j.GetPadName(0))
// 	time.Sleep(time.Second)
// 	count := j.GetStickCount()
// 	x := 0
// 	for {
// 		j.BeginPad()
// 		for i := range count {
// 			stg := j.Pads[i]
// 			if j.IsPadButtonDown(i, 0) {
// 				code := stg.ButtonBase + 0
// 				fmt.Println(evdev.KEYToString[code], "DOWN", i, x)
// 				x++
// 			}
// 		}
// 		time.Sleep(time.Millisecond << 4)
// 	}

// }

// func TestPressed(t *testing.T) {
// 	j := NewGPads()
// 	j.BeginPad()
// 	fmt.Println("Count", j.GetStickCount())
// 	fmt.Printf("Selecting %s\n", j.GetPadName(0))
// 	time.Sleep(time.Second)
// 	count := j.GetStickCount()
// 	x := 0
// 	for {
// 		j.BeginPad()
// 		for i := range count {
// 			stg := j.Pads[i]
// 			if j.IsPadButtonDown(i, 0) {
// 				code := stg.ButtonBase + 0
// 				fmt.Println(evdev.KEYToString[code], "DOWN", i, x)
// 				x++
// 			}
// 		}
// 		time.Sleep(time.Millisecond << 4)
// 	}

// }
