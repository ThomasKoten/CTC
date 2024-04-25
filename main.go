package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type mainQueue struct {
	queue chan Car
}

var globalWG sync.WaitGroup

func main() {
	start := time.Now()
	nCars := 1000
	globalWG.Add(nCars)
	mq := mainQueue{make(chan Car, 5)}

	nRegisters := 2
	nGasPumps := 2
	nLPGPumps := 1
	nElectricPumps := 1
	nDieselPumps := 2
	gs := &GasStation{Pumps: make(map[int][]*Station)}
	gs.AddRegisters(nRegisters, 1, 3, 1)
	gs.AddPumps(Gas, nGasPumps, 2, 5, 3)
	gs.AddPumps(LPG, nLPGPumps, 4, 7, 3)
	gs.AddPumps(Electric, nElectricPumps, 5, 10, 3)
	gs.AddPumps(Diesel, nDieselPumps, 3, 6, 3)

	go mq.distibuteCars(gs)
	go GenerateCars(nCars, 1, 2, &mq)

	// Wait for all cars to arrive
	globalWG.Wait()
	toPrint := fmt.Sprintf("%+v\n", globalStats)
	toPrint = strings.Replace(toPrint, " ", "\n", -1)
	toPrint = strings.Replace(toPrint, ":{", ":{\n", -1)
	toPrint = strings.Replace(toPrint, "}", "}\n", -1)
	fmt.Println(toPrint)
	fmt.Print(time.Since(start))
}

func GenerateCars(nCars, minTime, maxTime int, mq *mainQueue) {
	for i := 0; i < nCars; i++ {
		time.Sleep(GetRandomDurationMS(minTime, maxTime))
		car := NewCar(i)
		car.ArrivalTime = time.Now()
		mq.queue <- car
	}
}

func (mq *mainQueue) distibuteCars(gs *GasStation) {
	for {
		car := <-mq.queue
		pumps := gs.Pumps[car.CarType]
		queueIndex := FindShortQueue(pumps)
		car.QueueUpForPump = time.Now()
		pumps[queueIndex].WaitQueue <- car
		// fmt.Printf("Car %d is in queue %d type %d \n", car.ID, queueIndex, pumps[queueIndex].Type)
	}
}
