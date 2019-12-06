package spi

import (
	"testing"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/gobottest"
)

var _ gobot.Driver = (*WS2801Driver)(nil)

func initTestWS2801Driver() *WS2801Driver {
	d := NewWS2801Driver(&TestConnector{}, 4)
	return d
}

func TestDriverWS2801Start(t *testing.T) {
	d := initTestDriver()
	gobottest.Assert(t, d.Start(), nil)
}

func TestDriverWS2801Halt(t *testing.T) {
	d := initTestDriver()
	d.Start()
	gobottest.Assert(t, d.Halt(), nil)
}

func TestDriverWS2801LEDs(t *testing.T) {
	d := initTestWS2801Driver()
	d.Start()
	d.SetColor(0, 255, 0, 0)
	d.SetColor(1, 0, 255, 0)
	d.SetColor(2, 0, 0, 255)
	d.SetColor(3, 255, 255, 255)

	gobottest.Assert(t, d.Draw(), nil)
}
