package uart

import "github.com/go-ble/ble"

// ServiceUART is the UUID of the Nordic UART Service.
// https://developer.nordicsemi.com/nRF_Connect_SDK/doc/latest/nrf/
var ServiceUART = ble.MustParse("6e400001-b5a3-f393-e0a9-e50e24dcca9e")

// UUIDs of the characteristics.
var (
	CharacteristicRx = ble.MustParse("6e400002-b5a3-f393-e0a9-e50e24dcca9e")
	CharacteristicTx = ble.MustParse("6e400003-b5a3-f393-e0a9-e50e24dcca9e")
)
