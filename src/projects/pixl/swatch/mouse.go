package swatch

import (
	"fyne.io/fyne/v2/driver/desktop"
)

// Mouse down event
func (s *Swatch) MouseDown(*desktop.MouseEvent) {
	s.clickHandler(s)
	s.Selected = true
	s.Refresh()
}

// Mouse up event
func (s *Swatch) MouseUp(*desktop.MouseEvent) {}
