package uart

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

// Test_slice tests the slice method.
func Test_slice(t *testing.T) {
	i := []byte{0x00, 0x00, 0x01, 0x01, 0x02}

	s := slice(i, 2)

	assert.Equal(t, []byte{0x00, 0x00}, s[0])
	assert.Equal(t, []byte{0x01, 0x01}, s[1])
	assert.Equal(t, []byte{0x02}, s[2])
}
