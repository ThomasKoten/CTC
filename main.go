package main

import (
	"fmt"
	"sync"
	"time"
)

type MainQueue struct {
	queue chan Car
}

var globalWG sync.WaitGroup

func main() {
	fmt.Println("Starting gas station simulation")
	start := time.Now()
	config := loadConfig("config.yaml")

	nCars := config.Cars.Count
	globalWG.Add(nCars)
	mq := MainQueue{make(chan Car, config.Cars.MainQueueLength)}
	gs := newGasStation()

	gs.AddRegisters(config.Registers.Count, config.Registers.HandleTimeMin, config.Registers.HandleTimeMax, config.Registers.QueueLengthMax)
	for fuelType, stationConfig := range config.Stations {
		gs.AddFuelPumps(fuelType, stationConfig.Count, stationConfig.ServeTimeMin, stationConfig.ServeTimeMax, stationConfig.QueueLengthMax)
	}

	go mq.DistributeCars(gs)
	go mq.GenerateCars(nCars, config.Cars.ArrivalTimeMin, config.Cars.ArrivalTimeMax)

	// Wait for all cars to leave and the process of stats evaluation
	globalWG.Wait()

	finalizeStats()
	printStats()
	fmt.Print("Total elapsed time: ", time.Since(start))
}
