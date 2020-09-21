package uart

import (
	"io"

	"github.com/acomagu/bufpipe"
	"github.com/basilfx/go-ble-utilities/characteristics"
	"github.com/basilfx/go-utilities/observable"
	"github.com/go-ble/ble"
)

// Service is the interface for the device information BLE service. Only
// properties that are non-nil will be advertised by this service.
type Service struct {
	rx struct {
		reader io.Reader
		writer io.Writer
	}
	tx struct {
		observer *observable.Observable
	}
}

// New initializes a new instance of Service.
func New() *Service {
	s := Service{}

	s.rx.reader, s.rx.writer = bufpipe.New(nil)
	s.tx.observer = observable.New()

	return &s
}

// Create will return an instance of ble.Service, that can be used to advertise
// the device information service.
func (s *Service) Create() *ble.Service {
	service := ble.NewService(ServiceUART)

	// Receive characteristic.
	char1 := service.NewCharacteristic(CharacteristicRx)

	char1.HandleWrite(
		ble.WriteHandlerFunc(func(req ble.Request, rsp ble.ResponseWriter) {
			s.rx.writer.Write(req.Data())
		}))

	// Transmit characteristic.
	char2 := service.NewCharacteristic(CharacteristicTx)

	char2.HandleNotify(
		characteristics.ObservableNotifyHandlerFunc(s.tx.observer))

	return service
}

// Read implements io.Reader.Read by reading from the receive buffer. It will
// block until data is received.
func (s *Service) Read(p []byte) (n int, err error) {
	return s.rx.reader.Read(p)
}

// Write implements io.Writer.Write by writing the data directly to the
// transmit observable.
func (s *Service) Write(p []byte) (n int, err error) {
	for _, c := range slice(p, 20) {
		s.tx.observer.SetValue(c)
	}

	return len(p), nil
}
