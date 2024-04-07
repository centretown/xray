package main

import (
	"github.com/centretown/xray/gizmo"
	"github.com/centretown/xray/model"
)

var (
	path = "/home/dave/xray/game_01/"
)

func main() {
	record := &model.Record{
		Major: -8107658525041914367,
		Minor: -854626809563736956}
	game := gizmo.LoadGame(path, record)
	game.Run()
}
