package layout

import (
	"image/color"

	"github.com/centretown/xray/notes"
)

var colorSet = ColorSet{
	LabelInput:   color.RGBA{255, 0, 255, 255},
	ValueInput:   color.RGBA{255, 255, 255, 255},
	LabelData:    color.RGBA{128, 0, 128, 255},
	ValueData:    color.RGBA{128, 128, 128, 255},
	LabelCurrent: color.RGBA{255, 255, 0, 255},
	ValueCurrent: color.RGBA{0, 255, 255, 255},
	Label:        color.RGBA{128, 128, 0, 255},
	Value:        color.RGBA{0, 128, 128, 255},
}

type ColorSet struct {
	Label        color.RGBA
	Value        color.RGBA
	LabelData    color.RGBA
	ValueData    color.RGBA
	LabelCurrent color.RGBA
	ValueCurrent color.RGBA
	LabelInput   color.RGBA
	ValueInput   color.RGBA
}

type Layout struct {
	FontSize int32
	LabelX   int32
	ValueX   int32
	DeltaY   int32
	Current  int
	Colors   *ColorSet
}

func NewLayout(fontsize int32) *Layout {
	lay := &Layout{}
	lay.Refresh(fontsize)
	lay.Colors = &colorSet
	return lay
}

func (lay *Layout) Refresh(fontsize int32) {
	lay.FontSize = fontsize
	lay.LabelX = fontsize + 3
	lay.ValueX = lay.LabelX + lay.FontSize*15
	lay.DeltaY = lay.FontSize * 2
}

func (lay *Layout) Layout(startY int32,
	notes *notes.Notes, language *notes.LanguageItem,
	draw func(y int32,
		label string, labelColor color.RGBA,
		value string, valueColor color.RGBA)) {

	var (
		y          = startY
		labelColor color.RGBA
		valueColor color.RGBA
	)

	notes.Fetch(language)

	for i, note := range notes.List {
		if i == lay.Current {
			if note.CanAct() {
				labelColor, valueColor = lay.Colors.LabelInput, lay.Colors.ValueInput
			} else {
				labelColor, valueColor = lay.Colors.LabelCurrent, lay.Colors.ValueCurrent
			}
		} else {
			if note.CanAct() {
				labelColor, valueColor = lay.Colors.LabelData, lay.Colors.ValueData
			} else {
				labelColor, valueColor = lay.Colors.Label, lay.Colors.Value
			}
		}
		draw(y, notes.Outputs[i].Label, labelColor,
			notes.Outputs[i].Value, valueColor)
		y += lay.DeltaY
	}
}
