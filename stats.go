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
type Stats struct {
	Pumps struct {
		Gas struct {
			CarsServed       int
			TotalServiceTime time.Duration
			TotalQueueTime   time.Duration
		}
		LPG struct {
			CarsServed       int
			TotalServiceTime time.Duration
			TotalQueueTime   time.Duration
		}
		Electric struct {
			CarsServed       int
			TotalServiceTime time.Duration
			TotalQueueTime   time.Duration
		}
		Diesel struct {
			CarsServed       int
			TotalServiceTime time.Duration
			TotalQueueTime   time.Duration
		}
	}
	Registers struct {
		CarsServed       int
		TotalServiceTime time.Duration
		TotalQueueTime   time.Duration
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
