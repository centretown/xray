package tools

import (
	"fmt"
	"testing"

	"github.com/centretown/gpads/pad"
	"github.com/centretown/xray/gdb"
	"github.com/centretown/xray/model"
	"github.com/centretown/xray/tools/categories"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	dbmem  = ":memory:"
	dir    = "/home/dave/xray/test/"
	dbfile = dir + "db/xray_game.db"
	picd   = dir + "pic/"
	dbcur  = dbmem
)

func TestUrl(t *testing.T) {
	res := model.NewFileResource("../2d/runt.png", int32(categories.Texture), "just a runt")
	if res.Err != nil {
		t.Fatal(res.Err)
	}

	res = model.NewFileResource("../2d/notthere.png", int32(categories.Texture), "not there")
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

func TestSave(t *testing.T) {
	savedb := dbcur
	dbcur = dbfile
	// create_mem_gamedata(t)
	create_mem_game(t)
	dbcur = savedb
}

func create_mem_game(t *testing.T) {
	data := gdb.NewGameData("sqlite3", dbcur)
	data.Open()
	if data.Err != nil {
		t.Fatal(data.Err)
	}
	defer data.Close()

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

	itm := func(f func(model.Recorder), i model.Recorder) {
		f(i)
		if data.Err != nil {
			t.Fatal(data.Err)
		}
		fmt.Println("inserted")
	}

	lnk := func(f func(*model.Link), i *model.Link) {
		f(i)
		if data.Err != nil {
			t.Fatal(data.Err)
		}
		fmt.Println("linked")
	}

	hole := NewTexture(picd + "polar.png")
	hole_mv := NewMover(hole, viewPort, 10, 10, 10)
	game.AddMover(hole_mv, 6)

	ball := NewCircle(20, Cyan)
	ball_mv := NewMover(ball, viewPort, 200, 100, 0)
	game.AddMover(ball_mv, 1)

	head := NewTexture(picd + "head_300.png")
	head_mv := NewMover(head, viewPort, 70, 140, 1.75)
	game.AddMover(head_mv, 8)

	gander := NewTexture(picd + "gander.png")
	gander_mv := NewMover(gander, viewPort, 300, 300, 0.5)
	game.AddMover(gander_mv, 4)

	itm(data.InsertItem, ball)
	itm(data.InsertItem, ball_mv)
	itm(data.InsertItem, hole)
	itm(data.InsertItem, hole_mv)
	itm(data.InsertItem, head)
	itm(data.InsertItem, head_mv)
	itm(data.InsertItem, gander)
	itm(data.InsertItem, gander_mv)

	lnk(data.InsertLink, model.NewLink(ball_mv, ball, 1, 1))
	lnk(data.InsertLink, model.NewLink(hole_mv, hole, 1, 1))
	lnk(data.InsertLink, model.NewLink(head_mv, head, 1, 1))
	lnk(data.InsertLink, model.NewLink(gander_mv, gander, 1, 1))
	lnk(data.InsertLink, model.NewLink(game, hole_mv, 1, 1))
	lnk(data.InsertLink, model.NewLink(game, ball_mv, 1, 1))
	lnk(data.InsertLink, model.NewLink(game, head_mv, 1, 1))
	lnk(data.InsertLink, model.NewLink(game, gander_mv, 1, 1))

	data.InsertItem(game)
	if data.Err != nil {
		t.Fatal(data.Err)
	}

	read_game(t, game, data)
}

func read_game(t *testing.T, game *Game, data *gdb.GameData) {
	fmt.Println(game)
	gameRec := data.GetItemRecord(game)
	if data.Err != nil {
		t.Fatal(data.Err)
	}

	fmt.Println(gameRec)

	gs := &Game{}
	err := gs.Decode(gameRec)
	if err != nil {
		t.Fatal(err)
	}
	linkRecs := data.GetLinks(gameRec)

	if data.Err != nil {
		t.Fatal(data.Err)
	}
	fmt.Println("linkRecs")
	for i, l := range linkRecs {
		fmt.Println(i, l)
	}

	gs.Link(linkRecs...)
	fmt.Println(gs)

	for i, a := range gs.Movers() {
		linkRecs = data.GetLinks(a.GetRecord())
		for i, l := range linkRecs {
			fmt.Println(i, l)
		}
		a.Link(linkRecs...)
		fmt.Println(i, a)
	}

	// fmt.Println(game)
	// fmt.Println(gameS)
}
