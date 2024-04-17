package gizzmo

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/centretown/xray/access"
	"github.com/centretown/xray/gizzmo/class"
	"github.com/centretown/xray/gizzmodb"
	"github.com/centretown/xray/gizzmodb/model"
)

func LoadGame() (err error) {
	path := "."
	gameKeys, err := access.LoadGameKey(filepath.Join(path, "game_keys.yaml"))
	if err != nil {
		log.Fatalln(err)
		return
	}

	log.Println("LoadGameKey", gameKeys.Minor, gameKeys.Major)

	data := gizzmodb.NewGameData("sqlite3", filepath.Join(path, "xray_game.db"))
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

	recorder = makeLink(&record)
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

func link(data *gizzmodb.Data, parent model.Parent, records []*model.Record) {
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

		recorder = makeLink(record)

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

func makeLink(record *model.Record) (recorder model.Recorder) {
	fmt.Println("MakeLink", *record)
	cat := class.Class(record.Classn)
	switch cat {
	case class.Game:
		return NewGameFromRecord(record)
	case class.Texture:
		return NewTextureFromRecord(record)
	case class.Ellipse:
		return NewEllipseFromRecord(record)
	case class.CellsOrg:
		return NewCellsOrgFromRecord(record)
	case class.Tracker:
		return NewTrackerFromRecord(record)
	case class.LifeMover:
		return NewLifeMoverFromRecord(record)
	case class.LifeGrid:
		return NewLifeGridFromRecord(record)
	}

	err := fmt.Errorf("unknown Class %d(%s)", cat, cat)
	log.Fatal(err)
	return nil
}
