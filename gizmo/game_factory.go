package gizmo

import (
	"fmt"
	"log"

	"github.com/centretown/xray/gizmo/class"
	"github.com/centretown/xray/model"
)

func MakeLink(record *model.Record) (recorder model.Recorder) {
	fmt.Println("MakeLink", *record)
	recorder = makeCategory(class.Class(record.Category), record)
	return
}

func makeCategory(cat class.Class, rec *model.Record) model.Recorder {
	switch cat {
	case class.Game:
		return NewGameFromRecord(rec)
	case class.Texture:
		return NewTextureFromRecord(rec)
	case class.Ellipse:
		return NewEllipseFromRecord(rec)
	case class.CellsOrg:
		return NewCellsOrgFromRecord(rec)
	case class.Tracker:
		return NewTrackerFromRecord(rec)
	case class.LifeMover:
		return NewLifeMoverFromRecord(rec)
	case class.LifeGrid:
		return NewLifeGridFromRecord(rec)
	}

	err := fmt.Errorf("unknown category %d(%s)", cat, cat)
	log.Fatal(err)
	return nil
}
