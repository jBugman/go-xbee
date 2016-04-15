package xbee

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleByteFrame(t *testing.T) {
	data := []byte{0x00}
	expectedFrame := []byte{0x7E, 0x00, 0x01, 0x00, 0xFF}
	assert.Equal(t, expectedFrame, NewFrame(data).Bytes())
}

func TestMultiByteFrame(t *testing.T) {
	data := []byte{0x01, 0x02, 0x03}
	expectedFrame := []byte{0x7E, 0x00, 0x03, 0x01, 0x02, 0x03, 0xF9}
	assert.Equal(t, expectedFrame, NewFrame(data).Bytes())
}
