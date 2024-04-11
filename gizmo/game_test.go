package gizmo

import (
	"image/color"
	"log"
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
	if data.HasErrors() {
		t.Fatal(data.Err)
	}
	defer data.Close()

	if data.HasErrors() {
		t.Fatal(data.Err)
	}

	const (
		baseInterval = .02
		screenWidth  = 1280
		screenHeight = 720
		fps          = 30
		captureFps   = 25
	)

	game := NewGameSetup(dir, screenWidth, screenHeight, fps)
	data.Create(game.Record, &model.Version{
		Major: 0,
		Minor: 1,
	})

	if data.HasErrors() {
		t.Fatal(data.Err)
	}
	viewPort := game.GetViewPort()

	hole := NewTexture(picd + "polar.png")
	hole_mv := NewMover(viewPort, 10, 10, 10)
	hole_mv.AddDrawer(hole)
	game.AddActor(hole_mv, 6)

	ball := NewCircle(20, Cyan)
	ball_mv := NewMover(viewPort, 200, 100, 0)
	ball_mv.AddDrawer(ball)
	game.AddActor(ball_mv, 1)

	head := NewTexture(picd + "head_300.png")
	head_mv := NewMover(viewPort, 70, 140, 1.75)
	head_mv.AddDrawer(head)
	game.AddActor(head_mv, 8)

	gander := NewTexture(picd + "gander.png")
	gander_mv := NewMover(viewPort, 300, 300, 0.5)
	gander_mv.AddDrawer(gander)
	game.AddActor(gander_mv, 4)

	game.FixedPalette = fixedPalette
	data.Save(game)
	if data.HasErrors() {
		t.Fatal(data.Err)
	}
	// read_game(t, game, data)
	load_game(t, game, data)
}

func read_game(t *testing.T, saved *Game, data *dbg.Data) {
	buf, _ := yaml.Marshal(saved)
	log.Println("---------")
	log.Println("read game")
	log.Println(string(buf))

	gameRec := data.GetItemRecord(saved)
	if data.HasErrors() {
		t.Fatal(data.Err)
	}

	log.Println(gameRec)

	game := &Game{Record: gameRec}
	err := model.Decode(game)
	if err != nil {
		t.Fatal(err)
	}

	linkRecs := data.GetLinks(gameRec)

	if data.HasErrors() {
		t.Fatal(data.Err)
	}
	log.Println("linkRecs")
	for i, l := range linkRecs {
		log.Println(i, l)
	}

	game.Link(linkRecs...)
	log.Println(game)

	for _, a := range game.actors {
		if linker, ok := a.(model.Linker); ok {
			linkRecs = data.GetLinks(a.GetRecord())
			for i, l := range linkRecs {
				log.Println(i, l)
			}
			linker.Link(linkRecs...)
		}
	}

	buf, _ = yaml.Marshal(game)
	log.Println(string(buf))

}

func load_game(t *testing.T, saved *Game, data *dbg.Data) {

	game := &Game{
		Record: saved.Record,
	}

	data.Load(game)
	if data.HasErrors() {
		t.Fatal(data.Err)
	}

	game.Dump()
}
