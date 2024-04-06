package structs

import (
	"math/rand"
	"sync"
	"time"
)

type CarQueue []Car

func CarJourney(stations []*Station, car *Car, pumpLock *sync.Mutex) {
	// Sleep for arrival time
	time.Sleep(time.Millisecond * time.Duration(car.Arrival_time))
	// Record arrival time
	arrival := time.Now()

	// Service at pump station
	servicePump(stations, car, pumpLock)
}

func findAvailablePump(stations []*Station, fuelType string) (int, *Station) {
	// Iterate over stations to find an available pump
	for _, station := range stations {
		// Check if the station matches the fuel type
		if station.Station_type == fuelType {
			// Iterate over pump queues of the station to find an available pump
			for i, queue := range station.Station_queue {
				if len(queue) == 0 { // Found an available pump
					return i, station
				}
			}
		}
	}
	return -1, nil // No available pump found
}

func servicePump(stations []*Station, car *Car, pumpLock *sync.Mutex) {

	// Find an available pump
	pumpIndex, station := findAvailablePump(stations, car.Car_type)

	if pumpIndex == -1 { // No available pump, car goes to queue
		// Enqueue the car into the appropriate queue based on fuel type
		switch car.Car_type {
		case "gas", "lpg", "electric":
			station.Station_queue[0] = append(station.Station_queue[0], *car)
		}
		//car.QueueEnterTime = time.Now()
		return
	}

	serviceTime := getRandomDuration(station.Serve_time_min, station.Serve_time_max)
	time.Sleep(time.Duration(serviceTime) * time.Millisecond)

	switch car.Car_type {
	case "gas", "lpg", "electric":
		removeCarFromQueue(&station.Station_queue, car, pumpIndex)
	}
}

func removeCarFromQueue(queue *[]CarQueue, car *Car, pumpIndex int) {
	queuedCars := (*queue)[pumpIndex]
	for i, queuedCar := range queuedCars {
		if queuedCar.ID == car.ID {
			// Remove car from the queue by slicing the slice
			(*queue)[pumpIndex] = append(queuedCars[:i], queuedCars[i+1:]...)
			break
		}
	}
}
func getRandomDuration(min, max int) int {
	return (rand.Intn((max+1)-min) + min)

}

// func CarFillUp(car *Car) time.Duration {
// 	//Auto přijede začne tankovat určitou dobu nebo si stoupne do fronty/čeká
// 	//Dokud je někdo v určité frontě bude se přičítat čas

// 	waitPump := time.Since(arrival)

// 	return waitPump

// }

// func CheckPumps(pumps *[]CarQueue, pumpLock *sync.Mutex) {
// 	for {
// 		time.Sleep(time.Millisecond * 3) // Check pumps every second

// 		pumpLock.Lock()

// 		// Iterate through pumps and service cars in queues if pump is available
// 		for i, queue := range *pumps {
// 			if len(queue) > 0 { // Queue is not empty
// 				car := queue[0] // Get the first car in the queue
// 				// servicePump(pumps, &car, i, pumpLock)
// 			}
// 		}

// 		pumpLock.Unlock()
// 	}

// }
