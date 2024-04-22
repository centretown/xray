package gizzmo

import (
	"fmt"
	"image"
	"image/color"
	"math/rand"
	"time"

	"github.com/centretown/xray/gizzmodb"
	"github.com/centretown/xray/gizzmodb/model"

	"github.com/centretown/gpads/gpads"
	"github.com/centretown/gpads/pad"
	rl "github.com/centretown/raylib-go/raylib"
	"github.com/centretown/xray/gizzmo/class"
)

var _ = rand.New(rand.NewSource(time.Now().UnixNano()))
var _ model.Recorder = (*Game)(nil)
var _ model.Parent = (*Game)(nil)

type GameItem struct {
	Title        string
	Description  string
	Rules        string
	Instructions string
	Author       string

	Start         float64
	Current       float64
	InputInterval float64
	FrameRate     int32
	FramesCounter int32

	Width     float32
	Height    float32
	Depth     float32
	FixedSize bool

	CaptureCount    int
	CaptureInterval float64
	Capturing       bool
	Paused          bool
	DarkMode        bool         // defaults to true
	FixedPalette    []color.RGBA // gizzmo colors plus colors added
	BackGround      color.RGBA   // defaults to black
	CaptureDelay    int
	CaptureStart    int

	palette  color.Palette
	colorMap map[color.Color]uint8

	screen          rl.RenderTexture2D
	captureImage    *image.RGBA
	nextInput       float64
	previousCapture float64

	// capture go routine channels
	stopChan chan int
	scrChan  chan *image.RGBA

	// note: movers are also drawers
	movers      []Mover      // movers as loaded
	drawers     []Drawer     // drawers as loaded
	drawerList  []Drawer     // all drawers
	depthList   []DeepDrawer // all drawers plus depth
	textureList []*Texture   // all textures from all drawers
	inputters   []Inputer    // all drawers that are inputters
	gamepad     pad.PadG
}

type Game struct {
	model.RecorderClass[GameItem]
	data *gizzmodb.Data
}

func NewGameFromRecord(record *model.Record) *Game {
	gs := &Game{}
	model.Decode(gs, record)
	gs.setup()
	return gs
}

func (gs *Game) NewGameSetup(width, height, fps int32) {
	model.InitRecorder[GameItem](gs, class.Game.String(),
		int32(class.Game))
	item := &gs.Content
	item.Width = float32(width)
	item.Height = float32(height)
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
	gs.setup()
}

func (gs *Game) setup() {
	item := &gs.Content
	item.stopChan = make(chan int)
	item.scrChan = make(chan *image.RGBA)
	item.gamepad = gpads.NewGPads()
	item.movers = make([]Mover, 0)
	item.drawers = make([]Drawer, 0)
	item.inputters = make([]Inputer, 0)
	item.depthList = make([]DeepDrawer, 0)
	item.textureList = make([]*Texture, 0)
}

func (gs *Game) SetPad(pad pad.PadG)             { gs.Content.gamepad = pad }
func (gs *Game) AddActor(a Mover, depth float32) { gs.Content.movers = append(gs.Content.movers, a) }
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

func (gs *Game) Save(data *gizzmodb.Data) (err error) {
	gs.data = data
	data.Save(gs)
	return data.Err
}
