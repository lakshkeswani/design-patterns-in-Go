package displays

import "fmt"

type AccWeather struct {
	temperature float64
	humidity    float64
	pressure    float64
}

func (aw AccWeather) Update(temprature float64, humidity float64, pressure float64) {
	aw.temperature = temprature
	aw.humidity = humidity
	aw.pressure = pressure
	aw.Display()
}

func (aw AccWeather) Display() {
	s := fmt.Sprintf(" Temprature = %.2f , Humidity = %.2f ,Pressure = %.2f  by Acc weather \n", aw.temperature, aw.humidity, aw.pressure)
	fmt.Println(s)
}
