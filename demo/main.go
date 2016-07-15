package main

import (
	"log"
	"time"

	"github.com/jBugman/go-xbee"
	"github.com/tarm/serial"
)

const (
	OUTPUT_DEVICE = "/dev/tty.usbmodem1A1221"
	INPUT_DEVICE  = "/dev/tty.usbserial-A9014IAY"
	ADDRESS       = 2
	BAUD_RATE     = 115200
)

type xbInterface struct {
	serialPort *serial.Port
	address    int
}

func (xb xbInterface) Write(payload []byte) (n int, err error) {
	frame := xbee.NewTxFrame(xb.address, payload).Bytes()
	// log.Println("Written:", frame)
	return xb.serialPort.Write(frame)
}

func write() {
	port := &serial.Config{Name: OUTPUT_DEVICE, Baud: BAUD_RATE}
	s, err := serial.OpenPort(port)
	if err != nil {
		log.Fatal(err)
	}
	xb := xbInterface{s, ADDRESS}
	for {
		xb.Write([]byte{1, 1, 1})
		time.Sleep(2000 * time.Millisecond)
	}
}

func read() {
	port := &serial.Config{Name: INPUT_DEVICE, Baud: BAUD_RATE}
	s, err := serial.OpenPort(port)
	if err != nil {
		log.Fatal(err)
	}
	b := make([]byte, 100)
	var ok, total int
	for {
		n, err := s.Read(b)
		if err != nil {
			log.Println(err)
		}
		if n == 12 {
			ok++
		}
		total++
		log.Println(ok, "/", total, int(100*ok/total))
		// log.Println(n, err)
		// log.Println(b)
	}
}

func main() {
	go read()
	write()
}
