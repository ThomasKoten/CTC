package main

import (
	"sync"
	"time"
)

var globalStatsLock sync.Mutex
var globalStats Stats

// CarsServed => total count of cars getting serviced at the station
// TotalServiceTime => time car spent waiting in the queue + time spent refueling/paying + time spent waiting to get to another queue
// TotalQueueTime => total time all cars spent waiting in the queue
// AverageTime => average time cars spent waiting in the queue
type Stats struct {
	Pumps struct {
		Gas struct {
			CarsServed       int
			TotalServiceTime time.Duration
			TotalQueueTime   time.Duration
			AverageTime      time.Duration
		}
		LPG struct {
			CarsServed       int
			TotalServiceTime time.Duration
			TotalQueueTime   time.Duration
			AverageTime      time.Duration
		}
		Electric struct {
			CarsServed       int
			TotalServiceTime time.Duration
			TotalQueueTime   time.Duration
			AverageTime      time.Duration
		}
		Diesel struct {
			CarsServed       int
			TotalServiceTime time.Duration
			TotalQueueTime   time.Duration
			AverageTime      time.Duration
		}
	}
	Registers struct {
		CarsServed       int
		TotalServiceTime time.Duration
		TotalQueueTime   time.Duration
		AverageTime      time.Duration
	}
}

// Updates the stats when cars leaves
func updateStats(car Car) {
	defer globalStatsLock.Unlock()

	timeInPumpQueue := car.RefuelTime.Sub(car.QueueUpForPump)
	timeOnPump := car.QueueUpForRegister.Sub(car.QueueUpForPump)
	timeInRegisterQueue := car.PayTime.Sub(car.QueueUpForRegister)
	timeOnRegister := car.ExitTime.Sub(car.QueueUpForRegister)

	globalStatsLock.Lock()

	switch car.CarType {
	//Gas pump stats
	case Gas:
		globalStats.Pumps.Gas.CarsServed++
		globalStats.Pumps.Gas.TotalQueueTime += timeInPumpQueue
		globalStats.Pumps.Gas.TotalServiceTime += timeOnPump

	//LPG pump stats
	case LPG:
		globalStats.Pumps.LPG.CarsServed++
		globalStats.Pumps.LPG.TotalQueueTime += timeInPumpQueue
		globalStats.Pumps.LPG.TotalServiceTime += timeOnPump

	//Electric pump stats
	case Electric:
		globalStats.Pumps.Electric.CarsServed++
		globalStats.Pumps.Electric.TotalQueueTime += timeInPumpQueue
		globalStats.Pumps.Electric.TotalServiceTime += timeOnPump

	//Diesel pump stats
	case Diesel:
		globalStats.Pumps.Diesel.CarsServed++
		globalStats.Pumps.Diesel.TotalQueueTime += timeInPumpQueue
		globalStats.Pumps.Diesel.TotalServiceTime += timeOnPump
	}

	//Registers stats
	globalStats.Registers.CarsServed++
	globalStats.Registers.TotalQueueTime += timeInRegisterQueue
	globalStats.Registers.TotalServiceTime += timeOnRegister

	globalWG.Done()
}

// Reckon all the remaining stats
func finalizeStats() {

	if globalStats.Pumps.Gas.CarsServed > 0 {
		globalStats.Pumps.Gas.AverageTime = (globalStats.Pumps.Gas.TotalQueueTime / time.Duration(globalStats.Pumps.Gas.CarsServed))
	}

	if globalStats.Pumps.LPG.CarsServed > 0 {
		globalStats.Pumps.LPG.AverageTime = (globalStats.Pumps.LPG.TotalQueueTime / time.Duration(globalStats.Pumps.LPG.CarsServed))
	}

	if globalStats.Pumps.Electric.CarsServed > 0 {
		globalStats.Pumps.Electric.AverageTime = (globalStats.Pumps.Electric.TotalQueueTime / time.Duration(globalStats.Pumps.Electric.CarsServed))
	}

	if globalStats.Pumps.Diesel.CarsServed > 0 {
		globalStats.Pumps.Diesel.AverageTime = (globalStats.Pumps.Diesel.TotalQueueTime / time.Duration(globalStats.Pumps.Diesel.CarsServed))
	}

	if globalStats.Registers.CarsServed > 0 {
		globalStats.Registers.AverageTime = globalStats.Registers.TotalQueueTime / time.Duration(globalStats.Registers.CarsServed)
	}
}
