package gizzmo

import (
	"fmt"
	"log"

	"github.com/centretown/xray/check"
	"github.com/centretown/xray/gizzmo/class"
	"github.com/centretown/xray/gizzmodb/model"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var _ Mover = (*Tracker)(nil)
var _ model.Parent = (*Tracker)(nil)

const (
	XAxis int = iota
	YAxis
	ZAxis
)

type TrackerItem struct {
	Rectangle    rl.Rectangle
	Position     rl.Vector2
	Source       rl.Rectangle
	PixelRateX   float64
	PixelRateY   float64 // pixels per second
	PixelRateZ   float64 // pixels per second
	Depth        float32
	Rotation     float32
	RotationRate float32
	Axes         [3]Axis
	drawer       Drawer
}

type Tracker struct {
	model.RecorderG[TrackerItem]
}

func NewTrackerFromRecord(record *model.Record) *Tracker {
	tr := &Tracker{}
	model.Decode(tr, record)
	tr.Refresh(rl.GetTime(), rl.Vector4{
		X: tr.Content.Rectangle.Width,
		Y: tr.Content.Rectangle.Height})
	return tr
}

func NewTracker(bounds rl.Rectangle,
	pixelRateX, pixelRateY float64, rotationRate float32) *Tracker {

	tr := &Tracker{}
	item := &tr.Content
	item.Rectangle = bounds
	item.PixelRateX = pixelRateX
	item.PixelRateY = pixelRateY
	item.RotationRate = rotationRate
	item.Rotation = 0

	now := rl.GetTime()
	item.Axes[0] = Axis{Extent: item.Rectangle.Width, Direction: 1}
	item.Axes[1] = Axis{Extent: item.Rectangle.Height, Direction: 1}
	item.Axes[2] = Axis{Extent: item.Depth, Direction: 1}
	tr.Refresh(now, rl.Vector4{X: bounds.Width, Y: bounds.Height})
	model.InitRecorder[TrackerItem](tr,
		class.Tracker.String(), int32(class.Tracker))
	return tr
}

func (tr *Tracker) AddDrawer(dr Drawer) {
	mv := &tr.Content
	mv.drawer = dr
	mv.Source = dr.Bounds()
	tr.Refresh(rl.GetTime(), rl.Vector4{X: mv.Rectangle.Width, Y: mv.Rectangle.Height})
}

func (tr *Tracker) GetDrawer() Drawer    { return tr.Content.drawer }
func (tr *Tracker) Bounds() rl.Rectangle { return tr.Content.Rectangle }
func (tr *Tracker) Draw(v rl.Vector4)    { tr.Content.drawer.Draw(v) }

func (tr *Tracker) Move(can_move bool, now float64) {
	mv := &tr.Content
	x, y := &mv.Axes[0], &mv.Axes[1]

	mv.drawer.Draw(rl.Vector4{
		X: mv.Rectangle.X + float32(x.Position),
		Y: mv.Rectangle.Y + float32(y.Position), W: mv.Rotation})

	m := check.As[float64](can_move)
	y.Move(now, mv.PixelRateY*m)

	p := x.Position
	p -= x.Move(now, mv.PixelRateX*m)
	mv.Rotation += mv.RotationRate * float32(p)
}

func (tr *Tracker) Refresh(now float64, bounds rl.Vector4, fs ...func(any)) {
	mv := &tr.Content
	mv.Rectangle = rl.Rectangle{Width: bounds.X, Height: bounds.Y}
	mv.Rectangle.X += mv.Source.Width / 2
	mv.Rectangle.Y += mv.Source.Height / 2
	mv.Rectangle.Width -= mv.Source.Width
	mv.Rectangle.Height -= mv.Source.Height
	mv.Axes[0].Refresh(now, mv.Rectangle.Width)
	mv.Axes[1].Refresh(now, mv.Rectangle.Height)
}

func (tr *Tracker) LinkChild(recorder model.Recorder) {
	dr, ok := recorder.(Drawer)
	if ok {
		tr.AddDrawer(dr)
		fmt.Println("Tracker Added Drawer")
	} else {
		log.Fatal(fmt.Errorf("TrackerLinkChildren: not a Drawer"))
	}
}

func (tr *Tracker) Children() []model.Recorder {
	return []model.Recorder{tr.Content.drawer}
}

func (tr *Tracker) SetPixelRate(pixelRateX, pixelRateY float64) {
	tr.Content.PixelRateY = pixelRateY
	tr.Content.PixelRateX = pixelRateX
}

func (tr *Tracker) GetPixelRate() (float64, float64) {
	return tr.Content.PixelRateX, tr.Content.PixelRateY
}
