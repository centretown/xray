package access

const DRIVER_MYSQL = "mysql"
const DRIVER_SQLLITE3 = "sqlite3"
const DRIVER_POSTGRES = "postgres"
const DRIVER_CODE = "code"

type KeyValue struct {
	Key   string
	Value string
}

type OutHandler interface {
	Create(name string) error
	CreateEffect(folder, title string, frame any) error
	UpdateEffect(folder, title string, frame any) error
	CreateFolder(folder string) error
	OnExit() error
}

type InHandler interface {
	ReadEffect(folder, title string) (any, error)
	ListEffects(string) ([]string, error)
	ListKeys(string) ([]KeyValue, error)
	ListFolders() ([]string, error)
	OnExit() error
}

type IoHandler interface {
	InHandler
	OutHandler
}
