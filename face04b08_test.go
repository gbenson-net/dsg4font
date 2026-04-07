package microfont

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"

	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
	"gotest.tools/v3/assert"
)

func TestDefinedGlyphIndex(t *testing.T) {
	for r, want := range map[rune]int{
		' ':      0,
		'!':      1,
		'/':      15,
		'0':      16,
		'9':      25,
		':':      26,
		'@':      32,
		'A':      33,
		'Z':      58,
		'[':      59,
		'`':      64,
		'a':      33, // mapped
		'z':      58, // mapped
		'{':      65, // straight after '`'
		'~':      68,
		'\ufffd': 69,
	} {
		got, ok := Face04B08.glyphIndex(r)
		assert.Check(t, ok)
		assert.Equal(t, got, want)
	}
}

func TestUndefinedGlyphIndex(t *testing.T) {
	const rr = "\n\u007f\u0080\u00ff\ufffc\ufffe\uffff£€¥Ядаȝვეპხის⠊呪術廻戦"
	for _, r := range rr {
		got, ok := Face04B08.glyphIndex(r)
		assert.Check(t, ok)
		assert.Equal(t, got, 69)
	}
}

func TestGlyphWidths(t *testing.T) {
	widths := make(map[rune]int)
	for width, runes := range map[int]string{
		2: "',.:;|",
		3: "!()[]`",
		4: "\"*1<>Ii^{}",
		5: "~",
		7: "&",
	} {
		for _, r := range runes {
			widths[r] = width
		}
	}

	for r := range rune(256) {
		want, ok := widths[r]
		if !ok {
			want = 6
		}

		got, ok := Face04B08.GlyphAdvance(r)
		assert.Check(t, ok)
		assert.Equal(t, got, fixed.I(want))
	}
}

func TestDrawString(t *testing.T) {
	for filename, s := range map[string]string{
		"hello-world": "Hello, world!",
	} {
		f, err := os.Open("04b_08/examples/" + filename + ".png")
		assert.NilError(t, err)
		defer f.Close()

		want, err := png.Decode(f)
		assert.NilError(t, err)
		wantCM := want.ColorModel()

		got := image.NewGray(image.Rect(0, 0, 12*6, 6))
		d := font.Drawer{
			Src:  &image.Uniform{C: color.Gray{Y: 255}},
			Dst:  got,
			Face: Face04B08,
			Dot:  fixed.P(0, 5),
		}
		d.DrawString(s)

		// f, err = os.Create(filename + ".png")
		// defer f.Close()
		// assert.NilError(t, png.Encode(f, got))

		b := got.Bounds()
		assert.Equal(t, b, want.Bounds())

		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				wantColor := want.At(x, y)
				gotColor := wantCM.Convert(got.At(x, y))

				assert.Equal(t, gotColor, wantColor)
			}
		}
	}
}
