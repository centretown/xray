package model

type Link struct {
	Major       int64
	Minor       int64
	LinkedMajor int64 `db:"linked_major"`
	LinkedMinor int64 `db:"linked_minor"`
	Repeated    int32
	Weight      float64
}

func NewLink(item, linked Recorder, repeated int32, weight float64) *Link {
	i := item.GetRecord()
	r := linked.GetRecord()
	link := &Link{
		Major:       i.Major,
		Minor:       i.Minor,
		LinkedMajor: r.Major,
		LinkedMinor: r.Minor,
		Repeated:    repeated,
		Weight:      weight,
	}
	return link
}
