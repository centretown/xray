package gizmo

import (
	"fmt"
	"image/color"
	"testing"

	"github.com/centretown/xray/dbg"
	"github.com/centretown/xray/model"
	"gopkg.in/yaml.v3"
)

const (
	dbmem  = ":memory:"
	dir    = "/home/dave/xray/game_01/"
	dbfile = dir + "xray_game.db"
	picd   = dir + "pic/"
)

var dbcur = dbmem

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

var fixedPalette = []color.RGBA{
	White,
	Black,
	Red,
	Yellow,
	Green,
	Cyan,
	Blue,
	Magenta,
}

func create_mem_game(t *testing.T) {
	data := dbg.NewGameData("sqlite3", dbcur)
	data.Open()
	if data.Err != nil {
		t.Fatal(data.Err)
	}
	defer data.Close()

	if data.Err != nil {
		t.Fatal(data.Err)
	}

	data.Create()

	const (
		baseInterval = .02
		screenWidth  = 1280
		screenHeight = 720
		fps          = 30
		captureFps   = 25
	)

	game := NewGameSetup(screenWidth, screenHeight, fps)
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

	game.FixedPalette = fixedPalette
	data.Save(game)
	if data.Err != nil {
		t.Fatal(data.Err)
	}
	// read_game(t, game, data)
	load_game(t, game, data)
}

func read_game(t *testing.T, saved *Game, data *dbg.Data) {
	buf, _ := yaml.Marshal(saved)
	fmt.Println("---------")
	fmt.Println("read game")
	fmt.Println(string(buf))

	gameRec := data.GetItemRecord(saved)
	if data.Err != nil {
		t.Fatal(data.Err)
	}

	fmt.Println(gameRec)

	game := &Game{Record: gameRec}
	err := model.Decode(game)
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

	game.Link(linkRecs...)
	fmt.Println(game)

	for _, a := range game.Movers() {
		linkRecs = data.GetLinks(a.GetRecord())
		for i, l := range linkRecs {
			fmt.Println(i, l)
		}
		a.Link(linkRecs...)
	}

	buf, _ = yaml.Marshal(game)
	fmt.Println(string(buf))

}

func load_game(t *testing.T, saved *Game, data *dbg.Data) {

	game := &Game{
		Record: saved.Record,
	}

	data.Load(game)
	if data.Err != nil {
		t.Fatal(data.Err)
	}

	game.Dump()
}
