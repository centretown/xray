package model

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

const IN_HOUSE = "018e9522-01c9-77c0-be6c-65526f21ec1a"

var (
	origin                   uuid.UUID
	originMajor, originMinor uint64
)

func init() {
	origin, _ = uuid.Parse(IN_HOUSE)
	originMajor, originMinor = RecordID(origin)
}

func RecordID(id uuid.UUID) (major, minor uint64) {
	for i := range 8 {
		major |= uint64(id[i]) << (i * 8)
		minor |= uint64(id[i+8]) << (i * 8)
	}
	return
}

func RecordUUID(major, minor uint64) (id uuid.UUID) {
	for i := range 8 {
		id[i] = uint8(major >> (i * 8))
		id[i+8] = uint8(minor >> (i * 8))
	}
	return
}

type Recordable interface {
	GetRecord() *Record
}

type Record struct {
	Title       string
	Category    string
	Content     string
	Encoding    Encoding
	Major       uint64
	Minor       uint64
	OriginMajor uint64 `db:"origin_major"`
	OriginMinor uint64 `db:"origin_minor"`
	Created     time.Time
	Updated     time.Time
}

func NewRecord(title string, category Category, v any) *Record {

	id, _ := uuid.NewV7()
	major, minor := RecordID(id)

	content, err := json.Marshal(v)
	if err != nil {
		content = []byte(err.Error())
	}

	return &Record{
		Major:       major,
		Minor:       minor,
		OriginMajor: originMajor,
		OriginMinor: originMinor,
		Title:       title,
		Category:    category.String(),
		Created:     time.Now(),
		Updated:     time.Now(),
		Encoding:    JSON,
		Content:     string(content),
	}
}
