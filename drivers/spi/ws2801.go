package spi

import (
	"gobot.io/x/gobot"
)

// WS2801Driver is a driver for the WS2801 programmable RGB LEDs.
type WS2801Driver struct {
	name       string
	connector  Connector
	connection Connection
	Config
	gobot.Commander

	ledData []byte
}

// NewWS2801Driver creates a new Gobot Driver for WS2801 RGB LEDs.
//
// Params:
//      a *Adaptor - the Adaptor to use with this Driver.
//      count int - how many LEDs are in the array controlled by this driver.
//
// Optional params to use with this driver.
//      spi.WithBits(int):    	number of bits to use with this driver.
//      spi.WithSpeed(int64):   speed in Hz to use with this driver.
//
func NewWS2801Driver(a Connector, count int, options ...func(Config)) *WS2801Driver {
	d := &WS2801Driver{
		name:      gobot.DefaultName("WS2801"),
		connector: a,
		ledData:   make([]byte, count*3),
		Config:    NewConfig(),
	}
	for _, option := range options {
		option(d)
	}
	return d
}

// Name returns the name of the device.
func (d *WS2801Driver) Name() string { return d.name }

// SetName sets the name of the device.
func (d *WS2801Driver) SetName(n string) { d.name = n }

// Connection returns the Connection of the device.
func (d *WS2801Driver) Connection() gobot.Connection { return d.connection.(gobot.Connection) }

// Start initializes the driver.
func (d *WS2801Driver) Start() (err error) {
	bus := d.GetBusOrDefault(d.connector.GetSpiDefaultBus())
	chip := d.GetChipOrDefault(d.connector.GetSpiDefaultChip())
	mode := d.GetModeOrDefault(d.connector.GetSpiDefaultMode())
	bits := d.GetBitsOrDefault(d.connector.GetSpiDefaultBits())
	maxSpeed := d.GetSpeedOrDefault(d.connector.GetSpiDefaultMaxSpeed())

	d.connection, err = d.connector.GetSpiConnection(bus, chip, mode, bits, maxSpeed)
	if err != nil {
		return err
	}
	return nil
}

// Halt stops the driver.
func (d *WS2801Driver) Halt() (err error) {
	d.connection.Close()
	return
}

// NumberOfLeds return the number of LEDs that is controlled by the driver.
func (d *WS2801Driver) NumberOfLeds() int {
	return len(d.ledData) / 3
}

// SetColor sets the ith LED's color to the given R,G,B value.
// A subsequent call to Draw is required to transmit values
// to the LED strip.
func (d *WS2801Driver) SetColor(i int, r, g, b byte) {
	j := i * 3
	d.ledData[j] = r
	d.ledData[j+1] = g
	d.ledData[j+2] = b
}

// Draw displays the RGB values set on the actual LED strip.
func (d *WS2801Driver) Draw() error {
	/*
		tx := make([]byte, 3*len(d.vals))
		var j int
		for i, c := range d.vals {
			j = i * 3
			tx[j] = c.R
			tx[j+1] = c.G
			tx[j+2] = c.B
		}
	*/
	return d.connection.Tx(d.ledData, nil)
}
