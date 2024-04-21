package main

import (
	"math/rand"
	"time"
)

type Car struct {
	ID                 int
	CarType            FuelType
	Queued             bool
	ArrivalTime        time.Time
	QueueUpForPump     time.Time
	RefuelTime         time.Time
	QueueUpForRegister time.Time
	PayTime            time.Time
	ExitTime           time.Time
}

func randomCarType() string {
	var CAR_TYPES = [3]string{"gas", "lpg", "electro"}
	return CAR_TYPES[rand.Intn(3)]
}
func NewCar(id int) Car {
	return Car{
		ID: id,
		// ArrivalTime: arrival,
		CarType: Gas,
		Queued:  false,
	}
}
