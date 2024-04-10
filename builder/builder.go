package builder

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"github.com/centretown/xray/access"
	"github.com/centretown/xray/cmdl"
	"github.com/centretown/xray/dbg"
	"github.com/centretown/xray/gizmo"
	"github.com/centretown/xray/model"
	"gopkg.in/yaml.v3"
)

var (
	installBase = "/home/dave/xray/"
	gameName    = "xray_game.db"
	gameKeys    = "game_keys.yaml"
)

func init() {
	cmdl.Setup("install", "path", "resource", "major", "minor")
}

func Build(builder func(*gizmo.Game, string)) (*gizmo.Game, error) {

	flag.Parse()

	cmd := &cmdl.Cmdl

	path := ":memory:"
	memory := true
	pathLen := len(cmd.Path)
	if pathLen > 0 {
		if cmd.Install {
			path = filepath.Join(installBase, cmd.Path)
		} else {
			path = filepath.Clean(cmd.Path)
		}
		memory = false
	}

	if len(cmdl.Cmdl.Resource) == 0 {
		cmd.Resource = path
	}

	buf, _ := yaml.Marshal(&cmd)

	fmt.Printf("path: %s, memory: %v\nCommand Line: \n%s",
		path, memory, string(buf))

	return create(path, cmd, builder, memory)
}

func create(path string, cmd *cmdl.CmdLineFlags,
	builder func(*gizmo.Game, string), memory bool) (game *gizmo.Game, err error) {

	fname := path
	if !memory {
		fname = filepath.Join(path, gameName)
	}

	data := dbg.NewGameData("sqlite3", fname)
	defer func() {
		data.Close()
		if data.HasErrors() {
			log.Println(data.Err)
			err = data.Err
		}
	}()

	data.Open()
	if data.HasErrors() {
		return
	}
	const (
		baseInterval = .02
		screenWidth  = 800
		screenHeight = 450
		fps          = 20
		captureFps   = 25
	)

	game = gizmo.NewGameSetup(screenWidth, screenHeight, fps)
	data.Create(game.Record, &model.Version{Major: 0, Minor: 1})
	if data.HasErrors() {
		return
	}

	builder(game, cmdl.Cmdl.Resource)

	data.Save(game)
	if data.HasErrors() {
		return
	}

	if !memory {
		access.SaveGameKeys(filepath.Join(path, gameKeys),
			access.NewGameKeys(game.Record.Major,
				game.Record.Minor))
	}

	cmd.Dump()
	game.Dump()
	return
}
