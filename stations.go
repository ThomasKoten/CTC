// station.go
package main

import (
	"sync"
)

type FuelType int

const (
	Gas FuelType = iota
	LPG
	Electric
	Diesel
)

type Station struct {
	Serve_time_min int
	Serve_time_max int
	Pump_type      FuelType
	// Pump_count     int
	// ServicePlace   []*Car
	WaitQueue  chan Car
	queueMutex sync.Mutex
}

type GasStation struct {
	Pumps map[FuelType][]*Station
}

// Function to create a new Station
func NewStation(serveTimeMin, serveTimeMax int, fuelType FuelType, maxQueueLength int) *Station {
	return &Station{
		Serve_time_min: serveTimeMin,
		Serve_time_max: serveTimeMax,
		Pump_type:      fuelType,
		WaitQueue:      make(chan Car, maxQueueLength),
	}
}

func (gs *GasStation) BuildStation(pumpType FuelType, nPumps, minTime, maxTime, maxQueue int, wg *sync.WaitGroup) {
	gs.Pumps[pumpType] = make([]*Station, nPumps)
	for i := 0; i < nPumps; i++ {
		gs.Pumps[pumpType][i] = NewStation(minTime, maxTime, pumpType, maxQueue)
		go gs.Pumps[pumpType][i].fillUpAndLeave()
	}
}
