package swatch

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"

	"pixl/apptype"
)

type Swatch struct {
	widget.BaseWidget
	Selected     bool
	Color        color.Color
	SwatchIndex  int
	clickHandler func(s *Swatch)
}

// Set color to swatch
func (s *Swatch) SetColor(c color.Color) {
	s.Color = c
	s.Refresh()
}

// New Swatch
func NewSwatch(state *apptype.State, color color.Color, swatchIndex int, clickHandler func(s *Swatch)) *Swatch {
	s := &Swatch{
		Selected:     false,
		Color:        color,
		clickHandler: clickHandler,
		SwatchIndex:  swatchIndex,
	}
	s.ExtendBaseWidget(s)
	return s
}

// Create renderer
func (s *Swatch) CreateRenderer() fyne.WidgetRenderer {
	square := canvas.NewRectangle(s.Color)

	return &SwatchRenderer{
		square:  *square,
		objects: []fyne.CanvasObject{square},
		parent:  s,
	}
}
