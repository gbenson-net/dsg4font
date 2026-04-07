package microfont

import (
	"slices"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

func TestCasefold(t *testing.T) {
	var got []string
	for pair, adj := range Face04B08.Kernings {
		assert.Equal(t, adj, -1)
		got = append(got, pair)
	}

	slices.Sort(got)
	assert.Equal(
		t,
		strings.Join(got, ":"),
		` &:& :L?:LT:L\:Lt:T/:TJ:T_:Tj:\T:\t:_T:_t:l?:lT:l\:lt:t/:tJ:t_:tj`,
	)
}
