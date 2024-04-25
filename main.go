package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type MainQueue struct {
	queue chan Car
}

var globalWG sync.WaitGroup

func main() {
	start := time.Now()
	nCars := 1000
	globalWG.Add(nCars)
	mq := MainQueue{make(chan Car, 10)}
	gs := newGasStation()

	nRegisters := 2
	nGasPumps := 2
	nLPGPumps := 1
	nElectricPumps := 1
	nDieselPumps := 2

	gs.AddRegisters(nRegisters, 1, 3, 1)
	gs.AddFuelPumps(Gas, nGasPumps, 2, 5, 3)
	gs.AddFuelPumps(LPG, nLPGPumps, 4, 7, 3)
	gs.AddFuelPumps(Electric, nElectricPumps, 5, 10, 3)
	gs.AddFuelPumps(Diesel, nDieselPumps, 30, 60, 3)

	go mq.DistributeCars(gs)
	go mq.GenerateCars(nCars, 1, 2)

	// Wait for all cars to leave and the process of stats evaluation
	globalWG.Wait()

	finalizeStats()

	toPrint := fmt.Sprintf("%+v\n", globalStats)
	toPrint = strings.Replace(toPrint, " ", "\n", -1)
	toPrint = strings.Replace(toPrint, ":{", ":{\n", -1)
	toPrint = strings.Replace(toPrint, "}", "}\n", -1)
	fmt.Println(toPrint)
	fmt.Print(time.Since(start))
}
