package builder

import (
	"flag"
	"log"
	"path/filepath"

	"github.com/centretown/xray/access"
	"github.com/centretown/xray/dbg"
	"github.com/centretown/xray/flagset"
	"github.com/centretown/xray/gizmo"
	"github.com/centretown/xray/model"
	"gopkg.in/yaml.v3"
)

var (
	installBase = "/home/dave/xray/"
	gameName    = "xray_game.db"
	gameKeys    = "game_keys.yaml"
	memoryPath  = ":memory:"
)

func init() {
	flagset.Setup("install", "test")
}

func Build(builder func(*gizmo.Game, string)) (*gizmo.Game, bool, error) {

	flag.Parse()

	var (
		flags        = &flagset.Flags
		databasePath string
		resourcePath string
		inMemory     = false
		install      = false
	)
	// test and install conflict. test has higher priority
	// because it is the safest option
	if flags.Test {
		inMemory = true
		databasePath = memoryPath
	} else if flags.Install != "" {
		install = true
		databasePath = filepath.Join(installBase, flags.Install)
		resourcePath = filepath.Base(databasePath)
	}

	buf, _ := yaml.Marshal(&flags)

	game, err := create(databasePath, resourcePath, flags, builder, inMemory)
	log.Printf("memory: %v, databasePath: %s, resourcePath: %s\nCommand Line: %s\n",
		inMemory, databasePath, resourcePath, string(buf))
	return game, install, err
}

func create(databasePath, resourcePath string, cmd *flagset.FlagSet,
	builder func(*gizmo.Game, string), memory bool) (game *gizmo.Game, err error) {

	fname := databasePath
	if !memory {
		fname = filepath.Join(databasePath, gameName)
	}

	data := dbg.NewGameData("sqlite3", fname)
	defer func() {
		if data != nil {
			data.Close()
			if data.HasErrors() {
				log.Println(data.Err)
				err = data.Err
			}
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

	game = gizmo.NewGameSetup(resourcePath, screenWidth, screenHeight, fps)

	data.Create(game.Record, &model.Version{Major: 0, Minor: 1})
	if data.HasErrors() {
		return
	}

	builder(game, resourcePath)

	data.Save(game)
	if data.HasErrors() {
		return
	}

	if !memory {
		access.SaveGameKey(filepath.Join(cmd.Install, gameKeys),
			access.NewGameKeys(game.Record.Major,
				game.Record.Minor))
	}

	cmd.Dump()
	game.Dump()
	return
}
