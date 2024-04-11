package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/centretown/xray/cmdl"
	"github.com/centretown/xray/gizmo"
)

var (
	installBase = "/home/dave/xray/"
)

func init() {
	cmdl.Setup("path")
}

func main() {
	flag.Parse()
	cmd := cmdl.Cmdl
	var (
		path, dir string
		err       error
	)

	path = filepath.Join(installBase, cmd.Path)
	dir, err = filepath.Abs(installBase)
	if err != nil {
		log.Fatalln(err)
	}

	err = os.Chdir(dir)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(path, cmd.MajorKey, cmd.MinorKey, cmd.Key)

	var (
		game *gizmo.Game
	)

	defer func() {
		if err != nil {
			flag.Usage()
			os.Exit(1)
		}
	}()

	game, err = gizmo.LoadGameKey(path)
	if err != nil {
		log.Println("Loading game:", err)
		j, n, k := game.Keys()
		log.Println(path, j, n, k)
		return
	}

	game.Run()
}
