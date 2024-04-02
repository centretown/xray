package dbio

import (
	"fmt"

	"github.com/centretown/xray/access"
	"github.com/centretown/xray/model"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type GameData struct {
	Schema *Schema
	Access *access.Access
	DB     *sqlx.DB
	Err    error
}

func NewGameData(driver, path string) *GameData {
	return &GameData{Schema: SchemaGame, Access: access.NewAccess(driver, path)}
}

func (gd *GameData) Open() {
	gd.DB, gd.Err = sqlx.Connect(gd.Access.Driver, gd.Access.Path)
}

func (gd *GameData) Close() {
	gd.Err = gd.DB.Close()
}

func (gd *GameData) Create() {

	for _, sch := range gd.Schema.Create {
		fmt.Println(sch)
		gd.DB.MustExec(sch)
	}
	tx := gd.DB.MustBegin()
	fmt.Println(gd.Schema.Version)
	tx.NamedExec(gd.Schema.InsertVersion, &gd.Schema.Version)
	gd.Err = tx.Commit()
}

func (gd *GameData) InsertItem(item *model.Item) {
	tx := gd.DB.MustBegin()
	fmt.Println(item)
	tx.NamedExec(gd.Schema.InsertItem, item)
	gd.Err = tx.Commit()
}
