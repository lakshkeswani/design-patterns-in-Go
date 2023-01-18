package main

import (
	"fmt"

	"observer/displays"
)

func main() {
	station := WeatherStation{
		temperature: 5,
		humidity:    6,
		pressure:    7,
	}
	accweather := displays.AccWeather{}
	realtimeTemp := displays.RealTimeTemp{}
	station.RegisterObserver(accweather)
	station.RegisterObserver(realtimeTemp)
	station.NotifyObservers()
	station.SetMeasurements(1, 2, .2)
	fmt.Sprintf("completed")
}
