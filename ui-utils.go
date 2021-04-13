package main

import (
	"github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func NewLabel(label string, text string, color termui.Color) *widgets.Paragraph {
	paragraph := widgets.NewParagraph()
	paragraph.Border = false
	paragraph.Title = label
	paragraph.TitleStyle.Fg = color
	paragraph.Text = text
	paragraph.TextStyle.Fg = termui.ColorWhite
	paragraph.PaddingLeft = 1
	return paragraph
}
