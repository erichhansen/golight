package golight

import (
	"github.com/karalabe/hid"
	"log"
)

var vendorId uint16 = 0x0FC5
var productId uint16 = 0xB080

// colors
const (
	Green     byte = 0x01
	Red       byte = 0x02
	Yellow    byte = 0x03
	Blue      byte = 0x04
	BlueGreen byte = 0x05
	Purple    byte = 0x06
	White     byte = 0x07
)

type command struct {
	color    byte
	flash    bool
	shutdown bool
}

func Flash(color byte) {
	// turn on flash
	comm := command{color: color, flash: true}
	comm.sendCommand()
	// then turn on the light
	comm = command{color: color}
	comm.sendCommand()
}

func On(color byte) {
	comm := command{color: color}
	comm.sendCommand()
}

func Off(color byte) {
	// turn off flash
	comm := command{color: color, shutdown: true, flash: true}
	comm.sendCommand()
	// then turn off the light
	comm = command{color: color, shutdown: true}
	comm.sendCommand()
}

func (comm command) sendCommand() {
	devices := hid.Enumerate(vendorId, productId)
	var leds *hid.Device
	var err error
	if len(devices) > 0 {
		leds, err = devices[0].Open()
		if err != nil {
			log.Fatalf("Could not open leds device: %s", err)
		}
		defer leds.Close()
	} else {
		log.Fatalf("Could not find leds device: %s", err)
	}

	data := getCommandData(comm)

	_, err = leds.Write(data)
	if err != nil {
		log.Fatalf("Could not write bytes to delcom LED action. %s\n", err)
	}
}

func getCommandData(comm command) []byte {
	/*
	   Major Command 1 Byte
	   Minor Command 1 Byte
	   Data LSB      1 Byte
	   Data MSB      1 Byte
	   DataHID[0..3] 4 Bytes
	   DataExt[0..7] 8 Bytes (Optional)

	   Only 2 major commands, 101 (0x65) and 102 (0x66)
	   101 - sending 8 bytes
	   102 - sending 16 bytes

	*/
	data := make([]byte, 8)
	data[0] = 0x65 // Major command
	if comm.flash {
		data[1] = 0x14 // Minor command
	} else {
		data[1] = 0x0C // Minor command
	}

	if comm.shutdown {
		if comm.flash {
			data[2] = comm.color // Data LSB
			data[3] = 0x00       // Data MSB
		} else {
			data[2] = 0x00       // Data LSB
			data[3] = comm.color // Data MSB
		}
	} else {
		if comm.flash {
			data[2] = 0x00       // Data LSB
			data[3] = comm.color // Data MSB
		} else {
			data[2] = comm.color // Data LSB
			data[3] = 0x00       // Data MSB
		}
	}

	data[4] = 0x00 // DataHID[0]
	data[5] = 0x00 // DataHID[1]
	data[6] = 0x00 // DataHID[2]
	data[7] = 0x00 // DataHID[3]
	return data
}
