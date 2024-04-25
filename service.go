package main

import (
	"math/rand"
	"time"
)

// Simulates the refueling of a car
func (pump *Station) Refuel(car *Car) {
	servedTime := GetRandomDurationMS(pump.ServeTimeMin, pump.ServeTimeMax)
	time.Sleep(servedTime)
}

// Function finds the optimal station queue for car to go to
func FindShortQueue(stations []*Station) int {
	bestStation := 0
	shortestQueue := len(stations[0].WaitQueue)
	if stations[0].Occupied {
		shortestQueue++
	}

	for id := 1; id < len(stations); id++ {
		tempQueue := len(stations[id].WaitQueue)
		if stations[id].Occupied {
			tempQueue++
		}
		if tempQueue < shortestQueue {
			shortestQueue = tempQueue
			bestStation = id
		}
	}

	return bestStation
}

// Simulates going trought a pump and queueing up for register
func (pump *Station) refuelAndGoPay(registers []*Station) {
	for {
		car := <-pump.WaitQueue
		car.RefuelTime = time.Now()
		pump.Occupied = true
		pump.Refuel(&car)
		queueIndex := FindShortQueue(registers)
		car.QueueUpForRegister = time.Now()
		registers[queueIndex].WaitQueue <- car
		pump.Occupied = false
	}
}

// Simulates paying at a register and leaving the Gas station
func (register *Station) payAndLeave() {
	for {
		car := <-register.WaitQueue
		car.PayTime = time.Now()
		register.Occupied = true
		payingTime := GetRandomDurationMS(register.ServeTimeMin, register.ServeTimeMax)
		time.Sleep(payingTime)
		car.ExitTime = time.Now()
		go updateStats(car)
		register.Occupied = false
	}
}

// Returns random duration between min and max in milliseconds
func GetRandomDurationMS(min, max int) time.Duration {
	return time.Duration(rand.Intn(max+1-min)+min) * time.Millisecond
}
