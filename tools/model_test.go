package tools

import (
	"testing"

	"github.com/centretown/xray/model"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func TestModel(t *testing.T) {
	rl.InitWindow(400, 400, "test model")

	f := NewPicture("../2d/runt.png")

	if f.Resource.Err != nil {
		t.Fatal(f.Resource.Err)
	}

	t.Log(f.Resource.Record, f.Resource.Item)
	f.Unload()

	rl.CloseWindow()
}

func TestUrl(t *testing.T) {
	res := model.NewFileResource("../2d/runt.png", model.Picture)
	if res.Err != nil {
		t.Fatal(res.Err)
	}

	t.Log(res.Record, res.Item)

	res = model.NewFileResource("../2d/notthere.png", model.Picture)
	if res.Err == nil {
		t.Fatal("should be an error")
	}

	t.Log(res.Err)
}
