package model

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

// var _ encoding.TextMarshaler = (*JsonSerializer)(nil)

type YamlSerializer struct{}

func (yml *YamlSerializer) UnmarshalText(buffer []byte, obj any) error {
	return yaml.Unmarshal(buffer, obj)
}

func (yml *YamlSerializer) MarshalText(obj any) ([]byte, error) {
	return yaml.Marshal(obj)
}

type JsonSerializer struct{}

func (jsn *JsonSerializer) Scan(buffer []byte, obj any) error {
	return json.Unmarshal(buffer, obj)
}

func (jsn *JsonSerializer) Format(obj any) ([]byte, error) {
	return json.Marshal(obj)
}

func (jsn *JsonSerializer) MarshalText(r Recorder) ([]byte, error) {
	return json.Marshal(r)
}
