package game

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"

	"github.com/centretown/xray/model"

	"github.com/centretown/gpads/gpads"
	"github.com/centretown/gpads/pad"
	"github.com/centretown/xray/game/categories"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var _ model.Recorder = (*Game)(nil)
var _ model.Linker = (*Game)(nil)

type GameItem struct {
	Start         float64
	Current       float64
	Width         int32
	Height        int32
	FPS           int32
	InputInterval float64

	CaptureCount    int
	CaptureInterval float64
	Capturing       bool
	Paused          bool

	backGround color.RGBA
	palette    color.Palette
	colorMap   map[color.Color]uint8

	nextInput       float64
	captureDelay    int
	captureStart    int
	previousCapture float64

	stopChan chan int
	scrChan  chan image.Image

	movers  []*Mover
	gamepad pad.Pad
}

type Game struct {
	GameItem
	Record *model.Record
}

func NewGame(gp pad.Pad, width, height, fps int32) *Game {
	gs := &Game{}
	gs.Start = 0
	gs.Current = rl.GetTime()
	gs.Width = width
	gs.Height = height
	gs.FPS = fps
	gs.InputInterval = .2

	gs.CaptureCount = 0
	gs.CaptureInterval = float64(rl.GetFrameTime()) * 2
	gs.Capturing = false
	gs.Paused = false

	gs.Record = model.NewRecord("game", int32(categories.Game), &gs.GameItem, model.JSON)

	gs.stopChan = make(chan int)
	gs.scrChan = make(chan image.Image)
	gs.gamepad = gpads.NewGPads()
	gs.captureStart = 250
	gs.captureDelay = 4
	gs.movers = make([]*Mover, 0)
	return gs
}

func (gs *Game) GetRecord() *model.Record { return gs.Record }
func (gs *Game) GetItem() any             { return &gs.GameItem }

func (gs *Game) AddMover(mv *Mover, after float64) {
	gs.movers = append(gs.movers, mv)
}

func (gs *Game) Movers() []*Mover {
	return gs.movers
}

func (gs *Game) Children() (rcds []model.Recorder) {
	rcds = make([]model.Recorder, len(gs.movers))
	for i := range gs.movers {
		rcds[i] = gs.movers[i]
	}
	return
}

func (gs *Game) SetColors(BG color.RGBA, pal color.Palette, m map[color.Color]uint8) {
	gs.backGround = BG
	gs.palette = pal
	gs.colorMap = m
}

func (gs *Game) Link(recs ...*model.Record) {
	for _, rec := range recs {
		mv := &Mover{}
		err := mv.Decode(rec)
		if err == nil {
			gs.movers = append(gs.movers, mv)
		} else {
			fmt.Println(err)
		}
	}
}

func (gs *Game) Decode(rec *model.Record) (err error) {
	gs.Record = rec

	cat := categories.Category(rec.Category)
	if cat == categories.Game {
		err = json.Unmarshal([]byte(rec.Content), &gs.GameItem)
	} else {
		err = fmt.Errorf("wrong category want %s have %s",
			categories.Game, cat)
	}

	return
}
