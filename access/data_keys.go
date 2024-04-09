package access

import "github.com/jmoiron/sqlx"

type DataKeys struct {
	Driver   string
	Path     string
	User     string
	Password string
	Host     string
	Port     string
	Database string
	DB       sqlx.DB
}

// hide the secret parts
type DataView struct {
	Driver   string
	Path     string
	Host     string
	Port     string
	Database string
}

func NewAccesView(acc *DataKeys) *DataView {
	acv := &DataView{
		Driver:   acc.Driver,
		Path:     acc.Path,
		Host:     acc.Host,
		Port:     acc.Port,
		Database: acc.Database,
	}
	return acv
}

func NewDataKeys(driver, path string) *DataKeys {
	dk := &DataKeys{
		Driver: driver,
		Path:   path,
	}
	return dk
}

func LoadDataKeys(path string) (dk *DataKeys, err error) {
	return Load(path, &DataKeys{})
}

func SaveDataKeys(path string, dk *DataKeys) (err error) {
	return Save(path, dk)
}
