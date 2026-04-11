package microfont

import (
	"image"

	"golang.org/x/image/math/fixed"
)

// Glyph56vw is a 6-pixel high glyph.
type Glyph56vw uint32

// draw returns the draw.DrawMask parameters (dr, mask, maskp) to
// draw the glyph at the sub-pixel destination location dot, and
// the glyph's advance width.
func (g Glyph56vw) draw(f *Face56vw, dot fixed.Point26_6) (
	dr image.Rectangle,
	mask image.Image,
	maskp image.Point,
	advance fixed.Int26_6,
) {
	im := image56vw(g)
	w := g.width()
	x := dot.X.Round()
	y := dot.Y.Round()
	dr = image.Rectangle{
		Min: image.Point{
			X: x,
			Y: y - f.Ascent,
		},
		Max: image.Point{
			X: x + w,
			Y: y + f.Descent,
		},
	}
	return dr, im, maskp, fixed.I(w)
}

// bounds returns the bounding box of the glyph, drawn at a dot
// equal to the origin, and that glyph's advance width.
func (g Glyph56vw) bounds(f *Face56vw) (
	bounds fixed.Rectangle26_6,
	advance fixed.Int26_6,
) {
	return fixed.R(0, -f.Ascent, g.width(), f.Descent), g.advance()
}

// advance returns the advance width of the glyph.
func (g Glyph56vw) advance() (advance fixed.Int26_6) {
	return fixed.I(g.width())
}

// width returns the width of the glyph.
func (g Glyph56vw) width() int {
	return image56vw(g).width() + 1
}
