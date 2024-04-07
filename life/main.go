package main

import "github.com/centretown/xray/gizmo"

func main() {
	game := gizmo.NewGameSetup(800, 450, 20)
	game.FramesCounter = 0
}
