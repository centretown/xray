package gcmd

import (
	"flag"
	"fmt"
	"time"
	"xray/gpads"
	"xray/pad"
	"xray/tools"
)

var js pad.Pad = gpads.NewGPads()

// var js jstick.Jstick = joystickc.NewJoyStickC()

var keys tools.StringSlice
var joysticks tools.IntSlice
var buttons tools.IntSlice
var axes tools.IntSlice
var seconds tools.IntSlice

type JoyCmd struct {
	Title    string
	Cmd      func(*JoyCmd)
	Joystick int
	Button   int
	Axis     int
	Delay    time.Duration
}

const (
	Last int = iota
	Up
	Down
	Press
	Release
	Move
	Dump
)

var KeyList = []string{
	"last",
	"up",
	"down",
	"press",
	"release",
	"move",
	"dump",
}

var KeyUsage = []string{
	"indicate last button pressed",
	"indicate if selected button is up",
	"indicate if selected button is down",
	"indicate if selected button has been pressed",
	"indicate if selected button has been released",
	"indicate selected axis movement",
	"dump maps and value corrections",
}

var JoyCmds = map[string]*JoyCmd{
	KeyList[Last]:    {Cmd: LastButtonPressed, Title: KeyList[Last], Joystick: 0, Button: 0, Axis: 0, Delay: 0},
	KeyList[Up]:      {Cmd: IsButtonUp, Title: KeyList[Up], Joystick: 0, Button: 0, Axis: 0, Delay: 0},
	KeyList[Down]:    {Cmd: IsButtonDown, Title: KeyList[Down], Joystick: 0, Button: 0, Axis: 0, Delay: 0},
	KeyList[Press]:   {Cmd: IsButtonPressed, Title: KeyList[Press], Joystick: 0, Button: 0, Axis: 0, Delay: 0},
	KeyList[Release]: {Cmd: IsButtonReleased, Title: KeyList[Release], Joystick: 0, Button: 0, Axis: 0, Delay: 0},
	KeyList[Move]:    {Cmd: GetAxisMovement, Title: KeyList[Move], Joystick: 0, Button: 0, Axis: 0, Delay: 0},
	KeyList[Dump]:    {Cmd: GetDumpJoystick, Title: KeyList[Dump], Joystick: 0, Button: 0, Axis: 0, Delay: 0},
}

const keysText = "one or more commands eg: -c down -c up -c last"
const joystickUsage = "one or more joysticks to test\neg: -c down -c up -j 0 -j 1 [runs:  up joystick 0, down joystick 1]"
const durationsUsage = "one or more durations in seconds\neg: -c down -c up -d 5 -d 6 [runs:  up 5s, down 6s]"
const buttonsUsage = "one or more buttons to test\neg: -c down -c up -c last -b 2 -b 5 [runs:  down button 2, up button 5, last button 5]"
const axesUsage = "one or more axes to test\neg: -c down -c up -c last -a 0 -a 1 -a 3 [runs:  down axis 0, up axis 1, last axis 3]"

var commandsUsage string

func SetFlagsVars() {
	if len(KeyList) != len(KeyUsage) {
		panic("keys and usage don't match")
	}
	keyAndUsage := func(k []string, u []string) string {
		ss := ""
		for i := range k {
			ss += fmt.Sprintf("\n%12s - %s", k[i], u[i])
		}
		return ss
	}
	sameAs := func(s string) string {
		return "same as " + s
	}

	commandsUsage = fmt.Sprintf("%s %s", keysText, keyAndUsage(KeyList, KeyUsage))
	flag.Var(&keys, "command", commandsUsage)
	flag.Var(&keys, "c", sameAs("command"))
	flag.Var(&joysticks, "joystick", joystickUsage)
	flag.Var(&joysticks, "j", sameAs("joystick"))
	flag.Var(&buttons, "button", buttonsUsage)
	flag.Var(&buttons, "b", sameAs("button"))
	flag.Var(&axes, "axis", axesUsage)
	flag.Var(&axes, "a", sameAs("axis"))
	flag.Var(&seconds, "duration", durationsUsage)
	flag.Var(&seconds, "d", sameAs("duration"))

}

func RunJoyCmds(cmds []*JoyCmd) {
	ch := make(chan int)
	for _, c := range cmds {
		fmt.Println("START", c.Title)
		go c.RunCmd(ch)
		time.Sleep(c.Delay)
		ch <- 1
		fmt.Println("DONE", c.Title)
		fmt.Println()
	}
}

func (cmd *JoyCmd) RunCmd(ch <-chan int) {
	const delay = 16 //ms

	js.BeginPad()
	showStick(cmd.Joystick)
	showCmd(cmd)
	for {
		js.BeginPad()

		cmd.Cmd(cmd)

		select {
		case <-ch:
			return
		default:
			time.Sleep(time.Millisecond * delay)
		}
	}
}

func LastButtonPressed(cmd *JoyCmd) {
	button := js.GetPadButtonPressed()
	up := js.IsPadButtonDown(cmd.Joystick, button)
	fmt.Printf("[%4d:%4d]\r", button, tools.Bool2int(up))
}

func IsButtonUp(cmd *JoyCmd) {
	up := js.IsPadButtonUp(cmd.Joystick, cmd.Button)
	if up {
		fmt.Printf("[%d:%d]\r", cmd.Button, tools.Bool2int(!up))
	}
}

func IsButtonDown(cmd *JoyCmd) {
	down := js.IsPadButtonDown(cmd.Joystick, cmd.Button)
	if down {
		fmt.Printf("[%d:%d]\r", cmd.Button, tools.Bool2int(down))
	}
}

func IsButtonReleased(cmd *JoyCmd) {
	released := js.IsPadButtonReleased(cmd.Joystick, cmd.Button)
	if released {
		fmt.Printf("[%d:%d]\r", cmd.Button, tools.Bool2int(released))
	}
}

func IsButtonPressed(cmd *JoyCmd) {
	pressed := js.IsPadButtonPressed(cmd.Joystick, cmd.Button)
	if pressed {
		fmt.Printf("[%d:%d]\r", cmd.Button, tools.Bool2int(pressed))
	}
}

func GetAxisValues(cmd *JoyCmd) {
	count := js.GetPadAxisCount(cmd.Joystick)
	fmt.Print("axes:  ")
	for i := range count {
		value := js.GetPadAxisValue(cmd.Joystick, i)
		fmt.Printf("[%d:%6d] ", i, value)
	}
	fmt.Print("\r")
}

func GetAxisMovement(cmd *JoyCmd) {
	count := js.GetPadAxisCount(cmd.Joystick)
	fmt.Print("axes:  ")
	for i := range count {
		move := js.GetPadAxisMovement(cmd.Joystick, i)
		fmt.Printf("[%d:%6.0f] ", i, move)
	}
	fmt.Print("\r")
}

func GetDumpJoystick(cmd *JoyCmd) {
	js.DumpPad()
}

func NewCmds() []*JoyCmd {

	cmds := make([]*JoyCmd, 0, len(keys))

	ensureOneReturnLast := func(is *tools.IntSlice, v int) int {
		if len(*is) < 1 {
			*is = []int{v}
		}
		return len(*is) - 1
	}
	jLast := ensureOneReturnLast(&joysticks, 0)
	bLast := ensureOneReturnLast(&buttons, 0)
	aLast := ensureOneReturnLast(&axes, 0)
	sLast := ensureOneReturnLast(&seconds, 0)
	joyNext, axisNext, btnNext, secNext := 0, 0, 0, 0

	for _, key := range keys {
		pCmd, ok := JoyCmds[key]
		if !ok {
			fmt.Printf("invalid command %s\n", key)
			continue
		}

		cmd := *pCmd
		cmd.Joystick = joysticks[joyNext]
		cmd.Axis = axes[axisNext]
		cmd.Button = buttons[btnNext]
		cmd.Delay = time.Duration(seconds[secNext]) * time.Second

		joyNext += tools.Bool2int(joyNext < jLast)
		axisNext += tools.Bool2int(axisNext < aLast)
		btnNext += tools.Bool2int(btnNext < bLast)
		secNext += tools.Bool2int(secNext < sLast)
		cmds = append(cmds, &cmd)
		showCmd(&cmd)
	}

	return cmds
}

func showCmd(c *JoyCmd) {
	fmt.Printf("command: %s, joystick: %d axis: %d, button = %d, duration = %v\n",
		c.Title, c.Joystick, c.Axis, c.Button, c.Delay)

}

func showStick(JoyStick int) {
	fmt.Printf("%s, available:%v, axes:%d, buttons:%d\n",
		js.GetPadName(JoyStick),
		js.IsPadAvailable(JoyStick),
		js.GetPadAxisCount(JoyStick),
		js.GetPadButtonCount(JoyStick),
	)
}
