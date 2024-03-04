package main

import "xray/tools"

func main() {
	runr := tools.NewRunner(1280, 720, 60)
	runr.Run3d()
}
