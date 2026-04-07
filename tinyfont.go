// Package microfont provides tiny fonts for tiny displays.
package microfont

import "golang.org/x/image/font/basicfont"

// Range maps a contiguous range of runes to a contiguous range of glyphs.
type Range = basicfont.Range

// R creates a new range.
func R(start, limit rune, offset int) Range {
	return Range{Low: start, High: limit, Offset: offset}
}
