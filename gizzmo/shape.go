package gizzmo

import (
	"image/color"

	"github.com/centretown/xray/gizzmodb/model"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ShapeItem[T any] struct {
	Color           color.RGBA
	Dimensions      rl.Vector4
	ScaleToScreen   bool
	KeepAspectRatio bool

	Custom T
}

type Shape[T any] struct {
	model.RecorderClass[ShapeItem[T]]
}

func ShapeFromRecord[T any](sh *Shape[T], record *model.Record) {
	model.Decode(sh, record)
}

func InitShape[T any](sh *Shape[T], name string, classn int32,
	color color.RGBA, width float32, height float32, depth float32, custom ...float32) {

	model.InitRecorder[Shape[T]](sh, name, int32(classn))
	sh.Content.Color = color
	sh.Content.Dimensions.X = width
	sh.Content.Dimensions.Y = height
	sh.Content.Dimensions.Z = depth
	var _ model.Recorder = sh
	var _ HasDepth = sh
}

func (sh *Shape[T]) Refresh(now float64, v rl.Vector4, f ...func(any)) {
	if sh.Content.ScaleToScreen {
		sh.Content.Dimensions.X = float32(rl.GetScreenWidth())
		sh.Content.Dimensions.Y = float32(rl.GetScreenHeight())
	}
}

func (sh *Shape[T]) Bounds() rl.Vector4 {
	return rl.Vector4{X: 0, Y: 0, Z: sh.Content.Dimensions.Z}
}

func (sh *Shape[T]) GetDepth() float32 { return sh.Content.Dimensions.Z }
