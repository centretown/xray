package gizmo

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"

	"github.com/centretown/xray/model"

	"github.com/centretown/gpads/gpads"
	"github.com/centretown/gpads/pad"
	"github.com/centretown/xray/gizmo/categories"
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
	FixedPalette    []color.RGBA

	path       string
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

func NewGameSetup(width, height, fps int32) *Game {
	gs := NewGame()
	gs.Width = width
	gs.Height = height
	gs.FPS = fps
	return gs
}

func NewGame() *Game {
	gs := &Game{}
	gs.Start = 0
	gs.Current = rl.GetTime()
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
func (gs *Game) SetPad(pad pad.Pad) {
	gs.gamepad = pad
}

func (gs *Game) AddMover(mv *Mover, after float64) {
	gs.movers = append(gs.movers, mv)
}

func (gs *Game) Movers() []*Mover {
	return gs.movers
}

func (gs *Game) Children() (children []model.Recorder) {
	children = make([]model.Recorder, len(gs.movers))
	for i := range gs.movers {
		children[i] = gs.movers[i]
	}
	return
}

func (gs *Game) SetColors() {
	palette := make(color.Palette, 0, len(gs.FixedPalette)+1)
	for _, c := range gs.FixedPalette {
		palette = append(palette, c)
	}
	gs.palette, gs.colorMap =
		CreatePaletteFromTextures(color.RGBA{0, 0, 0, 255}, palette, gs)
}

func (gs *Game) SetColorPalette(backGround color.RGBA,
	palette color.Palette,
	colorMap map[color.Color]uint8) {

	palette = append(palette, color.Transparent)

	gs.backGround = backGround
	gs.palette = palette
	gs.colorMap = colorMap
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
