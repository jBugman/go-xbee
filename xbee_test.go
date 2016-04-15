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

func TestTxFramePayload(t *testing.T) {
	data := []byte{0x00, 0x01, 0x02}
	address := 0x1234
	expectedData := []byte{0x01, 0x00, 0x12, 0x34, 0x00, 0x00, 0x01, 0x02}
	assert.Equal(t, expectedData, NewTxFrame(address, data).data)
}

func TestTxFrame(t *testing.T) {
	data := []byte("Hello World")
	address := 0xFFFF
	// by default 'options' is 0
	expectedFrame := []byte{0x7E, 0x00, 0x10, 0x01, 0x00, 0xFF, 0xFF, 0x00, 0x48, 0x65, 0x6C, 0x6C, 0x6F, 0x20, 0x57, 0x6F, 0x72, 0x6C, 0x64, 0xE4}
	assert.Equal(t, expectedFrame, NewTxFrame(address, data).Bytes())
}
