package xbee

import "bytes"

const (
	API_16BIT_TX_ID = 0x01
)

func NewTxFrame(address int, data []byte) frame {
	var id byte = 0
	// Identifies the UART data frame for the host to correlate with a subsequent ACK (acknowledgement).
	// Setting Frame ID to ‘0' will disable response frame.

	var options byte = 0x00
	// 0x01 = Disable ACK
	// 0x04 = Send packet with Broadcast Pan ID
	// All other bits must be set to 0.

	return ManualTxFrame(address, id, options, data)
}

func ManualTxFrame(address int, frameId byte, options byte, data []byte) frame {
	buf := new(bytes.Buffer)
	buf.WriteByte(API_16BIT_TX_ID)
	buf.WriteByte(frameId)
	buf.Write(intToBytes(address))
	buf.WriteByte(options)
	buf.Write(data)

	return frame{data: buf.Bytes()}
}
