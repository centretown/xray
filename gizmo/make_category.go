package gizmo

import (
	"fmt"
	"log"

	"github.com/centretown/xray/gizmo/categories"
	"github.com/centretown/xray/model"
)

func MakeCategory(cat categories.Category, rec *model.Record) model.Recorder {
	switch cat {
	case categories.Texture:
		return &Texture{Record: rec}
	case categories.Circle:
		return &Circle{Record: rec}
	case categories.CellsOrg:
		return &CellsOrg{Record: rec}
	case categories.Mover:
		return &Mover{Record: rec}
	case categories.CellsMover:
		return &CellsMover{Record: rec}
	case categories.Cells:
		return &Cells{Record: rec}
	}

	err := fmt.Errorf("unknown category %d(%s)", cat, cat)
	log.Fatal(err)
	return nil
}

func MakeLink[T any](add func(T), min, max int,
	recs ...*model.Record) (err error) {

	var (
		cat    categories.Category
		dr     model.Recorder
		ok     bool
		typ    T
		length int
	)

	defer func() {
		if err != nil {
			log.Fatal(err)
		}
	}()

	if length = len(recs); length < min || length > max {
		err = fmt.Errorf("too many links want range [%d-%d] have %d",
			min, max, len(recs))
		return
	}

	for _, rec := range recs {
		cat = categories.Category(rec.Category)
		dr = MakeCategory(cat, rec)
		if err = model.Decode(dr); err != nil {
			return
		}
		if typ, ok = dr.(T); ok {
			add(typ)
		}
	}

	return
}
