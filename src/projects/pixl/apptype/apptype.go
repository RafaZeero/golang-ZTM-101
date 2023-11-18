package apptype

import (
	"image/color"

	"fyne.io/fyne/v2"
)

type BrushType = int

type PxCanvasConfig struct {
	DrawingArea    fyne.Size     // Drawing area size in pixels
	CanvasOffset   fyne.Position // Canvas offset in pixels
	PxRows, PxCols int           // Number of rows and columns of pixels
	PxSize         int           // Total pixel size in pixels (square)
}

type State struct {
	BrushColor     color.Color
	BrushType      BrushType
	SwatchSelected int    // Index of selected swatch slice
	FilePath       string // Path to file. New one will be empty
}
