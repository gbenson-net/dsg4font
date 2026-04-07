package microfont

import (
	"image"
	"image/color"
	"math/bits"
)

// image55vw is a 5-pixel high alpha image.
type image55vw uint32

// ColorModel implements [image.Image].
func (im image55vw) ColorModel() color.Model {
	return color.AlphaModel
}

// Bounds implements [image.Image].
func (im image55vw) Bounds() image.Rectangle {
	return image.Rectangle{
		Max: image.Point{
			X: im.width(),
			Y: 5,
		},
	}
}

// width returns the width of the image.
func (im image55vw) width() int {
	if im == 0 {
		return 5 // blank
	}
	return ((bits.Len32(uint32(im)) - 1) / 5) + 1
}

// At implements [image.Image].
func (im image55vw) At(x, y int) color.Color {
	c := color.Alpha{}
	if (uint32(im)>>(x*5+y))&1 != 0 {
		c.A = 255
	}
	return c
}
