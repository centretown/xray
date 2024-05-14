package gizzmo

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/centretown/xray/access"
	"github.com/centretown/xray/class"
	"github.com/centretown/xray/gizzmodb"
	"github.com/centretown/xray/gizzmodb/model"
)

// LoadGame reconstructs a game from the database.
// A game is selected from the top entry of the key stack file.
// The stack and database are created or updated at the end
// of the build process.
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
		data.Err = fmt.Errorf("%s with keys%d:%d is not game",
			record.Class, record.Major, record.Minor)
		return
	}

	gameRecord := game.GetRecord()
	records := data.LoadLinks(gameRecord)
	if data.HasErrors() {
		return
	}

	game.data = data
	link(game.data, game, records)
	if data.HasErrors() {
		return
	}

	// game.Content.options = load

	game.Run()
	return
}

func link(data *gizzmodb.Data, parent model.Parent, records []*model.Record) {
	var (
		recorder model.Recorder
	)

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

			// dig a little deeper
			link(data, p, rs)
		}
	}
}

// makeLink constructs concrete structs from a database class record
func makeLink(record *model.Record) (recorder model.Recorder) {
	cls := class.Class(record.Classn)
	switch cls {
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

	log.Fatal(fmt.Errorf("unknown Class %d(%s)", cls, cls))
	return nil
}
