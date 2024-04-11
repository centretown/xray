package access

import (
	"io"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/yaml.v3"
)

func Load[T any](path string, dest *T) (err error) {
	var rdr *os.File
	rdr, err = os.Open(path)
	if err != nil {
		return
	}
	defer rdr.Close()

	var buf []byte
	buf, err = io.ReadAll(rdr)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(buf, dest)
	return
}

func Save[T any](path string, src *T) (err error) {
	var buf []byte
	buf, err = yaml.Marshal(src)
	if err != nil {
		return
	}
	err = os.WriteFile(path, buf, os.ModeType)
	return
}
