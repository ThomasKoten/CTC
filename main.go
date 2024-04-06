package main

import (
	"fmt"
	"sync"

	"gas_station/structs" // assuming "gas_station" is the name of your module
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

	stations = append(stations, structs.NewStation(5, 8, "lpg", lpgPumps))

	stations = append(stations, structs.NewStation(2, 5, "electro", electricPumps))

	for i := 0; i < 2; i++ {
		stations = append(stations, structs.NewStation(5, 6, "register", registers))
	}

	// Start goroutines to check pumps
	// go structs.CheckPumps(&gasPumps, pumpLock)
	// go structs.CheckPumps(&lpgPumps, pumpLock)
	// go structs.CheckPumps(&electricPumps, pumpLock)

	for i := 0; i < 5; i++ {
		cars = append(cars, structs.NewCar(2, 5, i)) // adjust the parameters as needed
		go structs.CarJourney(stations, cars[i], pumpLock)
		fmt.Printf("Type: %s, Wait: %s \n", cars[i].Car_type)

	}

	// Keep the main goroutine alive
	select {}

	// for _, s := range stations {
	// 	fmt.Printf("Station: %s, Uptime:%d, Downtime: %d\n", s.Station_type, s.Serve_time_min, s.Serve_time_max)
	// }
}
