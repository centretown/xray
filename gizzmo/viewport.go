package gizzmo

import rl "github.com/centretown/raylib-go/raylib"

type ViewPort struct {
	Width  float32
	Height float32
	Depth  float32
	Spin   float32
}

func (vp *ViewPort) FromVector4(v rl.Vector4) {
	vp.Width = v.X
	vp.Height = v.Y
	vp.Depth = v.Z
	vp.Spin = v.W
}
