package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/centretown/xray/gizmo"
	"github.com/centretown/xray/model"
	"github.com/centretown/xray/runner/cmd"
	"github.com/google/uuid"
)

var (
	path     = "/home/dave/xray/game_01/"
	lifepath = "/home/dave/xray/life_01/"
)

func init() {
	cmd.Setup()
}

func main() {
	flag.Parse()

	fmt.Println(cmd.Path, cmd.MajorKey, cmd.MinorKey, cmd.Key)
	// 	Major: -8107658525041914367
	// 	Minor: -854626809563736956
	//  Key:   018eb516-87cf-7b8f-8494-6dd871c123f4

	// Major: -7390079466426692095
	// Minor: -1757195284917677394

	var (
		record = &model.Record{}
		err    error
		game   *gizmo.Game
		id     uuid.UUID
	)

	defer func() {
		if err != nil {
			flag.Usage()
			os.Exit(1)
		}
	}()

	if len(cmd.Path) > 0 {
		if cmd.Path == "life" {
			path = lifepath
			cmd.MajorKey = -7316587727477305855
			cmd.MinorKey = -4065651591770192760
		} else if cmd.Path == "goose" {
			cmd.MajorKey = -3209655326295290367
			cmd.MinorKey = 6019352295432975494
		}
	}

	if cmd.MajorKey != 0 {
		record.Major, record.Minor = cmd.MajorKey, cmd.MinorKey
	} else {
		id, err = uuid.Parse(cmd.Key)
		if err != nil {
			fmt.Println("Wrong key", err, cmd.Key)
			return
		}
		record.Major, record.Minor = model.RecordID(id)
	}

	game, err = gizmo.LoadGame(path, record)
	if err != nil {
		fmt.Println("Loading game:", err)
		j, n, k := game.Keys()
		fmt.Println(path, j, n, k)
		return
	}
	// game.Dump()

	game.Run()
}
