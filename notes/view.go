package notes

type ViewItem struct {
	Width  int64
	Height int64
}

type View struct {
	Composite[ViewItem]
}

func NewScreen(label, format string, view *ViewItem) *View {
	scr := &View{}
	InitComposite(&scr.Composite, label, format, view)
	var _ Note = scr
	return scr
}

func (scr *View) Do(command COMMAND) {}

func (scr *View) Values() []any {
	custom := scr.Custom
	return []any{
		custom.Width,
		custom.Height,
	}
}
