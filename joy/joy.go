package main

import (
	"flag"
	"fmt"
	"xray/jcmd"
)

func init() {
	jcmd.SetFlagsVars()
}

func main() {
	flag.Parse()
	cmds := jcmd.NewCmds()
	jcmd.RunJoyCmds(cmds)
	fmt.Println("done!")

}
