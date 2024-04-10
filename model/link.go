package model

type Link struct {
	Major    int64
	Minor    int64
	Linked   int64
	Linkedn  int64
	Repeated int32
	Weight   float64
}

func NewLink(item, linked Recorder, repeated int32, weight float64) *Link {
	i := item.GetRecord()
	r := linked.GetRecord()
	link := &Link{
		Major:    i.Major,
		Minor:    i.Minor,
		Linked:   r.Major,
		Linkedn:  r.Minor,
		Repeated: repeated,
		Weight:   weight,
	}
	return link
}
