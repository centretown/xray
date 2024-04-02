package model

import (
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
)

const IN_HOUSE = "018e9522-01c9-77c0-be6c-65526f21ec1a"

var origin uuid.UUID

func init() {
	origin, _ = uuid.Parse(IN_HOUSE)
}

type Record struct {
	ID       uuid.UUID `db:"id"`
	Origin   uuid.UUID
	Title    string
	Category string
	Created  time.Time
	Updated  time.Time
	Encoding string
	Content  string
}

func NewItem(title string, category Category, v any) *Record {

	id, err := uuid.NewV7()
	if err != nil {
		log.Fatal(err)
	}

	content, err := json.Marshal(v)
	if err != nil {
		content = []byte(err.Error())
	}

	return &Record{
		ID:       id,
		Origin:   origin,
		Title:    title,
		Category: category.String(),
		Created:  time.Now(),
		Updated:  time.Now(),
		Encoding: "json",
		Content:  string(content),
	}
}
