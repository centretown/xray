package gizmo

import (
	"testing"

	"github.com/centretown/xray/dbg"
)

var (
	lifedir = "/home/dave/xray/life_01/"
	// dbfile = dir + "db/xray_game.db"
	// picd   = dir + "pic/"
	lifefile = lifedir + "xray_game.db"
	lifecur  = lifefile
)

func TestLifeCreate(t *testing.T) {
	createLife(t, dbmem)
}

func TestLifeSave(t *testing.T) {
	createLife(t, lifefile)
}

func createLife(t *testing.T, fname string) {
	data := dbg.NewGameData("sqlite3", fname)
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
		screenWidth  = 800
		screenHeight = 450
		fps          = 20
		captureFps   = 25
	)

	game := NewGameSetup(screenWidth, screenHeight, fps)
	viewPort := game.GetViewPort()

	cells := NewCells(viewPort.Width, viewPort.Height)
	game.AddDrawer(cells)
	game.FixedPalette = fixedPalette
	game.FrameRate = fps
	game.FixedSize = true

	data.Save(game)
	if data.Err != nil {
		t.Fatal(data.Err)
	}

	game.Dump()
}
