package main

import (
	"context"
	"fmt"
	"io"
	"os"

	uart "github.com/basilfx/go-ble-nordic-uart"
	"github.com/basilfx/go-ble-utilities/device"

	"github.com/go-ble/ble"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

func main() {
	// Create a new device.
	d, err := device.NewDevice()

	if err != nil {
		log.Fatalf("Unable to create new device: %v", err)
	}

	ble.SetDefaultDevice(d)

	// Add Nordic UART service.
	service := uart.New()

	err = ble.AddService(service.Create())

	if err != nil {
		log.Fatalf("Unable to add UART service: %v", err)
	}

	// Start reader and writer to read from and to the UART service.
	go io.Copy(os.Stdout, service)
	go io.Copy(service, os.Stdin)

	// Advertise for specified durantion, or until interrupted by user.
	ctx := ble.WithSigHandler(context.WithCancel(context.Background()))

	err = ble.AdvertiseNameAndServices(ctx, "My Device")

	switch errors.Cause(err) {
	case nil:
	case context.DeadlineExceeded:
		fmt.Printf("done\n")
	case context.Canceled:
		fmt.Printf("canceled\n")
	default:
		log.Fatalf(err.Error())
	}
}
