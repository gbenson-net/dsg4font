package microfont

import (
	"image"
	"iter"
	"maps"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

// Face55vw is a variable width bitmap font with 5-pixel high glyph data.
type Face55vw struct {
	// Ascent is the glyph ascent, in pixels.
	Ascent int
	// Descent is the glyph descent, in pixels.
	Descent int
	// Glyphs contains the glyphs.
	Glyphs []Glyph55vw
	// Ranges map runes to sub-images of Mask. The rune ranges must not
	// overlap, and must be in increasing rune order.
	Ranges []Range
	// Kernings maps rune pairs to kerning adjustments.
	Kernings map[string]int
}

// Close implements [io.Closer].
func (f *Face55vw) Close() error {
	return nil
}

// Glyph implements [font.Face].
func (f *Face55vw) Glyph(dot fixed.Point26_6, r rune) (
	dr image.Rectangle,
	mask image.Image,
	maskp image.Point,
	advance fixed.Int26_6,
	ok bool,
) {
	g, ok := f.glyph(r)
	if ok {
		dr, mask, maskp, advance = g.draw(f, dot)
	}
	return
}

// GlyphBounds implements [font.Face].
func (f *Face55vw) GlyphBounds(r rune) (
	bounds fixed.Rectangle26_6,
	advance fixed.Int26_6,
	ok bool,
) {
	g, ok := f.glyph(r)
	if ok {
		bounds, advance = g.bounds(f)
	}
	return
}

// GlyphAdvance implements [font.Face].
func (f *Face55vw) GlyphAdvance(r rune) (advance fixed.Int26_6, ok bool) {
	g, ok := f.glyph(r)
	if ok {
		advance = g.advance()
	}
	return
}

// Kern implements [font.Face].
func (f *Face55vw) Kern(r0, r1 rune) fixed.Int26_6 {
	return fixed.I(f.Kernings[string([]rune{r0, r1})])
}

// UpdateKernings adds the values from seq to the face's kernings.
// If an adjustment already exists, its value will be overwritten.
func (f *Face55vw) UpdateKernings(seq iter.Seq2[string, int]) {
	if f.Kernings == nil {
		f.Kernings = make(map[string]int)
	}
	maps.Insert(f.Kernings, seq)
}

// Metrics implements [font.Face].
func (f *Face55vw) Metrics() font.Metrics {
	return font.Metrics{
		Height:     fixed.I(f.Ascent + f.Descent),
		Ascent:     fixed.I(f.Ascent),
		Descent:    fixed.I(f.Descent),
		XHeight:    fixed.I(f.Ascent),
		CapHeight:  fixed.I(f.Ascent),
		CaretSlope: image.Point{X: 0, Y: 1},
	}
}

func (f *Face55vw) glyph(r rune) (g Glyph55vw, ok bool) {
	index, ok := f.glyphIndex(r)
	if ok {
		g = f.Glyphs[index]
	}
	return
}

func (f *Face55vw) glyphIndex(r rune) (index int, ok bool) {
	for {
		for _, rr := range f.Ranges {
			if (rr.Low <= r) && (r < rr.High) {
				return int(r-rr.Low) + rr.Offset, true
			}
		}
		if r == '\ufffd' {
			return 0, false
		}
		r = '\ufffd'
	}
}
