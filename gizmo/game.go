package gizmo

import (
	"image"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/centretown/xray/model"

	"github.com/centretown/gpads/gpads"
	"github.com/centretown/gpads/pad"
	"github.com/centretown/xray/gizmo/categories"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var _ = rand.New(rand.NewSource(time.Now().UnixNano()))
var _ model.Recorder = (*Game)(nil)
var _ model.Parent = (*Game)(nil)

type GameItem struct {
	Start         float64
	Current       float64
	Width         int32
	Height        int32
	FrameRate     int32
	InputInterval float64
	FramesCounter int32
	FixedSize     bool

	CaptureCount    int
	CaptureInterval float64
	Capturing       bool
	Paused          bool
	DarkMode        bool
	FixedPalette    []color.RGBA

	path       string
	BackGround color.RGBA
	palette    color.Palette
	colorMap   map[color.Color]uint8

	nextInput       float64
	captureDelay    int
	captureStart    int
	previousCapture float64

	stopChan chan int
	scrChan  chan image.Image

	actors    []Actor
	drawers   []Drawer
	inputters []Inputer
	gamepad   pad.PadG
}

type Game struct {
	GameItem
	Record *model.Record
}

func NewGameSetup(path string, width, height, fps int32) *Game {
	gs := NewGame()
	gs.Width = width
	gs.Height = height
	gs.FrameRate = fps
	gs.path = path
	return gs
}

func NewGame() *Game {
	gs := &Game{}
	record := model.NewRecord("game", int32(categories.Game), &gs.GameItem, model.JSON)
	gs.Setup(record, "")
	return gs
}

func (gs *Game) Setup(record *model.Record, path string) *Game {
	gs.Record = record
	gs.path = path

	gs.Start = 0
	gs.Current = rl.GetTime()
	gs.InputInterval = .2

	gs.CaptureCount = 0
	gs.CaptureInterval = float64(rl.GetFrameTime()) * 2
	gs.Capturing = false
	gs.Paused = false

	gs.BackGround = rl.Black
	gs.stopChan = make(chan int)
	gs.scrChan = make(chan image.Image)
	gs.gamepad = gpads.NewGPads()
	gs.captureStart = 250
	gs.captureDelay = 4
	gs.actors = make([]Actor, 0)
	gs.drawers = make([]Drawer, 0)
	gs.inputters = make([]Inputer, 0)
	return gs
}

func (gs *Game) GetRecord() *model.Record        { return gs.Record }
func (gs *Game) GetItem() any                    { return &gs.GameItem }
func (gs *Game) SetPad(pad pad.PadG)             { gs.gamepad = pad }
func (gs *Game) AddActor(a Actor, after float64) { gs.actors = append(gs.actors, a) }
func (gs *Game) Actors() []Actor                 { return gs.actors }
func (gs *Game) AddDrawer(dr Drawer)             { gs.drawers = append(gs.drawers, dr) }
func (gs *Game) Drawers() []Drawer               { return gs.drawers }

func (gs *Game) Children() (children []model.Recorder) {
	children = make([]model.Recorder, 0, len(gs.actors)+len(gs.drawers))
	for i := range gs.actors {
		children = append(children, gs.actors[i])
	}
	for i := range gs.drawers {
		children = append(children, gs.drawers[i])
	}
	return
}

func (gs *Game) LinkChildren(recs ...*model.Record) {
	var err error
	defer func() {
		if err != nil {
			log.Println(err)
		}
	}()

	for _, rec := range recs {
		cat := categories.Category(rec.Category)
		el := MakeCategory(cat, rec)
		err = model.Decode(el)
		if err == nil {

			switch t := el.(type) {
			case Actor:
				gs.actors = append(gs.actors, t)
			case Drawer:
				gs.drawers = append(gs.drawers, Drawer(t))
			}
		}
	}
}

func (gs Game) Keys() (key string, major, minor int64) {
	major, minor = gs.Record.Major, gs.Record.Minor
	uuid := model.RecordUUID(major, minor)
	key = uuid.String()
	return
}
