package gdb

import (
	"fmt"

	"github.com/centretown/xray/access"
	"github.com/centretown/xray/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type GameData struct {
	Schema *Schema
	Access *access.Access
	dbx    *sqlx.DB
	Err    error
}

func NewGameData(driver, path string) *GameData {
	return &GameData{Schema: SchemaGame, Access: access.NewAccess(driver, path)}
}

func (gdb *GameData) Open() {
	gdb.dbx, gdb.Err = sqlx.Connect(gdb.Access.Driver, gdb.Access.Path)
}

func (gdb *GameData) Close() {
	gdb.Err = gdb.dbx.Close()
}

func (gdb *GameData) Create() {
	for _, sch := range gdb.Schema.Create {
		fmt.Println(sch)
		gdb.dbx.MustExec(sch)
	}
	tx := gdb.dbx.MustBegin()
	tx.NamedExec(gdb.Schema.InsertVersion, &gdb.Schema.Version)
	gdb.Err = tx.Commit()
}

func (gdb *GameData) InsertItem(item model.Recorder) {
	tx := gdb.dbx.MustBegin()
	defer func() {
		tx.Commit()
		if gdb.Err != nil {
			fmt.Println("InsertItem", gdb.Err)
		}
	}()

	_, gdb.Err = tx.NamedExec(gdb.Schema.InsertItem, item.GetRecord())
}

func (gdb *GameData) InsertLink(link *model.Link) {
	tx := gdb.dbx.MustBegin()
	defer func() {
		tx.Commit()
		if gdb.Err != nil {
			fmt.Println("InsertLink", gdb.Err)
		}
	}()

	_, gdb.Err = tx.NamedExec(gdb.Schema.InsertLink, &link)
}

func (gdb *GameData) GetItemID(id string) *model.Record {
	var uid uuid.UUID
	uid, gdb.Err = uuid.Parse((id))
	if gdb.Err != nil {
		return &model.Record{}
	}
	return gdb.GetItemUUID(uid)
}

func (gdb *GameData) GetItemUUID(uid uuid.UUID) *model.Record {
	major, minor := model.RecordID(uid)
	return gdb.GetItem(major, minor)
}

func (gdb *GameData) GetItemRecord(item model.Recorder) *model.Record {
	rec := item.GetRecord()
	return gdb.GetItem(rec.Major, rec.Minor)
}

func (gdb *GameData) GetItem(major, minor int64) *model.Record {
	item := &model.Record{}
	gdb.Err = gdb.dbx.Get(item, gdb.Schema.GetItem, major, minor)
	return item
}

func (gdb *GameData) GetLinks(rec *model.Record) (recs []*model.Record) {
	recs = make([]*model.Record, 0)
	var (
		rows  *sqlx.Rows
		link  model.Link
		links = make([]*model.Link, 0)
	)

	rows, gdb.Err = gdb.dbx.Queryx(gdb.Schema.GetLinks, rec.Major, rec.Minor)
	if gdb.Err != nil {
		return
	}

	for rows.Next() {
		gdb.Err = rows.StructScan(&link)
		if gdb.Err != nil {
			return
		}
		links = append(links, &link)
	}

	for _, l := range links {
		rec = gdb.GetItem(l.LinkedMajor, l.LinkedMinor)
		if gdb.Err != nil {
			return
		}
		recs = append(recs, rec)
	}
	return
}
