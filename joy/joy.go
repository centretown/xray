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
		fmt.Println("duration", duration)
	}

	run(duration, []func(){
		buttonPressed,
		buttonUp,
		buttonDown,
	})

	fmt.Println()
	fmt.Println("done!")
}

func run(duration time.Duration, tests []func()) {
	ch := make(chan int)
	for _, f := range tests {
		go test(ch, f)
		time.Sleep(duration)
	}
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

func buttonPressed() {
	button := tools.GetJoystickButtonPressed()
	up := tools.IsJoystickButtonUp(0, button)
	fmt.Printf("P[%s:%d]",
		tools.GetButtonName(0, button), tools.B2int(up))
}

func buttonUp() {
	button := 3
	up := tools.IsJoystickButtonUp(0, button)
	fmt.Printf("U[%d:%d]", button, tools.B2int(up))
}

func buttonDown() {
	button := 3
	down := tools.IsJoystickButtonDown(0, button)
	fmt.Printf("U[%d:%d]", button, tools.B2int(down))
}
