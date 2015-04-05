package golight

import (
	"testing"
)

func TestGetCommandDataGreenOn(t *testing.T) {
	data := getCommandData(command{color: Green})

	if data == nil {
		t.Errorf("Error, data was null")
	}

	if len(data) != 8 {
		t.Errorf("Malformed command")
	}

	if data[0] != 0x65 {
		t.Errorf("Wrong major command")
	}

	if data[1] != 0x0C {
		t.Errorf("Wrong minor command")
	}

	if data[2] != Green {
		t.Errorf("Color in wrong place")
	}

	if data[3] != 0x00 {
		t.Errorf("Color in wrong place")
	}

	if data[4] != 0x00 {
		t.Errorf("Wrong HID1 value")
	}

	if data[5] != 0x00 {
		t.Errorf("Wrong HID2 value")
	}

	if data[6] != 0x00 {
		t.Errorf("Wrong HID3 value")
	}

	if data[7] != 0x00 {
		t.Errorf("Wrong HID4 value")
	}
}

func TestGetCommandDataGreenOff(t *testing.T) {
	data := getCommandData(command{color: Green, shutdown:true})

	if data == nil {
		t.Errorf("Error, data was null")
	}

	if len(data) != 8 {
		t.Errorf("Malformed command")
	}

	if data[0] != 0x65 {
		t.Errorf("Wrong major command")
	}

	if data[1] != 0x0C {
		t.Errorf("Wrong minor command")
	}

	if data[2] != 0x00 {
		t.Errorf("Color in wrong place")
	}

	if data[3] != Green {
		t.Errorf("Color in wrong place")
	}

	if data[4] != 0x00 {
		t.Errorf("Wrong HID1 value")
	}

	if data[5] != 0x00 {
		t.Errorf("Wrong HID2 value")
	}

	if data[6] != 0x00 {
		t.Errorf("Wrong HID3 value")
	}

	if data[7] != 0x00 {
		t.Errorf("Wrong HID4 value")
	}
}

func TestGetCommandDataYellowFlash(t *testing.T) {
	data := getCommandData(command{color: Yellow, flash: true})

	if data == nil {
		t.Errorf("Error, data was null")
	}

	if len(data) != 8 {
		t.Errorf("Malformed command")
	}

	if data[0] != 0x65 {
		t.Errorf("Wrong major command")
	}

	if data[1] != 0x14 {
		t.Errorf("Wrong minor command")
	}

	if data[2] != 0x00 {
		t.Errorf("Color in wrong place")
	}

	if data[3] != Yellow {
		t.Errorf("Color in wrong place")
	}

	if data[4] != 0x00 {
		t.Errorf("Wrong HID1 value")
	}

	if data[5] != 0x00 {
		t.Errorf("Wrong HID2 value")
	}

	if data[6] != 0x00 {
		t.Errorf("Wrong HID3 value")
	}

	if data[7] != 0x00 {
		t.Errorf("Wrong HID4 value")
	}
}

func TestGetCommandDataRedFlashOff(t *testing.T) {
	data := getCommandData(command{color: Red, flash: true, shutdown: true})

	if data == nil {
		t.Errorf("Error, data was null")
	}

	if len(data) != 8 {
		t.Errorf("Malformed command")
	}

	if data[0] != 0x65 {
		t.Errorf("Wrong major command")
	}

	if data[1] != 0x14 {
		t.Errorf("Wrong minor command")
	}

	if data[2] != Red {
		t.Errorf("Color in wrong place")
	}

	if data[3] != 0x00 {
		t.Errorf("Color in wrong place")
	}

	if data[4] != 0x00 {
		t.Errorf("Wrong HID1 value")
	}

	if data[5] != 0x00 {
		t.Errorf("Wrong HID2 value")
	}

	if data[6] != 0x00 {
		t.Errorf("Wrong HID3 value")
	}

	if data[7] != 0x00 {
		t.Errorf("Wrong HID4 value")
	}
}
