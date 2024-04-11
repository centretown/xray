package model

import (
	"image"
	"image/jpeg"
	"image/png"
	"io/fs"
	"log"
	"os"
	"path/filepath"
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
	Err    error `json:"-"`
}

func NewFileResource(path string, category int32, content any) (res *Resource) {
	var (
		clean string
		info  fs.FileInfo
		err   error
		errp  = &err
	)

	res = &Resource{}
	res.Scheme = File

	defer func() {
		res.Err = *errp
	}()

	clean = filepath.Clean(path)
	i := strings.LastIndexByte(clean, '.')
	if i > 0 {
		ext := clean[i+1:]
		log.Println("NewFileResource", clean, ext)
		res.Width, res.Height = GetDimensions(clean, ext)
	}

	res.Path = clean
	info, err = os.Stat(clean)

	if err == nil {
		res.Name = info.Name()
		res.Size = info.Size()
		res.IsDir = info.IsDir()
	}

	return
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
