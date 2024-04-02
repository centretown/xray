package access

import (
	"encoding/json"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	YAML = "yaml"
	JSON = "json"
)

type Serializer interface {
	Scan(buffer []byte, frame any) (err error)
	Format(frame any) (buffer []byte, err error)
	FileName(title string) string
}

type YamlSerializer struct {
}

func (yml *YamlSerializer) Scan(buffer []byte, frame any) error {
	return yaml.Unmarshal(buffer, frame)
}

func (yml *YamlSerializer) Format(frame any) ([]byte, error) {
	return yaml.Marshal(frame)
}

func (yml *YamlSerializer) FileName(title string) string {
	return strings.ReplaceAll(title, " ", "_.") + YAML
}

type JsonSerializer struct {
}

func (jsn *JsonSerializer) Scan(buffer []byte, frame any) error {
	return json.Unmarshal(buffer, frame)
}

func (jsn *JsonSerializer) Format(frame any) ([]byte, error) {
	return json.Marshal(frame)
}

func (jsn *JsonSerializer) FileName(title string) string {
	return strings.ReplaceAll(title, " ", "_.") + JSON
}

func UriSerializer(extension string) Serializer {
	switch extension {
	case YAML:
		return &YamlSerializer{}
	}
	return &JsonSerializer{}
}
