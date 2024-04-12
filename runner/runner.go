package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/centretown/xray/gizmo"
)

var (
	installBase = "/home/dave/xray/"
)

func main() {
	flag.Parse()
	// cmd := cmdl.Cmdl
	var (
		path, dir string
		err       error
	)

	selection := flag.Arg(0)
	if len(selection) == 0 {
		fmt.Println("Enter the name of the game.")
		os.Exit(1)
	}

	path = filepath.Join(installBase, selection)

	dir, err = filepath.Abs(installBase)
	if err != nil {
		log.Fatalln(err)
	}

	err = os.Chdir(dir)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(path) //, cmd.MajorKey, cmd.MinorKey, cmd.Key)

	var (
		game *gizmo.Game
	)

	game, err = gizmo.LoadGameKey(path)
	if err != nil {
		log.Println("Loading game:", err)
		j, n, k := game.Keys()
		log.Println(path, j, n, k)
		os.Exit(1)
	}

	game.Run()
}
