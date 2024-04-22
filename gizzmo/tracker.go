package gizzmo

import (
	"fmt"
	"log"

	rl "github.com/centretown/raylib-go/raylib"
	"github.com/centretown/xray/check"
	"github.com/centretown/xray/gizzmo/class"
	"github.com/centretown/xray/gizzmodb/model"
)

var _ Mover = (*Tracker)(nil)
var _ model.Parent = (*Tracker)(nil)

const (
	XAxis int = iota
	YAxis
	ZAxis
)

type TrackerItem struct {
	Bounds       rl.Vector4
	Position     rl.Vector4
	Source       rl.Vector4
	PixelRateX   float64
	PixelRateY   float64 // pixels per second
	PixelRateZ   float64 // pixels per second
	Rotation     float32
	RotationRate float32
	Axes         [3]Axis

	drawer Drawer
}

type Tracker struct {
	model.RecorderClass[TrackerItem]
}

func NewTrackerFromRecord(record *model.Record) *Tracker {
	tr := &Tracker{}
	model.Decode(tr, record)
	tr.Refresh(rl.GetTime(), rl.Vector4{
		X: tr.Content.Bounds.X,
		Y: tr.Content.Bounds.Y})
	return tr
}

func NewTracker(bounds rl.Vector4, rates rl.Vector4,
	minV rl.Vector3, maxV rl.Vector3) *Tracker {
	tr := &Tracker{}
	var _ HasDepth = tr

	item := &tr.Content
	item.Bounds = bounds
	item.PixelRateX = float64(rates.X)
	item.PixelRateY = float64(rates.Y)
	item.PixelRateZ = float64(rates.Z)
	item.RotationRate = rates.W
	item.Rotation = 0

	now := rl.GetTime()
	item.Axes[0].Setup(now, minV.X, maxV.X)
	item.Axes[1].Setup(now, minV.Y, maxV.Y)
	item.Axes[2].Setup(now, minV.Z, maxV.Z)
	// tr.Refresh(now, item.Bounds)
	// tr.Refresh(now, rl.Vector4{X: bounds.X, Y: bounds.Y})
	model.InitRecorder[TrackerItem](tr,
		class.Tracker.String(), int32(class.Tracker))
	return tr
}
func (tr *Tracker) GetDepth() float32 {
	return tr.Content.Axes[2].Position
}

func (tr *Tracker) AddDrawer(dr Drawer) {
	item := &tr.Content
	item.drawer = dr
	item.Source = dr.Bounds()
	tr.Refresh(rl.GetTime(), rl.Vector4{X: item.Bounds.X, Y: item.Bounds.Y})
}

func (tr *Tracker) GetDrawer() Drawer  { return tr.Content.drawer }
func (tr *Tracker) Bounds() rl.Vector4 { return tr.Content.Bounds }
func (tr *Tracker) Draw(v rl.Vector4)  { tr.Content.drawer.Draw(v) }

func (tr *Tracker) Move(can_move bool, now float64) {
	mv := &tr.Content
	x, y, z := &mv.Axes[0], &mv.Axes[1], &mv.Axes[2]
	mv.drawer.Draw(rl.Vector4{
		X: x.Position,
		Y: y.Position,
		Z: z.Position,
		W: mv.Rotation,
	})

	m := check.As[float64](can_move)
	y.Move(now, mv.PixelRateY*m)
	z.Move(now, mv.PixelRateZ*m)

	p := x.Position
	p -= x.Move(now, mv.PixelRateX*m)
	mv.Rotation += mv.RotationRate * float32(p)
}

func (tr *Tracker) Refresh(now float64, bounds rl.Vector4, fs ...func(any)) {
	mv := &tr.Content
	mv.Bounds = bounds
	mv.Bounds.X -= mv.Source.X
	mv.Bounds.Y -= mv.Source.Y
	mv.Axes[0].Refresh(now, rl.Vector2{X: 0, Y: mv.Bounds.X})
	mv.Axes[1].Refresh(now, rl.Vector2{X: 0, Y: mv.Bounds.Y})
	mv.Axes[2].Refresh(now, mv.Axes[2].Extent)
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
