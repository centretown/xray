package gizmo

import (
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
	FrameRate     int32
	InputInterval float64
	FramesCounter int32
	FixedSize     bool

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
	drawers []Drawer
	gamepad pad.PadG
}

type Game struct {
	GameItem
	Record *model.Record
}

func NewGameSetup(width, height, fps int32) *Game {
	gs := NewGame()
	gs.Width = width
	gs.Height = height
	gs.FrameRate = fps
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

	gs.stopChan = make(chan int)
	gs.scrChan = make(chan image.Image)
	gs.gamepad = gpads.NewGPads()
	gs.captureStart = 250
	gs.captureDelay = 4
	gs.movers = make([]*Mover, 0)
	gs.drawers = make([]Drawer, 0)
	return gs
}

func (gs *Game) GetRecord() *model.Record          { return gs.Record }
func (gs *Game) GetItem() any                      { return &gs.GameItem }
func (gs *Game) SetPad(pad pad.PadG)               { gs.gamepad = pad }
func (gs *Game) AddMover(mv *Mover, after float64) { gs.movers = append(gs.movers, mv) }
func (gs *Game) Movers() []*Mover                  { return gs.movers }
func (gs *Game) AddDrawer(dr Drawer)               { gs.drawers = append(gs.drawers, dr) }
func (gs *Game) Drawers() []Drawer                 { return gs.drawers }

func (gs *Game) Children() (children []model.Recorder) {
	children = make([]model.Recorder, 0, len(gs.movers)+len(gs.drawers))
	for i := range gs.movers {
		children = append(children, gs.movers[i])
	}
	for i := range gs.drawers {
		children = append(children, gs.drawers[i])
	}
	return
}

func (gs *Game) SetColors() {
	palette := make(color.Palette, 0, len(gs.FixedPalette))
	for _, c := range gs.FixedPalette {
		palette = append(palette, c)
	}
	gs.palette, gs.colorMap =
		CreatePaletteFromTextures(color.RGBA{0, 0, 0, 255}, palette, gs)
}

func (gs *Game) SetColorPalette(backGround color.RGBA,
	palette color.Palette,
	colorMap map[color.Color]uint8) {

	// palette = append(palette, color.Transparent)

	gs.backGround = backGround
	gs.palette = palette
	gs.colorMap = colorMap
}

func (gs *Game) Link(recs ...*model.Record) {
	for _, rec := range recs {
		cat := categories.Category(rec.Category)
		switch cat {
		case categories.Mover:
			{
				mv := &Mover{Record: rec}
				err := model.Decode(mv)
				if err == nil {
					gs.movers = append(gs.movers, mv)
				} else {
					fmt.Println(err)
				}
			}
		case categories.Cells:
			{
				cs := &Cells{Record: rec}
				err := model.Decode(cs)
				if err == nil {
					gs.drawers = append(gs.drawers, cs)
				} else {
					fmt.Println(err)
				}
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
