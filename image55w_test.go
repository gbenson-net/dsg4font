package microfont

import (
	"image/color"
	"testing"

	"gotest.tools/v3/assert"
)

func TestColorModel(t *testing.T) {
	var im image55vw
	assert.Equal(t, im.ColorModel(), color.AlphaModel)
}

func TestWidth(t *testing.T) {
	for im, want := range map[image55vw]int{
		0:             5,
		1:             1,
		0b11111:       1,
		0b100000:      2,
		0b11111 << 5:  2,
		0b100000 << 5: 3,
		0xffffffff:    7,
	} {
		assert.Equal(t, im.width(), want)
	}
}
