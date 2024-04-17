package gizzmodb

import (
	"testing"

	"github.com/centretown/xray/gizzmodb/model"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func TestSchema(t *testing.T) {

	testCases := []struct {
		desc string
		f    func(t *testing.T)
	}{
		{
			desc: "Ping sqlite db",
			f:    ping,
		},
		{
			desc: "Ping sqlite db using gamedata",
			f:    ping_mem_gamedata,
		},
		{
			desc: "create sqlite mem db using gamedata",
			f:    create_mem_gamedata,
		},
		// {
		// 	desc: "create sqlite file db using gamedata",
		// 	f:    create_data,
		// },
		{
			desc: "create sqlite file db using gamedata",
			f:    create_mem_game,
		},
	}

	for _, tC := range testCases {
		if tC.f != nil {
			t.Run(tC.desc, tC.f)
			t.Log("RUN", tC.desc)
		} else {
			t.Log("TODO", tC.desc)
		}
	}
}

func ping(t *testing.T) {
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// force a connection and test that it worked
	err = db.Ping()
	if err != nil {
		t.Fatal(err)
	}

}

func ping_mem_gamedata(t *testing.T) {
	data := NewGameData("sqlite3", ":memory:")
	if data.HasErrors() {
		t.Fatal(data.Err)
	}
	defer data.Close()

	data.Err = data.dbx.Ping()

	if data.HasErrors() {
		t.Fatal(data.Err)
	}
}

func create_mem_gamedata(t *testing.T) {
	data := NewGameData("sqlite3", ":memory:")
	if data.HasErrors() {
		t.Fatal(data.Err)
	}
	defer data.Close()

	data.Err = data.dbx.Ping()

	if data.HasErrors() {
		t.Fatal(data.Err)
	}

	data.Create(&model.Record{}, &model.Version{})

}

func create_mem_game(t *testing.T) {
	// gd := NewGameData("sqlite3", "/home/dave/xray/test/db/gd2.db")
	gd := NewGameData("sqlite3", ":memory:")
	if gd.Err != nil {
		t.Fatal(gd.Err)
	}
	defer gd.Close()

	gd.Err = gd.dbx.Ping()

	if gd.Err != nil {
		t.Fatal(gd.Err)
	}

	gd.Create(&model.Record{}, &model.Version{})

	var paths = []string{
		// "../2d/head.png",
		// "../2d/head_90.png",
		// "../2d/head_300.png",
		// "../2d/hole.png",
		// "../2d/gander.png",
		// "../2d/runt.png",
		// "../2d/polar.png",
		// "../2d/swirl.png",
		// "../2d/GJwBkohXoAAiWN9.jpeg",
	}

	for _, path := range paths {
		var resource = model.Resource{}
		model.InitResource(&resource, path, 0)
		if resource.Err != nil {
			t.Fatal(gd.Err)
		}
		// gd.InsertItem(item)
		if gd.Err != nil {
			t.Fatal(gd.Err)
		}

	}
}
