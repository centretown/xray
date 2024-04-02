package access

import (
	"io"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/yaml.v3"
)

type Access struct {
	Driver   string
	Path     string
	User     string
	Password string
	Host     string
	Port     string
	Database string
	DB       sqlx.DB
}

func NewAccess(driver, path string) *Access {
	acc := &Access{
		Driver: driver,
		Path:   path,
	}
	return acc
}

func (acc *Access) Load(path string) (err error) {
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

	acc = &Access{}
	err = yaml.Unmarshal(buf, acc)
	return
}

func (acc *Access) Save(path string) (err error) {
	var buf []byte
	buf, err = yaml.Marshal(acc)
	if err != nil {
		return
	}
	err = os.WriteFile(path, buf, os.ModePerm)
	return
}

// hide the secret parts
type AccessView struct {
	Driver   string
	Path     string
	Host     string
	Port     string
	Database string
}

func NewAccesView(acc *Access) *AccessView {
	acv := &AccessView{
		Driver:   acc.Driver,
		Path:     acc.Path,
		Host:     acc.Host,
		Port:     acc.Port,
		Database: acc.Database,
	}
	return acv
}
