package examples

import (
	"github.com/dangerousHobo/go-bluetooth/api"
	"github.com/dangerousHobo/go-bluetooth/devices"
	"github.com/prometheus/common/log"
	debug "github.com/dangerousHobo/go-debug"
)

var dbg = debug.Debug("bluetooth:main")

var adapterID = "hci0"
var tagAddress = "B0:B4:48:C9:4B:01"

//SensorTagTemperatureExample example of reading temperature from a TI sensortag
func SensorTagTemperatureExample() {

	dev, err := api.GetDeviceByAddress(tagAddress)
	if err != nil {
		panic(err)
	}

	if dev == nil {
		panic("Device not found")
	}

	err = dev.Connect()
	if err != nil {
		panic(err)
	}

	sensorTag, err := devices.NewSensorTag(dev)
	if err != nil {
		panic(err)
	}

	// var readTemperature = func() {
	// 	temp, err := sensorTag.Temperature.Read()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	log.Printf("Temperature %v°", temp)
	// }

	var notifyTemperature = func(fn func(temperature float64)) {
		sensorTag.Temperature.StartNotify()
		select {}
	}

	// readTemperature()
	notifyTemperature(func(t float64) {
		log.Infof("Temperature update: %f", t)
	})

}
