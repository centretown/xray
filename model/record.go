package model

import (
	"time"

	"github.com/google/uuid"
)

const IN_HOUSE = "018e9522-01c9-77c0-be6c-65526f21ec1a"

// Decode(rec *model.Record) (err error)
type Recorder interface {
	GetRecord() *Record
	GetItem() any
	Decode(rec *Record) (err error)
}

type Linker interface {
	Recorder
	Link(...*Record)
	Children() []Recorder
}

// type Encoder interface {
// 	Encode(any) (string, error)
// }

var (
	origin                   uuid.UUID
	originMajor, originMinor int64
)

func init() {
	origin, _ = uuid.Parse(IN_HOUSE)
	originMajor, originMinor = RecordID(origin)
}

func RecordID(id uuid.UUID) (major, minor int64) {
	for i := range 8 {
		major |= int64(id[i]) << (i * 8)
		minor |= int64(id[i+8]) << (i * 8)
	}
	return
}

func RecordUUID(major, minor int64) (id uuid.UUID) {
	for i := range 8 {
		id[i] = uint8(major >> (i * 8))
		id[i+8] = uint8(minor >> (i * 8))
	}
	return
}

type Record struct {
	Title       string
	Category    int32
	Content     string
	Encoding    Encoding
	Major       int64
	Minor       int64
	OriginMajor int64 `db:"origin_major"`
	OriginMinor int64 `db:"origin_minor"`
	Created     time.Time
	Updated     time.Time
}

func NewRecord(title string, category int32,
	v any, encoding Encoding) *Record {

	id, _ := uuid.NewV7()
	major, minor := RecordID(id)

	buf, err := encoding.Encode(v)
	if err != nil {
		buf, _ = encoding.Encode(err.Error())
	}
	content := string(buf)

	return &Record{
		Major:       major,
		Minor:       minor,
		OriginMajor: originMajor,
		OriginMinor: originMinor,
		Title:       title,
		Category:    category,
		Created:     time.Now(),
		Updated:     time.Now(),
		Encoding:    encoding,
		Content:     content,
	}
}

func (rec *Record) UpdateContent(v any) string {
	buf, err := rec.Encoding.Encode(v)
	if err != nil {
		buf, _ = rec.Encoding.Encode(err.Error())
	}
	rec.Content = string(buf)
	return rec.Content
}
