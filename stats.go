package main

import (
	"fmt"
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

// Prints formated stats output
func printStats() {
	fmt.Println("Pumps:")
	printPumpStats("Gas")
	printPumpStats("LPG")
	printPumpStats("Electric")
	printPumpStats("Diesel")

	fmt.Println("\nRegisters:")
	fmt.Printf("\tCars Served: %d\n", globalStats.Registers.CarsServed)
	fmt.Printf("\tTotal Service Time: %s\n", globalStats.Registers.TotalServiceTime)
	fmt.Printf("\tTotal Queue Time: %s\n", globalStats.Registers.TotalQueueTime)
	fmt.Printf("\tAverage Time: %s\n", globalStats.Registers.AverageTime)
}

// Function for uniform printing of pump stats
func printPumpStats(pumpType string) {
	pump := globalStats.Pumps.Gas
	fmt.Printf("%s:\n", pumpType)
	switch pumpType {

	case "Gas":
		pump = globalStats.Pumps.Gas
	case "LPG":
		pump = globalStats.Pumps.LPG
	case "Electric":
		pump = globalStats.Pumps.Electric
	case "Diesel":
		pump = globalStats.Pumps.Diesel
	}

	fmt.Printf("\tCars Served: %d\n", pump.CarsServed)
	fmt.Printf("\tTotal Service Time: %s\n", pump.TotalServiceTime)
	fmt.Printf("\tTotal Queue Time: %s\n", pump.TotalQueueTime)
	fmt.Printf("\tAverage Time: %s\n", pump.AverageTime)
}
