package main

import (
	"fmt"
	"reflect"
)

type WeatherStation struct {
	observers   []Observer
	temperature float64
	humidity    float64
	pressure    float64
}

func (ws WeatherStation) RemoveObserver(obs Observer) {
	for i, observer := range ws.observers {
		if reflect.DeepEqual(observer, obs) {
			ws.observers = append(ws.observers[:i-1], ws.observers[i+1:]...)
		}
	}
}

func (ws WeatherStation) NotifyObservers() {
	for _, observer := range ws.observers {
		observer.Update(ws.temperature, ws.humidity, ws.pressure)
	}
}

func (ws *WeatherStation) RegisterObserver(obs Observer) {
	if len(ws.observers) == 0 {
		ws.observers = []Observer{obs}
	} else {
		ws.observers = append(ws.observers, obs)
	}
	fmt.Println("registered sucessfully")

}

func (ws WeatherStation) SetMeasurements(temprature float64, humidity float64, pressure float64) {
	ws.temperature = temprature
	ws.humidity = humidity
	ws.pressure = pressure
	fmt.Println("values updated sucessfully")
	ws.NotifyObservers()
}
