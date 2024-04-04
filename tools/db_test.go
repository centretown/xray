package tools

import (
	"fmt"
	"testing"

	"github.com/centretown/gpads/pad"
	"github.com/centretown/xray/dbio"
	"github.com/centretown/xray/model"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jmoiron/sqlx"
)

var (
	dbmem  = ":memory:"
	dir    = "/home/dave/xray/test/"
	dbfile = dir + "db/xray_game.db"
	picd   = dir + "pic/"
	dbcur  = dbmem
)

func TestUrl(t *testing.T) {
	res := model.NewFileResource("../2d/runt.png", model.Picture, "just a runt")
	if res.Err != nil {
		t.Fatal(res.Err)
	}

	t.Log(res.Record, res.Item)

	res = model.NewFileResource("../2d/notthere.png", model.Picture, "not there")
	if res.Err == nil {
		t.Fatal("should be an error")
	}

	t.Log(res.Err)
}

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
		{
			desc: "create sqlite gamedata",
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
	db, err := sqlx.Open("sqlite3", dbcur)
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
	gd := dbio.NewGameData("sqlite3", dbcur)
	gd.Open()
	if gd.Err != nil {
		t.Fatal(gd.Err)
	}
	defer gd.Close()

	gd.Err = gd.DB.Ping()

	if gd.Err != nil {
		t.Fatal(gd.Err)
	}
}

func create_mem_gamedata(t *testing.T) {
	gd := dbio.NewGameData("sqlite3", dbcur)
	gd.Open()
	if gd.Err != nil {
		t.Fatal(gd.Err)
	}
	defer gd.Close()

	gd.Err = gd.DB.Ping()

	if gd.Err != nil {
		t.Fatal(gd.Err)
	}

	gd.Create()
}

func TestSave(t *testing.T) {
	savedb := dbcur
	dbcur = dbfile
	// create_mem_gamedata(t)
	create_mem_game(t)
	dbcur = savedb
}

func create_mem_game(t *testing.T) {
	data := dbio.NewGameData("sqlite3", dbcur)
	data.Open()
	if data.Err != nil {
		t.Fatal(data.Err)
	}
	defer data.Close()

	data.Err = data.DB.Ping()

	if data.Err != nil {
		t.Fatal(data.Err)
	}

	data.Create()

	var (
		gp pad.Pad
	)

	const (
		baseInterval = .02
		screenWidth  = 1280
		screenHeight = 720
		fps          = 30
		captureFps   = 25
	)

	game := NewGame(gp, fps)
	viewPort := rl.RectangleInt32{X: 0, Y: 0, Width: 1000, Height: 500}

	chk := func(f func(model.Recordable), i model.Recordable) {
		f(i)
		if data.Err != nil {
			t.Fatal(data.Err)
		}
		fmt.Println("inserted")
	}

	hole := NewPicture(picd + "polar.png")
	chk(data.InsertItem, hole)
	hole_mv := NewMover(hole, viewPort, 10, 10, 10)
	chk(data.InsertItem, hole_mv)
	game.AddActor(hole_mv, 6)

	ball := NewCircle(20, Cyan)
	chk(data.InsertItem, ball)
	ball_mv := NewMover(ball, viewPort, 200, 100, 0)
	chk(data.InsertItem, ball_mv)
	game.AddActor(ball_mv, 1)

	head := NewPicture(picd + "head_300.png")
	chk(data.InsertItem, head)
	head_mv := NewMover(head, viewPort, 70, 140, 1.75)
	chk(data.InsertItem, head_mv)
	game.AddActor(head_mv, 8)

	gander := NewPicture(picd + "gander.png")
	chk(data.InsertItem, gander)
	gander_mv := NewMover(gander, viewPort, 300, 300, 0.5)
	chk(data.InsertItem, gander_mv)
	game.AddActor(gander_mv, 4)

	data.InsertItem(game)
	if data.Err != nil {
		t.Fatal(data.Err)
	}
	fmt.Println(game)

}
