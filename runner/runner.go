package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/centretown/xray/cmdl"
	"github.com/centretown/xray/gizmo"
)

var (
	path     = "/home/dave/xray/game_01/"
	lifepath = "/home/dave/xray/life_01/"
	gameName = "game"
	lifeName = "life"
)

func init() {
	cmdl.Setup()
}

func main() {
	flag.Parse()

	fmt.Println(cmdl.Path, cmdl.MajorKey, cmdl.MinorKey, cmdl.Key)
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

	if len(cmdl.Path) > 0 {
		if cmdl.Path == "life" {
			path = lifepath
		}
	}

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
