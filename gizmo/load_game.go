package gizmo

import (
	"log"
	"path/filepath"

	"github.com/centretown/xray/access"
	"github.com/centretown/xray/gizmodb"
	"github.com/centretown/xray/model"
)

func LoadGame() (err error) {
	path := "."
	gameKeys, err := access.LoadGameKey(filepath.Join(path, "game_keys.yaml"))
	if err != nil {
		log.Fatalln(err)
		return
	}

	log.Println("LoadGameKey", gameKeys.Minor, gameKeys.Major)

	data := gizmodb.NewGameData("sqlite3", filepath.Join(path, "xray_game.db"))
	defer func() {
		if data.HasErrors() {
			err = data.Err
			log.Fatalln(data.Err)
			return
		}
		data.Close()
	}()

	var (
		record   model.Record
		recorder model.Recorder
		ok       bool
		game     *Game
	)

	record.Major = gameKeys.Major
	record.Minor = gameKeys.Minor

	data.GetRecord(&record)
	if data.HasErrors() {
		return
	}

	recorder = MakeLink(&record)
	game, ok = recorder.(*Game)
	if !ok {
		log.Fatal()
	}

	gameRecord := game.GetRecord()
	records := data.LoadLinks(gameRecord)
	if data.HasErrors() {
		return
	}

	game.data = data
	link(game.data, game, records)

	game.Run()
	return
}

func link(data *gizmodb.Data, parent model.Parent, records []*model.Record) {
	var (
		recorder model.Recorder
	)

	defer func() {
		if data.HasErrors() {
			log.Fatal(data.Err)
		}
	}()

	for _, record := range records {

		data.GetRecord(record)
		if data.HasErrors() {
			return
		}

		recorder = MakeLink(record)

		parent.LinkChild(recorder)

		p, ok := recorder.(model.Parent)
		if ok {
			rs := data.LoadLinks(p.GetRecord())
			if data.HasErrors() {
				return
			}
			link(data, p, rs)
		}
	}
}
