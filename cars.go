package main

import (
	"math/rand"
	"time"
)

type Car struct {
	ID                 int
	CarType            int
	Queued             bool
	ArrivalTime        time.Time
	QueueUpForPump     time.Time
	RefuelTime         time.Time
	QueueUpForRegister time.Time
	PayTime            time.Time
	ExitTime           time.Time
}

func randomCarType() int {
	return rand.Intn(4)
}
func NewCar(id int) Car {
	return Car{
		ID:      id,
		CarType: randomCarType(),
		Queued:  false,
	}
}
