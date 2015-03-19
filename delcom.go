package main

import(
    "log"
    "github.com/GeertJohan/go.hid"
)
    var vendorId uint16 = 0x0FC5
    var productId uint16 = 0xB080

    // colors
    var green byte = 0x01
    var red byte = 0x02
    var yellow byte = 0x04

func on(color byte) {
    leds, err := hid.Open(vendorId, productId, "")
    if err != nil {
        log.Fatalf("Could not open leds device: %s", err)
    }
    log.Println("Opened connection")
    defer leds.Close()

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
    data[0] = 0x65   // Major command
    data[1] = 0x0C   // Minor command
    data[2] = color  // Data LSB
    data[3] = 0x00   // Data MSB
    data[4] = 0x00   // DataHID[0]
    data[5] = 0x00   // DataHID[1]
    data[6] = 0x00   // DataHID[2]
    data[7] = 0x00   // DataHID[3]

    bytesWritten, err := leds.SendFeatureReport(data)
    if err != nil {
        log.Fatalf("Could not send feature report to do dummy action. %s\n", err)
    }

    if bytesWritten == -1 {
        log.Fatalf("No bytes written")
    }
}

func off(color byte) {
    leds, err := hid.Open(vendorId, productId, "")
    if err != nil {
        log.Fatalf("Could not open leds device: %s", err)
    }
    log.Println("Opened connection")
    defer leds.Close()

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
    data[0] = 0x65   // Major command
    data[1] = 0x0C   // Minor command
    data[2] = 0x00  // Data LSB
    data[3] = color   // Data MSB
    data[4] = 0x00   // DataHID[0]
    data[5] = 0x00   // DataHID[1]
    data[6] = 0x00   // DataHID[2]
    data[7] = 0x00   // DataHID[3]

    bytesWritten, err := leds.SendFeatureReport(data)
    if err != nil {
        log.Fatalf("Could not send feature report to do dummy action. %s\n", err)
    }

    if bytesWritten == -1 {
        log.Fatalf("No bytes written")
    }
}


func main() {
    off(green)
    off(red)
}