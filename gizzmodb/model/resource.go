package model

import (
	"image"
	"image/jpeg"
	"image/png"
	"io/fs"
	"log"
	"os"
	"strings"

	"golang.org/x/image/webp"
)

type Resource struct {
	Name   string
	Path   string
	Scheme Scheme
	IsDir  bool
	Size   int64
	Width  int32
	Height int32
	Err    error `json:"-" yaml:"-"`
}

func SetupResource(res *Resource, path string, classn int32) {
	var (
		info fs.FileInfo
		err  error
	)

	res.Path = path
	res.Scheme = File

	defer func() {
		res.Err = err
	}()

	// path = filepath.Clean(path)
	i := strings.LastIndexByte(path, '.')
	if i > 0 {
		ext := path[i+1:]
		log.Println("NewFileResource", path, ext)
		res.Width, res.Height = GetDimensions(path, ext)
	}

	info, err = os.Stat(path)

	if err == nil {
		res.Name = info.Name()
		res.Size = info.Size()
		res.IsDir = info.IsDir()
	}
}

func GetDimensions(path, ext string) (width int32, height int32) {
	r, err := os.Open(path)
	if err != nil {
		return
	}
	var (
		cfg image.Config
	)
	switch ext {
	case "png":
		cfg, err = png.DecodeConfig(r)
	case "jpeg":
	case "jpg":
		cfg, err = jpeg.DecodeConfig(r)
	case "webp":
		cfg, err = webp.DecodeConfig(r)
	}
	defer r.Close()
	if err != nil {
		return
	}
	return int32(cfg.Width), int32(cfg.Height)
}
