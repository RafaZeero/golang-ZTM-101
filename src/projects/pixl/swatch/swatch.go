package swatch

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
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
func NewSwatch(c color.Color, swatchIndex int, clickHandler func(s *Swatch)) *Swatch {
	s := &Swatch{
		Color:        c,
		SwatchIndex:  swatchIndex,
		clickHandler: clickHandler,
	}
	s.ExtendBaseWidget(s)
	return s
}

// Create renderer
func (s *Swatch) CreateRenderer() fyne.WidgetRenderer {
	square := canvas.NewRectangle(s.Color)
	// square.StrokeWidth = 0
	// square.StrokeColor = color.NRGBA(255, 255, 255)
	// square.SetMinSize(fyne.NewSize(20, 20))
	// square.SetPosition(fyne.NewPos(0, 0))
	return &SwatchRenderer{
		square:  *square,
		objects: []fyne.CanvasObject{square},
		parent:  s,
	}
}
