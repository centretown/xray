package gizmo

import (
	"image/color"

	"github.com/centretown/xray/gizmo/categories"
	"github.com/centretown/xray/model"
	rl "github.com/gen2brain/raylib-go/raylib"
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

func NewEllipse(rgba color.RGBA, width float32, height float32) *Ellipse {
	ell := &Ellipse{}
	InitShape[EllipseItem](&ell.Shape, categories.Ellipse.String(),
		int32(categories.Ellipse), rgba, width, height)
	var _ model.Recorder = ell
	var _ Drawer = ell
	return ell
}

func (ell *Ellipse) Draw(v rl.Vector4) {
	rl.DrawEllipse(int32(v.X), int32(v.Y),
		ell.Content.Dimensions.X, ell.Content.Dimensions.Y,
		ell.Content.Color)
}
