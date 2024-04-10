package model

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

var (
	origin                   uuid.UUID
	originMajor, originMinor int64
	inMajor, inMinor         int64
)

const IN_HOUSE = "018e9522-01c9-77c0-be6c-65526f21ec1a"

func init() {
	origin, _ = uuid.Parse(IN_HOUSE)
	originMajor, originMinor = RecordID(origin)
	inMajor, inMinor = originMajor, originMinor
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
	Title    string
	Category int32
	Content  string
	Encoding Encoding
	Major    int64
	Minor    int64
	Origin   int64
	Originn  int64
	Created  time.Time
	Updated  time.Time
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
		Major:    major,
		Minor:    minor,
		Origin:   originMajor,
		Originn:  originMinor,
		Title:    title,
		Category: category,
		Created:  time.Now(),
		Updated:  time.Now(),
		Encoding: encoding,
		Content:  content,
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

type Recorder interface {
	GetRecord() *Record
	GetItem() any
}

func Decode(recorder Recorder) (err error) {
	rec := recorder.GetRecord()
	err = json.Unmarshal([]byte(rec.Content), recorder.GetItem())
	if err != nil {
		fmt.Println(rec.Content)
		panic(err)
	}
	return
}

type Linker interface {
	Recorder
	Link(...*Record)
	Children() []Recorder
}
