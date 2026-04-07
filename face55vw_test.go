package microfont

import (
	"image"
	"testing"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"gotest.tools/v3/assert"
)

func TestClose(t *testing.T) {
	assert.NilError(t, Face04B08.Close())
	assert.NilError(t, Face04B08.Close())
}

func TestFontMetrics(t *testing.T) {
	assert.Equal(t, Face04B08.Metrics(), font.Metrics{
		Height:     fixed.I(6),
		Ascent:     fixed.I(5),
		Descent:    fixed.I(1),
		XHeight:    fixed.I(5),
		CapHeight:  fixed.I(5),
		CaretSlope: image.Point{0, 1},
	})
}

func TestGlyphMetrics(t *testing.T) {
	for r, width := range map[rune]int{
		'.': 2,
		'(': 3,
		'I': 4,
		'~': 5,
		'A': 6,
		'&': 7,
		'↑': 6, // not found
	} {
		wantBounds := fixed.R(0, -5, width, +1)
		wantAdvance := fixed.I(width)

		advance, ok := Face04B08.GlyphAdvance(r)
		assert.Check(t, ok)
		assert.Equal(t, advance, wantAdvance)

		bounds, advance, ok := Face04B08.GlyphBounds(r)
		assert.Check(t, ok)
		assert.Equal(t, bounds, wantBounds)
		assert.Equal(t, advance, wantAdvance)

		_, _, _, advance, ok = Face04B08.Glyph(fixed.Point26_6{}, r)
		assert.Check(t, ok)
		assert.Equal(t, advance, wantAdvance)
	}
}

func TestNoGlyph(t *testing.T) {
	var f Face55vw
	dr, mask, maskp, advance, ok := f.Glyph(fixed.Point26_6{}, ' ')
	assert.Check(t, !ok)
	assert.Equal(t, dr, image.Rectangle{})
	assert.Equal(t, mask, nil)
	assert.Equal(t, maskp, image.Point{})
	assert.Equal(t, advance, fixed.I(0))
}

func TestNoGlyphBounds(t *testing.T) {
	var f Face55vw
	bounds, advance, ok := f.GlyphBounds(' ')
	assert.Check(t, !ok)
	assert.Equal(t, bounds, fixed.Rectangle26_6{})
	assert.Equal(t, advance, fixed.I(0))
}

func TestNoGlyphAdvance(t *testing.T) {
	var f Face55vw
	advance, ok := f.GlyphAdvance(' ')
	assert.Check(t, !ok)
	assert.Equal(t, advance, fixed.I(0))
}
