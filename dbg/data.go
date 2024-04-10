package dbg

import (
	"fmt"

	"github.com/centretown/xray/access"
	"github.com/centretown/xray/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type Data struct {
	Schema *Schema
	Keys   *access.DataKeys
	dbx    *sqlx.DB
	Err    error
}

func NewGameData(driver, path string) *Data {
	return &Data{Schema: SchemaGame,
		Keys: access.NewDataKeys(driver, path)}
}

func (data *Data) Open() *Data {
	data.dbx, data.Err = sqlx.Connect(data.Keys.Driver, data.Keys.Path)
	return data
}

func (data *Data) HasErrors() bool {
	return data.Err != nil
}

func (data *Data) Close() {
	data.Err = data.dbx.Close()
}

func (data *Data) Create(game *model.Record, version *model.Version) {
	version.ItemMajor = game.Major
	version.ItemMinor = game.Minor
	data.GetVersion(version)

	if data.HasErrors() {
		for _, sch := range data.Schema.Create {
			fmt.Println(sch)
			data.dbx.MustExec(sch)
		}
		tx := data.dbx.MustBegin()
		tx.NamedExec(data.Schema.InsertVersion, &data.Schema.Version)
		data.Err = tx.Commit()
	}
}

func (data *Data) InsertItems(items ...model.Recorder) {
	tx := data.dbx.MustBegin()
	defer func() {
		tx.Commit()
		if data.HasErrors() {
			fmt.Println("InsertItem", data.Err)
		}
	}()
	for _, item := range items {
		item.GetRecord().UpdateContent(item.GetItem())
		if !data.HasErrors() {
			_, data.Err = tx.NamedExec(data.Schema.InsertItem, item.GetRecord())
		}
		if data.HasErrors() {
			return
		}
	}
}

func (data *Data) InsertLinks(links ...*model.Link) {
	tx := data.dbx.MustBegin()
	defer func() {
		tx.Commit()
		if data.HasErrors() {
			fmt.Println("InsertLink", data.Err)
		}
	}()
	for _, link := range links {
		_, data.Err = tx.NamedExec(data.Schema.InsertLink, link)
		if data.HasErrors() {
			return
		}
	}
}

func (data *Data) GetItemID(id string) *model.Record {
	var uid uuid.UUID
	uid, data.Err = uuid.Parse((id))
	if data.HasErrors() {
		return &model.Record{}
	}
	return data.GetItemUUID(uid)
}

func (data *Data) GetItemUUID(uid uuid.UUID) *model.Record {
	major, minor := model.RecordID(uid)
	return data.GetItem(major, minor)
}

func (data *Data) GetItemRecord(item model.Recorder) *model.Record {
	rec := item.GetRecord()
	return data.GetItem(rec.Major, rec.Minor)
}

func (data *Data) GetItem(major, minor int64) *model.Record {
	item := &model.Record{}
	data.Err = data.dbx.Get(item, data.Schema.GetItem, major, minor)
	return item
}

func (data *Data) GetVersion(version *model.Version) *model.Version {
	data.Err = data.dbx.Get(version, data.Schema.GetVersion,
		version.ItemMajor, version.ItemMinor, version.Major, version.Minor)
	return version
}

func (data *Data) GetLinks(rec *model.Record) (recs []*model.Record) {
	recs = make([]*model.Record, 0)
	var (
		rows  *sqlx.Rows
		links = make([]*model.Link, 0)
	)

	rows, data.Err = data.dbx.Queryx(data.Schema.GetLinks, rec.Major, rec.Minor)
	if data.HasErrors() {
		return
	}

	for rows.Next() {
		link := &model.Link{}
		data.Err = rows.StructScan(link)
		if data.HasErrors() {
			return
		}
		links = append(links, link)
	}

	for _, l := range links {
		rec = data.GetItem(l.LinkedMajor, l.LinkedMinor)
		if data.HasErrors() {
			return
		}
		recs = append(recs, rec)
	}
	return
}

func (data *Data) Load(item model.Recorder) {
	data.Err = model.Decode(item)
	if data.HasErrors() {
		return
	}

	linker, isLinker := item.(model.Linker)
	if isLinker {
		data.addLinks(linker)
	}
}

func (data *Data) addLinks(item model.Linker) {
	linkRecs := data.GetLinks(item.GetRecord())
	if data.HasErrors() {
		return
	}

	item.Link(linkRecs...)

	for _, child := range item.Children() {
		linker, isLinker := child.(model.Linker)
		if isLinker {
			data.addLinks(linker)
		}
	}
}

func (data *Data) Save(rec model.Recorder) {
	data.InsertItems(rec)
	if !data.HasErrors() {
		linker, isLinker := rec.(model.Linker)
		if isLinker {
			list, links := data.addLists(linker)
			fmt.Println(list)
			fmt.Println(links)
			data.InsertItems(list...)
			if data.HasErrors() {
				return
			}
			data.InsertLinks(links...)
			if data.HasErrors() {
				return
			}
		}
	}
}

func (data *Data) addLists(parent model.Linker) (list []model.Recorder, links []*model.Link) {
	children := len(parent.Children())
	list = make([]model.Recorder, 0, children)
	list = append(list, parent.Children()...)
	links = make([]*model.Link, 0, children)

	for _, item := range parent.Children() {
		link := model.NewLink(parent, item, 1, 1)
		links = append(links, link)
		linker, isLinker := item.(model.Linker)
		if isLinker {
			r, l := data.addLists(linker)
			list = append(list, r...)
			links = append(links, l...)
		}
	}
	return
}
