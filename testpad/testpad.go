package main

import (
	"flag"
	"fmt"
	"xray/gcmd"
)

func init() {
	gcmd.SetFlagsVars()
}

func main() {
	flag.Parse()
	cmds := gcmd.NewCmds()
	gcmd.RunJoyCmds(cmds)
	fmt.Println("done!")

}
