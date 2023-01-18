package main

type Subject interface {
	RegisterObserver(obs Observer)
	RemoveObserver(obs Observer)
	NotifyObservers()
}
