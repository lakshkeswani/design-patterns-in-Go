package displays

import "fmt"

type RealTimeTemp struct {
	temperature float64
	humidity    float64
	pressure    float64
}

func (rtt RealTimeTemp) Update(temprature float64, humidity float64, pressure float64) {
	rtt.temperature = temprature
	rtt.Display()
}

func (aw RealTimeTemp) Display() {
	s := fmt.Sprintf(" Temprature = %.1f by Real Time Temp \n", aw.temperature)
	fmt.Println(s)
}
