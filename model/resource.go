package model

import (
	"io/fs"
	"os"
	"path/filepath"
)

type ResourceItem struct {
	Title    string
	Path     string
	Category Category
	Scheme   Scheme
	IsDir    bool
	Size     int64
	Content  any
}

type Resource struct {
	Record *Record
	Item   ResourceItem
	Err    error
}

func NewFileResource(path string, category Category, content any) (res *Resource) {
	var (
		abs  string
		info fs.FileInfo
		err  error
		errp = &err
	)

	res = &Resource{}
	res.Item.Content = content
	res.Item.Category = category
	res.Item.Scheme = File

	defer func() {
		res.Err = *errp
	}()

	abs, err = filepath.Abs(path)
	if err == nil {
		res.Item.Path = abs
		info, err = os.Stat(abs)

		if err == nil {
			res.Item.Title = info.Name()
			res.Item.Size = info.Size()
			res.Item.IsDir = info.IsDir()
			res.Record = NewItem(res.Item.Title, category, &res.Item)
		}
	}

	return
}
