package gizmo

import (
	"fmt"
	"log"

	"github.com/centretown/xray/gizmo/categories"
	"github.com/centretown/xray/model"
)

func MakeLink(record *model.Record) (recorder model.Recorder) {
	fmt.Println("MakeLink", *record)
	recorder = makeCategory(categories.Category(record.Category), record)
	return
}

func makeCategory(cat categories.Category, rec *model.Record) model.Recorder {
	switch cat {
	case categories.Game:
		return NewGameFromRecord(rec)
	case categories.Texture:
		return NewTextureFromRecord(rec)
	case categories.Ellipse:
		return NewEllipseFromRecord(rec)
	case categories.CellsOrg:
		return NewCellsOrgFromRecord(rec)
	case categories.Tracker:
		return NewTrackerFromRecord(rec)
	case categories.LifeMover:
		return NewLifeMoverFromRecord(rec)
	case categories.LifeGrid:
		return NewLifeGridFromRecord(rec)
	}

	err := fmt.Errorf("unknown category %d(%s)", cat, cat)
	log.Fatal(err)
	return nil
}
