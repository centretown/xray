package gizzmo

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/centretown/xray/gizzmodb"
	"github.com/centretown/xray/gizzmodb/model"
	"github.com/centretown/xray/layout"
	"github.com/centretown/xray/notes"
	msg "github.com/centretown/xray/notes"

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

	InputInterval float64
	FrameRate     int64

	Width  float32
	Height float32
	Depth  float32

	FixedWidth  float32
	FixedHeight float32
	FixedDepth  float32
	FixedSize   bool

	BackGround      color.RGBA // defaults to black
	CaptureDelay    float64
	CaptureDuration float64

	built              float64
	paused             bool
	fullscreen         bool
	screenstate        ResizeState
	monitorNum         int
	monitorWidth       int
	monitorHeight      int
	monitorRefreshRate int
	currentFrameRate   int64
	screenWidth        int64
	screenHeight       int64

	currentTime float64

	beginCapturing bool
	capturing      bool
	endCapturing   bool

	renderTexture rl.RenderTexture2D
	captureImage  *rl.Image
	captureCount  int64
	captureTotal  int64
	captureEnd    float64
	// capture go routine channels
	captureStop   chan int
	captureSource chan *rl.Image

	aspectRatio   float32
	commandState  bool
	notes         *msg.Notes
	captureNotes  *msg.Notes
	layout        *layout.Layout
	languages     *notes.LanguageList
	language      *notes.LanguageItem
	languageIndex int
	note          int

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
	content := &gs.Content
	content.FixedWidth, content.Width = float32(width), float32(width)
	content.FixedHeight, content.Height = float32(height), float32(height)
	content.built = rl.GetTime()
	content.currentTime = rl.GetTime()
	content.InputInterval = .15
	content.BackGround = rl.Black
	content.CaptureDuration = 15
	content.FrameRate = 30
	//TODO?
	content.CaptureDelay = 4
	gs.setup()
}

func (gs *Game) setup() {
	content := &gs.Content
	content.paused = false
	content.aspectRatio = content.Width / content.Height
	content.captureStop = make(chan int)
	content.captureSource = make(chan *rl.Image)
	content.capturing = false
	content.screenstate = RESIZE_NORMAL

	content.layout = layout.NewLayout(20)
	content.gamepad = gpads.NewGPads()

	content.movers = make([]Mover, 0)
	content.drawers = make([]Drawer, 0)
	content.inputters = make([]Inputer, 0)
	content.depthList = make([]DeepDrawer, 0)
	content.textureList = make([]*Texture, 0)
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
		gs.Content.movers = append(gs.Content.movers, mover)
	} else {
		drawer, ok := recorder.(Drawer)
		if ok {
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
