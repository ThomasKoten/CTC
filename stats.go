package main

import (
	"sync"
	"time"
)

var globalStatsLock sync.Mutex
var globalStats Stats

type Stats struct {
	SharedQueue struct {
		CarsServed     int
		TotalQueueTime time.Duration
	}
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

func updateStats(car Car) {
	defer globalStatsLock.Unlock()

	timeAtEntrance := car.QueueUpForPump.Sub(car.ArrivalTime)
	// fmt.Println(timeAtEntrance.Milliseconds())
	timeInPumpQueue := car.RefuelTime.Sub(car.QueueUpForPump)
	timeOnPump := car.QueueUpForRegister.Sub(car.QueueUpForPump)
	timeInRegisterQueue := car.PayTime.Sub(car.QueueUpForRegister)
	timeOnRegister := car.ExitTime.Sub(car.QueueUpForRegister)

	globalStatsLock.Lock()
	globalStats.SharedQueue.CarsServed++
	globalStats.SharedQueue.TotalQueueTime += timeAtEntrance

	switch car.CarType {
	case Gas:
		globalStats.Pumps.Gas.CarsServed++
		globalStats.Pumps.Gas.TotalQueueTime += timeInPumpQueue
		globalStats.Pumps.Gas.TotalServiceTime += timeOnPump
	case LPG:
		globalStats.Pumps.LPG.CarsServed++
		globalStats.Pumps.LPG.TotalQueueTime += timeInPumpQueue
		globalStats.Pumps.LPG.TotalServiceTime += timeOnPump
	case Electric:
		globalStats.Pumps.Electric.CarsServed++
		globalStats.Pumps.Electric.TotalQueueTime += timeInPumpQueue
		globalStats.Pumps.Electric.TotalServiceTime += timeOnPump
	case Diesel:
		globalStats.Pumps.Diesel.CarsServed++
		globalStats.Pumps.Diesel.TotalQueueTime += timeInPumpQueue
		globalStats.Pumps.Diesel.TotalServiceTime += timeOnPump
	}
	globalStats.Registers.CarsServed++
	globalStats.Registers.TotalQueueTime += timeInRegisterQueue
	globalStats.Registers.TotalServiceTime += timeOnRegister
	globalWG.Done()
}
