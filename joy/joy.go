package main

import (
	"flag"
	"fmt"
	"time"
	"xray/tools"
)

func main() {
	duration := 10 * time.Second
	flag.Parse()
	if len(flag.Args()) > 0 {
		fmt.Println(flag.Arg(0))
		var num time.Duration
		n, err := fmt.Sscanf(flag.Arg(0), "%d", &num)
		if err == nil && n == 1 {
			duration = num * time.Second
		}
		fmt.Println("num", num, duration)
	}

	buttonPressed := func() {
		button := tools.GetJoystickButtonPressed()
		down := tools.IsJoystickButtonUp(0, button)
		fmt.Printf("[%s:%d]",
			tools.GetButtonName(0, button), tools.B2int(down))
	}

	ch := make(chan int)
	go test(ch, buttonPressed)
	time.Sleep(duration)
	fmt.Println()
	fmt.Println("done!")
}

func test(ch <-chan int, f func()) {
	tools.BeginJoystick()
	fmt.Printf("%s, available:%v, axes:%d\n",
		tools.GetJoystickName(0),
		tools.IsJoystickAvailable(0),
		tools.GetJoystickAxisCount(0),
	)
	for {
		tools.BeginJoystick()
		f()

		select {
		case <-ch:
			return
		default:
			time.Sleep(time.Millisecond * 16)
		}
	}
}
