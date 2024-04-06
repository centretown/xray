package main

import (
	"github.com/centretown/xray/gizmo"
	"github.com/centretown/xray/model"
)

var (
	dir    = "/home/dave/xray/test/"
	dbfile = dir + "db/xray_game.db"
)

func main() {
	record := &model.Record{
		Major: -8107658525041914367,
		Minor: -854626809563736956}

	game := gizmo.LoadGame("sqlite3", dbfile, record)
	game.Dump()
	game.Run()
}
