package model

import (
	"log"
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
	Class    string
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

func InitRecord(rec *Record, class string, category int32,
	vContent any, encoding Encoding) {

	id, _ := uuid.NewV7()
	major, minor := RecordID(id)

	buf, err := encoding.Encode(vContent)
	if err != nil {
		buf, _ = encoding.Encode(err.Error())
	}
	content := string(buf)

	rec.Major = major
	rec.Minor = minor
	rec.Origin = originMajor
	rec.Originn = originMinor
	rec.Class = class
	rec.Category = category
	rec.Created = time.Now()
	rec.Updated = time.Now()
	rec.Encoding = encoding
	rec.Content = content
}

func (rec *Record) UpdateContent(v any) string {
	buf, err := rec.Encoding.Encode(v)
	if err != nil {
		buf, _ = rec.Encoding.Encode(err.Error())
	}
	rec.Content = string(buf)
	return rec.Content
}

func (rec *Record) Copy(record *Record) {
	rec.Major = record.Major
	rec.Minor = record.Minor
	rec.Origin = record.Origin
	rec.Originn = record.Originn
	rec.Class = record.Class
	rec.Category = record.Category
	rec.Created = record.Created
	rec.Updated = record.Updated
	rec.Encoding = record.Encoding
	rec.Content = record.Content
}

func (rec *Record) Dump() {
	log.Println(
		"Major", rec.Major,
		"Minor", rec.Minor,
		"Origin", rec.Origin,
		"Originn", rec.Originn,
		"Class", rec.Class,
		"Category", rec.Category,
		"Created", rec.Created,
		"Updated", rec.Updated,
		"Encoding", rec.Encoding,
		"Content", rec.Content)
}

func DumpList(records []*Record) {
	for _, record := range records {
		record.Dump()
	}
}
