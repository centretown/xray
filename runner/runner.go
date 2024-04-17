package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/centretown/xray/gizzmo"
)

var (
	installBase = "/home/dave/xray/"
)

func main() {
	flag.Parse()
	if len(flag.Args()) < 2 {
		log.Println("Enter a title. (eg: life02)")
		return
	}

	var (
		selection    = flag.Arg(1)
		runDirectory string
		err          error
	)

	defer func() {
		if err != nil {
			log.Printf("Unable run %s. Cause: %v\n", selection, err)
			os.Exit(1)
		}
	}()

	runDirectory, err = filepath.Abs(filepath.Join(installBase, selection))
	if err != nil {
		return
	}

	err = os.Chdir(runDirectory)
	if err != nil {
		return
	}

	err = gizzmo.LoadGame()
	if err != nil {
		return
	}
}
