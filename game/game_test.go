package game

import (
	"fmt"
	"testing"

	"github.com/centretown/gpads/pad"
	"github.com/centretown/xray/gdb"
	"github.com/centretown/xray/model"
	"gopkg.in/yaml.v3"
)

var (
	dbmem  = ":memory:"
	dir    = "/home/dave/xray/test/"
	dbfile = dir + "db/xray_game.db"
	picd   = dir + "pic/"
	dbcur  = dbmem
)

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

	game := NewGame(gp, screenWidth, screenHeight, fps)
	viewPort := game.GetViewPort()

	hole := NewTexture(picd + "polar.png")
	hole_mv := NewMover(viewPort, 10, 10, 10).AddDrawer(hole)
	game.AddMover(hole_mv, 6)

	ball := NewCircle(20, Cyan)
	ball_mv := NewMover(viewPort, 200, 100, 0).AddDrawer(ball)
	game.AddMover(ball_mv, 1)

	head := NewTexture(picd + "head_300.png")
	head_mv := NewMover(viewPort, 70, 140, 1.75).AddDrawer(head)
	game.AddMover(head_mv, 8)

	gander := NewTexture(picd + "gander.png")
	gander_mv := NewMover(viewPort, 300, 300, 0.5).AddDrawer(gander)
	game.AddMover(gander_mv, 4)

	data.InsertItems(ball, ball_mv,
		hole, hole_mv,
		head, head_mv,
		gander, gander_mv)
	if data.Err != nil {
		t.Fatal(data.Err)
	}
	data.InsertLinks(model.NewLink(ball_mv, ball, 1, 1),
		model.NewLink(hole_mv, hole, 1, 1),
		model.NewLink(head_mv, head, 1, 1),
		model.NewLink(gander_mv, gander, 1, 1),
		model.NewLink(game, hole_mv, 1, 1),
		model.NewLink(game, ball_mv, 1, 1),
		model.NewLink(game, head_mv, 1, 1),
		model.NewLink(game, gander_mv, 1, 1))

	data.InsertItems(game)
	if data.Err != nil {
		t.Fatal(data.Err)
	}

	read_game(t, game, data)
}

func read_game(t *testing.T, game *Game, data *gdb.Data) {
	buf, _ := yaml.Marshal(game)
	fmt.Println("---------")
	fmt.Println("read game")
	fmt.Println(string(buf))

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

	for _, a := range gs.Movers() {
		linkRecs = data.GetLinks(a.GetRecord())
		for i, l := range linkRecs {
			fmt.Println(i, l)
		}
		a.Link(linkRecs...)
	}

	buf, _ = yaml.Marshal(gs)
	fmt.Println(string(buf))

}
