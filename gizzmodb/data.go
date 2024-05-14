package gizzmodb

import (
	"fmt"
	"log"
	"strings"

	"github.com/centretown/xray/access"
	"github.com/centretown/xray/gizzmodb/model"
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
	data := &Data{Schema: SchemaGame,
		Keys: access.NewDataKeys(driver, path)}

	data.dbx, data.Err = sqlx.Connect(data.Keys.Driver, data.Keys.Path)
	if data.HasErrors() {
		log.Fatal(data.Err)
	}
	return data
}

func (data *Data) HasErrors() bool {
	return data.Err != nil
}

func (data *Data) HasError(err error) bool {
	return data.Err == err
}

func (data *Data) Close() {
	data.Err = data.dbx.Close()
}

var ErrNotCreated = fmt.Errorf("no such")

func (data *Data) Create(game *model.Record, version *model.Version) {

	version.Item = game.Major
	version.Itemn = game.Minor

	versions := make([]*model.Version, 0)

	data.GetVersions()

	log.Println(versions, data.Err)

	if data.Err != nil {
		text := data.Err.Error()
		if strings.Index(text, ErrNotCreated.Error()) != 0 {
			log.Panicln("something else", data.Err)
		}
		// first time create
		data.Err = ErrNotCreated
	}

	if data.HasError(ErrNotCreated) {
		for _, sch := range data.Schema.Create {
			log.Println(sch)
			data.dbx.MustExec(sch)
		}
	}
	tx := data.dbx.MustBegin()
	tx.NamedExec(data.Schema.InsertVersion, version)
	data.Err = tx.Commit()
}

func (data *Data) InsertItems(items ...model.Recorder) {
	tx := data.dbx.MustBegin()
	defer func() {
		tx.Commit()
		if data.HasErrors() {
			log.Println("InsertItem", data.Err)
		}
	}()
	for _, item := range items {
		item.GetRecord().UpdateContent(item.GetItem())
		if !data.HasErrors() {
			_, data.Err = tx.NamedExec(data.Schema.InsertItem, item.GetRecord())
			// fmt.Println(data.Schema.InsertItem, item.GetRecord())
			// panic(data.Err)
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
			log.Println("InsertLink", data.Err)
		}
	}()
	for _, link := range links {
		_, data.Err = tx.NamedExec(data.Schema.InsertLink, link)
		if data.HasErrors() {
			return
		}
	}
}

func (data *Data) GetRecord(record *model.Record) {
	data.Err = data.dbx.Get(record, data.Schema.GetItem,
		record.Major, record.Minor)
}

func (data *Data) GetVersion(version *model.Version) *model.Version {
	data.Err = data.dbx.Get(version, data.Schema.GetVersion,
		version.Item, version.Itemn, version.Major, version.Minor)
	return version
}

func (data *Data) GetVersions() error {
	rows := data.dbx.QueryRow(data.Schema.GetVersions)
	data.Err = rows.Err()
	return data.Err
}

func (data *Data) Save(rec model.Recorder) {
	data.InsertItems(rec)

	if !data.HasErrors() {
		parent, isParent := rec.(model.Parent)
		if isParent {
			children, links := data.getChildLinks(parent)
			log.Println(children)
			log.Println(links)
			data.InsertItems(children...)
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

func (data *Data) getChildLinks(parent model.Parent) (children []model.Recorder, links []*model.Link) {
	count := len(parent.Children())
	children = make([]model.Recorder, 0, count)
	children = append(children, parent.Children()...)
	links = make([]*model.Link, 0, count)

	for _, item := range parent.Children() {
		link := model.NewLink(parent, item, 1, 1)
		links = append(links, link)
		linker, isLinker := item.(model.Parent)
		if isLinker {
			r, l := data.getChildLinks(linker)
			children = append(children, r...)
			links = append(links, l...)
		}
	}
	return
}

func (data *Data) LoadLinks(record *model.Record) (records []*model.Record) {

	records = make([]*model.Record, 0)

	var rows *sqlx.Rows

	rows, data.Err = data.dbx.Queryx(data.Schema.GetLinks, record.Major, record.Minor)
	if data.HasErrors() {
		return
	}

	for rows.Next() {
		link := &model.Link{}
		data.Err = rows.StructScan(link)
		if data.HasErrors() {
			return
		}
		rec := &model.Record{Major: link.Linked, Minor: link.Linkedn}
		data.GetRecord(record)
		if data.HasErrors() {
			return
		}
		records = append(records, rec)
	}

	return
}
