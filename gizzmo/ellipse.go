package gizzmo

import (
	"image/color"

	rl "github.com/centretown/raylib-go/raylib"
	"github.com/centretown/xray/gizzmo/class"
	"github.com/centretown/xray/gizzmodb/model"
)

type EllipseItem struct {
}

type Ellipse struct {
	Shape[EllipseItem]
}

func NewEllipseFromRecord(record *model.Record) *Ellipse {
	ell := &Ellipse{}
	ShapeFromRecord(&ell.Shape, record)
	return ell
}

func NewEllipse(rgba color.RGBA, width float32, height float32, depth float32) *Ellipse {
	ell := &Ellipse{}
	SetupShape[EllipseItem](&ell.Shape, class.Ellipse.String(),
		int32(class.Ellipse), rgba, width, height, depth)
	var _ model.Recorder = ell
	var _ Drawer = ell
	return ell
}

func (ell *Ellipse) Draw(v rl.Vector4) {
	rl.DrawEllipse(int32(v.X), int32(v.Y),
		ell.Content.Dimensions.X/v.Z, ell.Content.Dimensions.Y/v.Z,
		ell.Content.Color)
}
