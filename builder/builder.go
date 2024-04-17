package builder

import (
	"flag"
	"log"
	"path/filepath"

	"github.com/centretown/xray/access"
	"github.com/centretown/xray/flagset"
	"github.com/centretown/xray/gizmo"
	"github.com/centretown/xray/gizmodb"
	"github.com/centretown/xray/gizmodb/model"
)

var (
	installBase = "/home/dave/xray/"
	gameName    = "xray_game.db"
	gameKeys    = "game_keys.yaml"
	memoryPath  = ":memory:"
)

func init() {
	flagset.Setup("install", "quick")
}

func Build(custom func(*gizmo.Game)) (*gizmo.Game, bool, error) {

	flag.Parse()

	var (
		flags               = &flagset.Flags
		databasePath string = ""
		inMemory            = false
		install             = false
	)
	// test and install conflict
	// test has higher priority because it's the safest option
	if flags.Quick {
		inMemory = true
		databasePath = memoryPath
	} else if flags.Install != "" {
		install = true
		databasePath = filepath.Join(installBase, flags.Install)
	}

	game, err := create(databasePath, flags, custom, inMemory, install)

	log.Printf("memory: %v, databasePath: %s\n",
		inMemory, databasePath)
	flags.Dump()

	return game, install, err
}

func create(databasePath string, cmd *flagset.FlagSet,
	custom func(*gizmo.Game),
	memory bool, install bool) (game *gizmo.Game, err error) {

	fname := databasePath
	if !memory {
		fname = filepath.Join(databasePath, gameName)
	}

	data := gizmodb.NewGameData("sqlite3", fname)
	defer func() {
		if data != nil {
			data.Close()
			if data.HasErrors() {
				log.Println(data.Err)
				err = data.Err
			}
		}
	}()
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

	game = &gizmo.Game{}
	game.NewGameSetup(screenWidth, screenHeight, fps)
	data.Create(game.GetRecord(), &model.Version{Major: 0, Minor: 1})
	if data.HasErrors() {
		return
	}

	custom(game)

	data.Save(game)
	if data.HasErrors() {
		return
	}

	if !memory {
		if install {
			access.SaveGameKey(filepath.Join(databasePath, gameKeys),
				access.NewGameKeys(game.Record.Major,
					game.Record.Minor))

		} else {
			access.SaveGameKey(filepath.Join(cmd.Install, gameKeys),
				access.NewGameKeys(game.Record.Major,
					game.Record.Minor))
		}
	}

	cmd.Dump()
	game.Dump()
	return
}
