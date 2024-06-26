package main

import (
	"github.com/centretown/xray/numbers"

	"github.com/centretown/gpads/gpads"
	"github.com/holoplot/go-evdev"
)

func Example(gpad *gpads.GPad) {
	var (
		isDown, wasDown bool
		button          int
		code            evdev.EvCode
	)

	state, err := gpad.Device.State(evdev.EV_KEY)
	if err == nil {
		for button, code = range gpad.ButtonCodes {
			// previous state
			wasDown = gpad.ButtonState[button]
			// current state
			isDown = state[code]

			// flag gets set only button was up and is now down
			// if b2i returns 1 and is shifted left button positions
			// if b2i returns 0 nothing happens
			gpad.PressedOnce |= numbers.As[uint64](!wasDown && isDown) << button
			// flag gets set only button was down and is now up
			gpad.ReleasedOnce |= numbers.As[uint64](wasDown && !isDown) << button
			gpad.ButtonState[button] = isDown
		}
	}

}
