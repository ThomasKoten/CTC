package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

type mainQueue struct {
	queue chan Car
}

var globalWG sync.WaitGroup

func main() {
	nCars := 10
	wg := &sync.WaitGroup{}
	globalWG.Add(nCars)
	mq := mainQueue{make(chan Car, 50)}

	nGasPumps := 3
	nLPGPumps := 2
	nElectricPumps := 1
	gs := &GasStation{Pumps: make(map[FuelType][]*Station)}
	gs.BuildStation(Gas, nGasPumps, 30, 50, 3, wg)
	gs.BuildStation(LPG, nLPGPumps, 4, 8, 5, wg)
	gs.BuildStation(Electric, nElectricPumps, 5, 10, 5, wg)
	fmt.Print(gs)

	go mq.distibuteCars(gs)
	go GenerateCars(nCars, 5, 10, &mq)

	// Wait for all cars to arrive
	globalWG.Wait()
	toPrint := fmt.Sprintf("%+v\n", globalStats)
	toPrint = strings.Replace(toPrint, " ", "\n", -1)
	toPrint = strings.Replace(toPrint, ":{", ":{\n", -1)
	toPrint = strings.Replace(toPrint, "}", "}\n", -1)
	fmt.Println(toPrint)
}

func GenerateCars(nCars, minTime, maxTime int, mq *mainQueue) {
	for i := 0; i < nCars; i++ {
		time.Sleep(time.Duration(rand.Intn(maxTime+1-minTime)+minTime) * time.Millisecond)
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
		fmt.Printf("Car %d is in queue %d type %d \n", car.ID, queueIndex, pumps[queueIndex].Pump_type)
	}
}
