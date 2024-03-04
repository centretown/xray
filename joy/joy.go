package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
	"xray/tools"
)

const (
	Last int = iota
	Up
	Down
)

var commandList = []string{
	"last",
	"up",
	"down",
}

var commandsText = "specify commands"
var commandsUsage string
var durationsUsage = "specify durations in seconds eg: -d 5 -d 6 (1st cmd 5s, 2nd... 6s)"
var buttonsUsage = "specify buttons to test eg: -b 2 -b 5 (1st cmd button 2, 2nd... button 5)"
var axesUsage = "specify axes to test eg: -a 0 -a 1 -a 3 (1st cmd axis 0, 2nd axis 1, 3rd... axis 3)"

var commands tools.StringSlice
var seconds tools.IntSlice
var buttons tools.IntSlice
var axes tools.IntSlice

type JoyCmd struct {
	f        func(*JoyCmd)
	title    string
	button   int
	axis     int
	duration time.Duration
}

var JoyCmds = map[string]*JoyCmd{
	commandList[Last]: {f: buttonPressed, title: commandList[Last], button: 0, axis: 0, duration: 0},
	commandList[Up]:   {f: buttonUp, title: commandList[Up], button: 0, axis: 0, duration: 0},
	commandList[Down]: {f: buttonDown, title: commandList[Down], button: 0, axis: 0, duration: 0},
}

func init() {
	commandsUsage = fmt.Sprintf("%s [%s]", commandsText, strings.Join(commandList, ", "))
	flag.Var(&commands, "commands", commandsUsage)
	flag.Var(&commands, "c", commandsUsage)

	flag.Var(&seconds, "duration", durationsUsage)
	flag.Var(&seconds, "d", durationsUsage)

	flag.Var(&buttons, "button", buttonsUsage)
	flag.Var(&buttons, "b", buttonsUsage)

	flag.Var(&axes, "axis", axesUsage)
	flag.Var(&axes, "a", axesUsage)
}

func main() {
	flag.Parse()
	cmds := createCmds()
	run(cmds)
	fmt.Println("done!")

}

func run(cmds []*JoyCmd) {
	ch := make(chan int)
	for _, c := range cmds {
		fmt.Println("START", c.title)
		go runCmd(ch, c)
		time.Sleep(c.duration)
		ch <- 1
		fmt.Println("DONE", c.title)
		fmt.Println()
	}
}

func runCmd(ch <-chan int, cmd *JoyCmd) {
	const delay = 16 //ms

	tools.BeginJoystick()
	fmt.Printf("%s, available:%v, axes:%d\n",
		tools.GetJoystickName(0),
		tools.IsJoystickAvailable(0),
		tools.GetJoystickAxisCount(0),
	)
	for {
		tools.BeginJoystick()
		cmd.f(cmd)

		select {
		case <-ch:
			return
		default:
			time.Sleep(time.Millisecond * delay)
		}
	}
}

func buttonPressed(cmd *JoyCmd) {
	button := tools.GetJoystickButtonPressed()
	up := tools.IsJoystickButtonUp(0, button)
	fmt.Printf("[%d:%d]",
		button, tools.B2int(up))
}

func buttonUp(cmd *JoyCmd) {
	button := cmd.button
	up := tools.IsJoystickButtonUp(0, button)
	if up {
		fmt.Printf("U[%d:%d]", button, tools.B2int(up))
	}
}

func buttonDown(cmd *JoyCmd) {
	button := cmd.button
	down := tools.IsJoystickButtonDown(0, button)
	if down {
		fmt.Printf("U[%d:%d]", button, tools.B2int(down))
	}
}

func createCmds() []*JoyCmd {

	cmds := make([]*JoyCmd, 0, len(commands))
	if len(buttons) < 1 {
		buttons = []int{0}
	}
	if len(axes) < 1 {
		axes = []int{0}
	}
	if len(seconds) < 1 {
		seconds = []int{0}
	}

	bLen := len(buttons) - 1
	aLen := len(axes) - 1
	sLen := len(seconds) - 1
	ai, bi, si := 0, 0, 0

	for _, k := range commands {
		c, ok := JoyCmds[k]
		if !ok {
			fmt.Printf("invalid command %s\n", k)
			continue
		}

		c.axis = axes[ai]
		c.button = buttons[bi]
		c.duration = time.Duration(seconds[si]) * time.Second

		ai += tools.B2int(ai < aLen)
		bi += tools.B2int(bi < bLen)
		si += tools.B2int(si < sLen)
		cmds = append(cmds, c)

		fmt.Printf("command: %s, axis: %d, button = %d, duration = %v\n",
			c.title, c.axis, c.button, c.duration)
	}

	return cmds
}
