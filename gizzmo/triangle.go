package gizzmo

// import (
// 	"image/color"

// 	"github.com/centretown/xray/gizzmo/class"
// 	"github.com/centretown/xray/gizzmodb/model"
// 	rl "github.com/gen2brain/raylib-go/raylib"
// )

// type TriangleItem struct {
// 	Points [3]rl.Vector2
// 	Color  rl.Color
// }

// type Triangle struct {
// 	Shape[TriangleItem]
// }

// func NewTriangle(rgba color.RGBA, v1, v2, v3 rl.Vector2) *Triangle {
// 	tri := &Triangle{}

// 	InitShape[TriangleItem](&tri.Shape, class.Ellipse.String(),
// 		int32(class.Ellipse), rgba, width, height)

// 	var _ model.Recorder = tri
// 	var _ Drawer = tri
// 	return tri
// }

// func (tri *Triangle) Draw(v rl.Vector4) {
// 	item := tri.Content
// 	v1, v2, v3 := item.v1, item.v2, item.v3

// 	rl.DrawTriangle(
// 		int32(v.X),
// 		int32(v.Y),
// 		ell.Content.Dimensions.X, ell.Content.Dimensions.Y,
// 		ell.Content.Color)
// }
