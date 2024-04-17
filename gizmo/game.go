package gizmo

import (
	"fmt"
	"image"
	"image/color"
	"math/rand"
	"time"

	"github.com/centretown/xray/gizmodb"
	"github.com/centretown/xray/model"

	"github.com/centretown/gpads/gpads"
	"github.com/centretown/gpads/pad"
	"github.com/centretown/xray/gizmo/class"
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
	BackGround      color.RGBA

	palette  color.Palette
	colorMap map[color.Color]uint8

	nextInput       float64
	CaptureDelay    int
	CaptureStart    int
	previousCapture float64

	stopChan chan int
	scrChan  chan image.Image

	movers    []Mover
	drawers   []Drawer
	inputters []Inputer
	gamepad   pad.PadG
}

type Game struct {
	model.RecorderG[GameItem]
	data *gizmodb.Data
}

func NewGameFromRecord(record *model.Record) *Game {
	gs := &Game{}
	model.Decode(gs, record)
	gs.Setup()
	return gs
}

func (gs *Game) NewGameSetup(width, height, fps int32) {
	model.InitRecorder[GameItem](gs, class.Game.String(), int32(class.Game))
	item := &gs.Content
	item.Width = width
	item.Height = height
	item.FrameRate = fps
	item.Start = 0
	item.Current = rl.GetTime()
	item.InputInterval = .2
	item.CaptureCount = 0
	item.CaptureInterval = float64(rl.GetFrameTime()) * 2
	item.Capturing = false
	item.Paused = false
	item.BackGround = rl.Black
	item.CaptureStart = 250
	item.CaptureDelay = 4
	gs.Setup()
}

func (gs *Game) Setup() {
	item := &gs.Content
	item.stopChan = make(chan int)
	item.scrChan = make(chan image.Image)
	item.gamepad = gpads.NewGPads()
	item.movers = make([]Mover, 0)
	item.drawers = make([]Drawer, 0)
	item.inputters = make([]Inputer, 0)
}

func (gs *Game) SetPad(pad pad.PadG)             { gs.Content.gamepad = pad }
func (gs *Game) AddActor(a Mover, after float64) { gs.Content.movers = append(gs.Content.movers, a) }
func (gs *Game) Actors() []Mover                 { return gs.Content.movers }
func (gs *Game) AddDrawer(dr Drawer)             { gs.Content.drawers = append(gs.Content.drawers, dr) }
func (gs *Game) Drawers() []Drawer               { return gs.Content.drawers }

func (gs *Game) Children() (children []model.Recorder) {
	children = make([]model.Recorder, 0,
		len(gs.Content.movers)+len(gs.Content.drawers))

	for i := range gs.Content.movers {
		children = append(children, gs.Content.movers[i])
	}
	for i := range gs.Content.drawers {
		children = append(children, gs.Content.drawers[i])
	}
	return
}

func (gs *Game) LinkChild(recorder model.Recorder) {
	mover, ok := recorder.(Mover)
	if ok {
		fmt.Println("Added Mover")
		gs.Content.movers = append(gs.Content.movers, mover)
	} else {
		drawer, ok := recorder.(Drawer)
		if ok {
			fmt.Println("Added Drawer")
			gs.Content.drawers = append(gs.Content.drawers, Drawer(drawer))
		} else {
			panic("Game LinkChildren bad child")
		}
	}
}
