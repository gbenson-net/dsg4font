package microfont

import (
	"image"
	"image/color"
	"math/bits"
)

// image56vw is a 6-pixel high alpha image.
type image56vw uint32

// ColorModel implements [image.Image].
func (im image56vw) ColorModel() color.Model {
	return color.AlphaModel
}

// Bounds implements [image.Image].
func (im image56vw) Bounds() image.Rectangle {
	return image.Rectangle{
		Max: image.Point{
			X: im.width(),
			Y: 6,
		},
	}
}

// width returns the width of the image.
func (im image56vw) width() int {
	if im == 0 {
		return 4 // blank
	}
	return ((bits.Len32(uint32(im)) - 1) / 6) + 1
}

// At implements [image.Image].
func (im image56vw) At(x, y int) color.Color {
	c := color.Alpha{}
	if (uint32(im)>>(x*6+y))&1 != 0 {
		c.A = 255
	}
	return c
}
