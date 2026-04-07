package microfont

import (
	"iter"
	"maps"
	"unicode"
)

func init() {
	kernFace04B08()
}

// kernFace04B08 applies kernings to Face04B08.
func kernFace04B08() {
	kernings := map[string]int{
		" &": -1,
		"& ": -1,
		"L?": -1,
		"LT": -1,
		`L\`: -1,
		"T/": -1,
		"TJ": -1,
		"T_": -1,
		`\T`: -1,
		"_T": -1,
	}

	Face04B08.UpdateKernings(casefoldKernings(maps.All(kernings)))
}

// casefoldKernings makes kernings case-insensitive.
func casefoldKernings(seq iter.Seq2[string, int]) iter.Seq2[string, int] {
	return func(yield func(string, int) bool) {
		for pair, adj := range seq {
			runes := []rune(pair)

			r0 := runes[0]
			r0s := []rune{unicode.ToUpper(r0), unicode.ToLower(r0)}

			r1 := runes[1]
			r1s := []rune{unicode.ToUpper(r1), unicode.ToLower(r1)}

			for _, r0 := range r0s {
				for _, r1 := range r1s {
					pair := string([]rune{r0, r1})
					if !yield(pair, adj) {
						return
					}
				}
			}
		}
	}
}
