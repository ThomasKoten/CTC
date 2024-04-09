package main

import (
	"fmt"
	"gas_station/structs" // assuming "gas_station" is the name of your module
	"sync"
	"time"
)

func main() {
	stations := []*structs.Station{}
	cars := []*structs.Car{}

	gasPumps := make([]structs.CarQueue, 3)      // 3 gas pumps
	lpgPumps := make([]structs.CarQueue, 2)      // 2 LPG pumps
	electricPumps := make([]structs.CarQueue, 1) // 1 electric pump
	registers := make([]structs.CarQueue, 2)     // 2 registers
	pumpLock := &sync.Mutex{}

	stations = append(stations, structs.NewStation(10, 20, "gas", gasPumps))

	stations = append(stations, structs.NewStation(7, 19, "lpg", lpgPumps))

	stations = append(stations, structs.NewStation(10, 15, "electro", electricPumps))

	stations = append(stations, structs.NewStation(5, 6, "register", registers))

	// Start goroutines to check pumps
	// go structs.CheckPumps(&gasPumps, stations, pumpLock)
	// go structs.CheckPumps(&lpgPumps, stations, pumpLock)
	// go structs.CheckPumps(&electricPumps, stations, pumpLock)
	// sum := 0
	numCars := 10 // Change this to the desired number of cars
	arrivalMin := 2
	arrivalMax := 5
	arrivalDeviation := 0
	wg := &sync.WaitGroup{}
	wg.Add(numCars)
	for i := 0; i < numCars; i++ {
		arrivalTime := structs.GetRandomDuration(arrivalMin, arrivalMax) + arrivalDeviation
		cars = append(cars, structs.NewCar(arrivalTime, i))
		arrivalDeviation = cars[i].Arrival_time
		cars[i].Arrival = time.Now().Add(time.Duration(cars[i].Arrival_time) * time.Millisecond)
		time.Sleep(time.Duration(cars[i].Arrival_time) * time.Millisecond) // Wait for the determined arrival time
		go structs.CarJourney(stations, cars[i], wg, pumpLock)
		// fmt.Println(time.Since(cars[i].Arrival))
		//fmt.Printf("Car %d generated with arrival time: %d seconds\n", cars[i].ID, cars[i].QueueTime)
	}
	wg.Wait()

	for _, car := range cars {
		fmt.Println(car.Car_type, car.QueueTime)
		// fmt.Printf("Nejdelši ceakni %s\n", station.Big)
	}
	// fmt.Println("")
	// for _, station := range stations {
	// 	fmt.Println(station.QueueTime / time.Duration(station.CarsServed))
	// }
	// fmt.Println(time.Duration(sum))

	// select {}
	// for i := 0; i < 10; i++ {
	// 	car := structs.NewCar(2, 5, i)
	// 	cars = append(cars, car) // adjust the parameters as needed
	// 	go structs.CarJourney(stations, car, pumpLock)
	// 	// w := <-wait
	// 	fmt.Printf("Type: %s, Wait: %t\n", cars[i].Car_type, car.Serviced)

	// }

	// Keep the main goroutine alive
}
