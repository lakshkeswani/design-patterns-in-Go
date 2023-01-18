package main

type Observer interface {
	Update(temprature float64, humidity float64, pressure float64)
}
