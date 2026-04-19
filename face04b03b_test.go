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

func Test04B03BDrawString(t *testing.T) {
	for filename, s := range map[string]string{
		"hello-world": "Hello, world!",      // no kerning
		"salt-n-bits": "SALT & TV Ta Tb Td", // kerning
		"arrows":      "←↑→↓",               // arrows
	} {
		f, err := os.Open("04b_03b/examples/" + filename + ".png")
		assert.NilError(t, err)
		defer f.Close()

		want, err := png.Decode(f)
		assert.NilError(t, err)
		wantCM := want.ColorModel()

		got := image.NewGray(image.Rect(0, 0, 12*6, 6))
		d := font.Drawer{
			Src:  &image.Uniform{C: color.Gray{Y: 255}},
			Dst:  got,
			Face: Face04B03B,
			Dot:  fixed.P(0, 5),
		}
		d.DrawString(s)

		// f, err = os.Create(filename + ".png")
		// assert.NilError(t, err)
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
