package gcmd

import (
	"flag"
	"fmt"
	"time"
	"xray/b2i"
	"xray/gpads"
	"xray/pad"
	"xray/tools"
)

var js pad.Pad = gpads.NewGPads()

// var js jstick.Jstick = joystickc.NewJoyStickC()

var pads tools.IntSlice
var keys tools.StringSlice
var buttons tools.IntSlice
var axes tools.IntSlice
var seconds tools.IntSlice

type GCmd struct {
	Title  string
	Cmd    func(*GCmd)
	Pad    int
	Button int
	Axis   int
	Delay  time.Duration
}

const (
	Last int = iota
	Up
	Down
	Press
	Release
	Move
	Keys
	Dump
)

var KeyList = []string{
	"last",
	"up",
	"down",
	"press",
	"release",
	"move",
	"keys",
	"dump",
}

var KeyUsage = []string{
	"indicate last button pressed",
	"indicate if selected button is up",
	"indicate if selected button is down",
	"indicate if selected button has been pressed",
	"indicate if selected button has been released",
	"indicate selected axis movement",
	"indicate any key pressed",
	"dump maps and value corrections",
}

var GCmds = map[string]*GCmd{
	KeyList[Last]:    {Cmd: LastButtonPressed, Title: KeyList[Last], Pad: 0, Button: 0, Axis: 0, Delay: 0},
	KeyList[Up]:      {Cmd: IsButtonUp, Title: KeyList[Up], Pad: 0, Button: 0, Axis: 0, Delay: 0},
	KeyList[Down]:    {Cmd: IsButtonDown, Title: KeyList[Down], Pad: 0, Button: 0, Axis: 0, Delay: 0},
	KeyList[Press]:   {Cmd: IsButtonPressed, Title: KeyList[Press], Pad: 0, Button: 0, Axis: 0, Delay: 0},
	KeyList[Release]: {Cmd: IsButtonReleased, Title: KeyList[Release], Pad: 0, Button: 0, Axis: 0, Delay: 0},
	KeyList[Move]:    {Cmd: GetAxisMovement, Title: KeyList[Move], Pad: 0, Button: 0, Axis: 0, Delay: 0},
	KeyList[Keys]:    {Cmd: TestKeys, Title: KeyList[Keys], Pad: 0, Button: 0, Axis: 0, Delay: 0},
	KeyList[Dump]:    {Cmd: DumpPad, Title: KeyList[Dump], Pad: 0, Button: 0, Axis: 0, Delay: 0},
}

const keysText = "one or more commands eg: -c down -c up -c last"
const joystickUsage = "one or more joysticks to test\neg: -c down -c up -j 0 -j 1 [runs:  up joystick 0, down joystick 1]"
const durationsUsage = "one or more durations in seconds\neg: -c down -c up -d 5 -d 6 [runs:  up 5s, down 6s]"
const buttonsUsage = "one or more buttons to test\neg: -c down -c up -c last -b 2 -b 5 [runs:  down button 2, up button 5, last button 5]"
const axesUsage = "one or more axes to test\neg: -c down -c up -c last -a 0 -a 1 -a 3 [runs:  down axis 0, up axis 1, last axis 3]"
const keyUsage = "test all keys\neg: -c keys [runs: keys]"

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
	flag.Var(&pads, "joystick", joystickUsage)
	flag.Var(&pads, "j", sameAs("joystick"))
	flag.Var(&buttons, "button", buttonsUsage)
	flag.Var(&buttons, "b", sameAs("button"))
	flag.Var(&axes, "axis", axesUsage)
	flag.Var(&axes, "a", sameAs("axis"))
	flag.Var(&seconds, "duration", durationsUsage)
	flag.Var(&seconds, "d", sameAs("duration"))
	flag.Var(&keys, "keys", keyUsage)
	flag.Var(&keys, "k", sameAs("keys"))

}

func RunJoyCmds(cmds []*GCmd) {
	ch := make(chan int)
	for _, c := range cmds {
		fmt.Println("START", c.Title)
		go c.RunCmd(ch)
		time.Sleep(c.Delay)
		ch <- 1
		fmt.Printf("DONE: %s\n\n", c.Title)
	}
}

func (cmd *GCmd) RunCmd(stopChan <-chan int) {
	var delay = time.Millisecond * 16 //ms

	js.BeginPad()
	showPad(cmd.Pad)
	showCmd(cmd)
	for {
		js.BeginPad()

		cmd.Cmd(cmd)

		select {
		case <-stopChan:
			return
		default:
			time.Sleep(delay)
		}
	}
}

func NewCmds() []*GCmd {

	cmds := make([]*GCmd, 0, len(keys))

	ensureOneReturnLast := func(is *tools.IntSlice, v int) int {
		if len(*is) < 1 {
			*is = []int{v}
		}
		return len(*is) - 1
	}
	jLast := ensureOneReturnLast(&pads, 0)
	bLast := ensureOneReturnLast(&buttons, 0)
	aLast := ensureOneReturnLast(&axes, 0)
	sLast := ensureOneReturnLast(&seconds, 0)
	padNext, axisNext, btnNext, secNext := 0, 0, 0, 0

	for _, key := range keys {
		pCmd, ok := GCmds[key]
		if !ok {
			fmt.Printf("invalid command %s\n", key)
			continue
		}

		cmd := *pCmd
		cmd.Pad = pads[padNext]
		cmd.Axis = axes[axisNext]
		cmd.Button = buttons[btnNext]
		cmd.Delay = time.Duration(seconds[secNext]) * time.Second

		padNext += b2i.Bool2int(padNext < jLast)
		axisNext += b2i.Bool2int(axisNext < aLast)
		btnNext += b2i.Bool2int(btnNext < bLast)
		secNext += b2i.Bool2int(secNext < sLast)
		cmds = append(cmds, &cmd)
		showCmd(&cmd)
	}

	return cmds
}

func showCmd(c *GCmd) {
	fmt.Printf("command: %s, joystick: %d axis: %d, button = %d, duration = %v\n",
		c.Title, c.Pad, c.Axis, c.Button, c.Delay)

}

func showPad(pad int) {
	fmt.Printf("%s, available:%v, axes:%d, buttons:%d\n",
		js.GetPadName(pad),
		js.IsPadAvailable(pad),
		js.GetPadAxisCount(pad),
		js.GetPadButtonCount(pad),
	)
}
