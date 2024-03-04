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
	// tools.BeginJoystick()
	// fd := tools.IsJoystickAvailable(0)
	// fmt.Println(tools.GetJoystickName(0))
	// fmt.Println(tools.GetJoystickButtonPressed())

	ch := make(chan int)
	f := func(ch chan int) {
		tools.BeginJoystick()
		fmt.Println(tools.GetJoystickName(0))
		fmt.Println("available", tools.IsJoystickAvailable(0))
		for {
			tools.BeginJoystick()
			fmt.Print(tools.GetJoystickButtonPressed(), ":")

			select {
			case <-ch:
				return
			default:
				time.Sleep(time.Millisecond * 16)
			}
		}
	}

	go f(ch)
	time.Sleep(duration)
	fmt.Println()
	fmt.Println("done!")
}
