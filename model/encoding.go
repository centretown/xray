package model

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

type Encoding int32

const (
	RAW Encoding = iota
	JSON
	YAML
	XML
	CSV
)

func (en Encoding) Encode(v any) (buf []byte, err error) {
	switch en {
	case RAW, XML, CSV:
		panic("NOT IMPLEMENTED")
	case YAML:
		return yaml.Marshal(v)
	case JSON:
	default:
	}
	return json.Marshal(v)
}

func (en Encoding) Decode(buf []byte, v any) (err error) {
	switch en {
	case RAW, XML, CSV:
		panic("NOT IMPLEMENTED")
	case YAML:
		return yaml.Unmarshal(buf, v)
	case JSON:
	default:
	}
	return json.Unmarshal(buf, v)
}

//go:generate stringer -type=Encoding
