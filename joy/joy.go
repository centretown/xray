package main

import (
	"flag"
	"fmt"
	"xray/stick"
)

func init() {
	stick.SetFlagsVars()
}

func main() {
	flag.Parse()
	cmds := stick.NewCmds()
	stick.RunJoyCmds(cmds)
	fmt.Println("done!")

}
