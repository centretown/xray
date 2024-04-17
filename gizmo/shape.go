package gizmo

import (
	"image/color"

	"github.com/centretown/xray/model"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ShapeItem[T any] struct {
	Color      color.RGBA
	Dimensions rl.Vector4
	Custom     T
}

type Shape[T any] struct {
	model.RecorderG[ShapeItem[T]]
}

func ShapeFromRecord[T any](sh *Shape[T], record *model.Record) {
	model.Decode(sh, record)
}

func InitShape[T any](sh *Shape[T], name string, category int32,
	color color.RGBA, width float32, height float32, custom ...float32) {

	model.InitRecorder[Shape[T]](sh, name, int32(category))
	sh.Content.Color = color
	sh.Content.Dimensions.X = width
	sh.Content.Dimensions.Y = height
	var _ model.Recorder = sh
	// var _ Drawer = sh
}

func (sh *Shape[T]) Refresh(float64, rl.Vector4, ...func(any)) {}

func (sh *Shape[T]) Bounds() rl.Rectangle {
	return rl.Rectangle{X: 0, Y: 0, Width: sh.Content.Dimensions.X, Height: sh.Content.Dimensions.Y}
}
