// Package microfont provides tiny fonts for tiny displays.
package microfont

import "golang.org/x/image/font/basicfont"

// Range maps a contiguous range of runes to a contiguous range of glyphs.
type Range = basicfont.Range

// R creates a new range.
func R(start, limit rune, offset int) Range {
	return Range{Low: start, High: limit, Offset: offset}
}

// GlyphIndex returns the index of the glyph representing r, or the
// index of the glyph representing the Unicode replacement character
// U+FFFD if no glyph specifically represents r.  It returns !ok if
// no glyph represents either r or the Unicode replacement character.
func GlyphIndex(ranges []Range, r rune) (index int, ok bool) {
	for {
		for _, rr := range ranges {
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
