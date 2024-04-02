package tools

import (
	"testing"

	"github.com/centretown/xray/model"
)

func TestModel(t *testing.T) {

	f := NewPicture("../2d/runt.png")

	if f.Resource.Err != nil {
		t.Fatal(f.Resource.Err)
	}

	t.Log(f.Resource.Record, f.Resource.Item)

}

func TestUrl(t *testing.T) {
	res := model.NewFileResource("../2d/runt.png", model.Picture, "just a runt")
	if res.Err != nil {
		t.Fatal(res.Err)
	}

	t.Log(res.Record, res.Item)

	res = model.NewFileResource("../2d/notthere.png", model.Picture, "not there")
	if res.Err == nil {
		t.Fatal("should be an error")
	}

	t.Log(res.Err)
}
