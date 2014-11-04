package mqtt

import (
	"fmt"
	"github.com/hybridgroup/gobot"
	"testing"
)

func initTestMqttAdaptor() *MqttAdaptor {
	return NewMqttAdaptor("mqtt", "localhost:1883")
}

func TestMqttAdaptorConnect(t *testing.T) {
	a := initTestMqttAdaptor()
	gobot.Assert(t, a.Connect(), true)
}

func TestMqttAdaptorFinalize(t *testing.T) {
	a := initTestMqttAdaptor()
	gobot.Assert(t, a.Finalize(), true)
}

func TestMqttAdaptorCannotPublishUnlessConnected(t *testing.T) {
	a := initTestMqttAdaptor()
	data := []byte("o")
	gobot.Assert(t, a.Publish("test", data), false)
}

func TestMqttAdaptorPublishWhenConnected(t *testing.T) {
	a := initTestMqttAdaptor()
	a.Connect()
	data := []byte("o")
	gobot.Assert(t, a.Publish("test", data), true)
}

func TestMqttAdaptorCannotOnUnlessConnected(t *testing.T) {
	a := initTestMqttAdaptor()
	gobot.Assert(t, a.On("hola", func(data interface{}) {
		fmt.Println("hola")
	}), false)
}

func TestMqttAdaptorOnWhenConnected(t *testing.T) {
	a := initTestMqttAdaptor()
	a.Connect()
	gobot.Assert(t, a.On("hola", func(data interface{}) {
		fmt.Println("hola")
	}), true)
}
