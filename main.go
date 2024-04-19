package main

import (
	"fmt"
	"gas_station/structs" // assuming "gas_station" is the name of your module
	"math/rand"
	"sync"
	"time"
)

func main() {
	// stations := []*structs.Station{}
	numCars := 10
	wg := &sync.WaitGroup{}
	wg.Add(numCars)
	completionChannel := make(chan bool, numCars)
	carArrivalChannel := make(chan *structs.Car, numCars)
	cars := make([]*structs.Car, numCars)

	gasPumps := structs.NewStation(30, 50, "gas", 3)
	lpgPumps := structs.NewStation(40, 60, "lpg", 2)
	electroPumps := structs.NewStation(20, 40, "electro", 1)

	pumps := []*structs.Station{gasPumps, lpgPumps, electroPumps}

	for i := 0; i < numCars; i++ {
		cars[i] = generateCar(i + 1)
	}

	for _, car := range cars {
		go func(car *structs.Car) {
			time.Sleep(time.Duration(car.Arrival_delay) * time.Millisecond)
			carArrivalChannel <- car // Send car to arrival channel
			wg.Done()
		}(car)
	}

	// Wait for all cars to arrive
	wg.Wait()

	// Process car arrivals
	for i := 0; i < numCars; i++ {
		car := <-carArrivalChannel // Receive car from arrival channel
		pump := structs.GetPumpForCar(car, pumps)
		pump.ProcessCarArrival(car, completionChannel)
	}

	for i := 0; i < numCars; i++ {
		<-completionChannel
	}
	for _, station := range pumps {
		fmt.Println(station.CarsServed, station.TotalQueueTime)
	}
}

// Funkce pro generování náhodného času s daným průměrem a rozptylem
func randomArrivalTime(mean, stddev int) int {
	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)
	return int(generator.NormFloat64()*float64(stddev) + float64(mean))
}

// Funkce pro generování auta s náhodným příjezdovým časem a typem paliva
func generateCar(id int) *structs.Car {
	arrivalTime := randomArrivalTime(1000+id*10, 300) // Průměrný příjezdový čas 1000 ms, rozptyl 300 ms
	return structs.NewCar(arrivalTime, id)
}
