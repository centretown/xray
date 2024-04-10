package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/centretown/xray/cmdl"
	"github.com/centretown/xray/gizmo"
)

var (
	installBase = "/home/dave/xray/"
)

func init() {
	cmdl.Setup("test", "path", "version")
}

func main() {
	flag.Parse()
	cmd := cmdl.Cmdl
	var path string
	if cmd.Test {
		path = filepath.Clean(cmd.Path)
	} else {
		path = filepath.Join(installBase, cmd.Path)
	}
	fmt.Println(path, cmd.MajorKey, cmd.MinorKey, cmd.Key)
	var (
		err  error
		game *gizmo.Game
	)

	defer func() {
		if err != nil {
			flag.Usage()
			os.Exit(1)
		}
	}()

	game, err = gizmo.LoadGameKeys(path)
	if err != nil {
		fmt.Println("Loading game:", err)
		j, n, k := game.Keys()
		fmt.Println(path, j, n, k)
		return
	}
	// game.Dump()

	game.Run()
}

// if cmdl.MajorKey != 0 {
// 	record.Major, record.Minor = cmdl.MajorKey, cmdl.MinorKey
// } else {
// 	id, err = uuid.Parse(cmdl.Key)
// 	if err != nil {
// 		fmt.Println("Wrong key", err, cmdl.Key)
// 		return
// 	}
// 	record.Major, record.Minor = model.RecordID(id)
// }
