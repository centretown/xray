package model

import (
	"io/fs"
	"os"
	"path/filepath"
)

type ResourceItem struct {
	Name   string
	Path   string
	Scheme Scheme
	IsDir  bool
	Size   int64
	Item   any
}

type Resource struct {
	*Record
	ResourceItem
	Err error
}

func NewFileResource(path string, category Category, content any) (res *Resource) {
	var (
		abs  string
		info fs.FileInfo
		err  error
		errp = &err
	)

	res = &Resource{}
	res.Item = content
	res.Scheme = File

	defer func() {
		res.Err = *errp
	}()

	abs, err = filepath.Abs(path)
	if err == nil {
		res.Path = abs
		info, err = os.Stat(abs)

		if err == nil {
			res.Name = info.Name()
			res.Size = info.Size()
			res.IsDir = info.IsDir()
			res.Record = NewRecord(res.Name, category, &res.ResourceItem)
		}
	}

	return
}

func (res *Resource) GetRecord() *Record {
	return res.Record
}
