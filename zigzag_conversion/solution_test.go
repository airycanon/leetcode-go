package zigzag_conversion

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLines(t *testing.T) {
	s := "PAYPALISHIRING"

	assert.Equal(t, "PAHNAPLSIIGYIR", convert(s, 3))
	assert.Equal(t, "PINALSIGYAHRPI", convert(s, 4))
}
