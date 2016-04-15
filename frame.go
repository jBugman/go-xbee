package xbee

import "bytes"

const START_BYTE = 0x7E

type frame struct {
	data []byte
}

func NewFrame(data []byte) frame {
	return frame{data: data}
}

func (f frame) checksum() byte {
	var sum int
	for i := 0; i < len(f.data); i++ {
		sum += int(f.data[i])
	}
	return 0xFF - byte(sum&0xFF)
}

func (f frame) Bytes() []byte {
	buf := new(bytes.Buffer)
	buf.WriteByte(START_BYTE)
	buf.Write(intToBytes(len(f.data)))
	buf.Write(f.data)
	buf.WriteByte(f.checksum())
	return buf.Bytes()
}

func intToBytes(x int) []byte {
	return []byte{byte(x >> 8), byte(x & 0xFF)}
}
