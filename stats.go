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
	Station struct {
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
}

func updateStats(car Car) {
	defer globalStatsLock.Unlock()

	timeAtEntrance := car.QueueUpForPump.Sub(car.ArrivalTime)
	// fmt.Println(timeAtEntrance.Milliseconds())
	timeInPumpQueue := car.RefuelTime.Sub(car.QueueUpForPump)
	timeOnPump := car.ExitTime.Sub(car.QueueUpForPump)

	globalStatsLock.Lock()
	globalStats.SharedQueue.CarsServed++
	globalStats.SharedQueue.TotalQueueTime += timeAtEntrance

	switch car.CarType {
	case Gas:
		globalStats.Station.Gas.CarsServed++
		globalStats.Station.Gas.TotalQueueTime += timeInPumpQueue
		globalStats.Station.Gas.TotalServiceTime += timeOnPump
	case LPG:
		globalStats.Station.LPG.CarsServed++
		globalStats.Station.LPG.TotalQueueTime += timeInPumpQueue
		globalStats.Station.LPG.TotalServiceTime += timeOnPump
	case Electric:
		globalStats.Station.Electric.CarsServed++
		globalStats.Station.Electric.TotalQueueTime += timeInPumpQueue
		globalStats.Station.Electric.TotalServiceTime += timeOnPump
	case Diesel:
		globalStats.Station.Diesel.CarsServed++
		globalStats.Station.Diesel.TotalQueueTime += timeInPumpQueue
		globalStats.Station.Diesel.TotalServiceTime += timeOnPump
	}
	globalWG.Done()
}
