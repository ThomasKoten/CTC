package main

import (
	"math/rand"
	"time"
)

type Car struct {
	ID                 int
	CarType            FuelType
	ArrivalTime        time.Time
	QueueUpForPump     time.Time
	RefuelTime         time.Time
	QueueUpForRegister time.Time
	PayTime            time.Time
	ExitTime           time.Time
}

func randomCarType() FuelType {
	return FuelType(rand.Intn(4) + 1)
}

func newCar(id int) Car {
	return Car{
		ID:      id,
		CarType: randomCarType(),
	}
}

// Generates cars with random fuel type and arrival delay
func (mq *MainQueue) GenerateCars(nCars, minTime, maxTime int) {
	for i := 0; i < nCars; i++ {
		time.Sleep(GetRandomDurationMS(minTime, maxTime))
		car := newCar(i)
		car.ArrivalTime = time.Now()
		mq.queue <- car
	}
}

// Distributes cars to the correct pump with the optimal length of queue
func (mq *MainQueue) DistributeCars(gs *GasStation) {
	for {
		car := <-mq.queue
		pumps := gs.Pumps[car.CarType]
		queueIndex := FindShortQueue(pumps)
		car.QueueUpForPump = time.Now()
		pumps[queueIndex].WaitQueue <- car
	}
}
