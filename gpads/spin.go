package gpads

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
