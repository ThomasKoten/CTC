package structs

//Auto přijede začne tankovat určitou dobu nebo si stoupne do fronty/čeká
// 	//Dokud je někdo v určité frontě bude se přičítat čas
import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type CarQueue []Car

var stationsMutex sync.Mutex

func CarJourney(stations []*Station, car *Car, wg *sync.WaitGroup, pumpLock *sync.Mutex) {
	defer wg.Done()
	// Sleep for arrival time

	// Service at pump station
	for !car.Serviced {
		servicePump(stations, car, pumpLock)
	}

}

func servicePump(stations []*Station, car *Car, pumpLock *sync.Mutex) {
	pumpLock.Lock()
	defer pumpLock.Unlock()
	sLock := &sync.Mutex{}
	// Find an available pump
	sLock.Lock()
	pumpIndex, station := findAvailablePump(stations, car.Car_type)
	sLock.Unlock()
	if pumpIndex == -1 { // No available pump, car goes to queue
		// Enqueue the car into the appropriate queue based on fuel type
		return
	}
	sLock.Lock() //At one point there was index one T_T
	station.Station_queue[pumpIndex] = append(station.Station_queue[pumpIndex], *car)
	// sLock.Unlock()
	fmt.Println(pumpIndex, station.Station_queue)
	queueExitTime := time.Now()

	// Calculate the queue time
	queueTime := queueExitTime.Sub(car.Arrival)

	// Accumulate the queue time for the car
	car.QueueTime = queueTime
	// fmt.Printf("Car: %s Time:%s\n", car.Car_type, car.QueueTime)

	// Accumulate the queue time for the station
	station.QueueTime += queueTime
	// station.QueueTime += end
	station.CarsServed++
	if queueTime > station.Big {
		station.Big = queueTime
		// fmt.Printf("Sattion: %s Big: %s \n", station.Station_type, station.Big)
	}
	serviceTime := GetRandomDuration(station.Serve_time_min, station.Serve_time_max)
	time.Sleep(time.Duration(serviceTime) * time.Millisecond)
	car.Serviced = true

	switch car.Car_type {
	case "gas", "lpg", "electric":
		removeCarFromQueue(&station.Station_queue, car, pumpIndex)
	}
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
func GetRandomDuration(min, max int) int {
	return (rand.Intn((max+1)-min) + min)

}

func CheckPumps(pumps *[]CarQueue, stations []*Station, pumpLock *sync.Mutex) {
	for {
		time.Sleep(time.Millisecond * 3) // Check pumps every second

		pumpLock.Lock()

		// Iterate through pumps and service cars in queues if pump is available
		for _, queue := range *pumps {
			if len(queue) > 0 { // Queue is not empty
				car := queue[0] // Get the first car in the queue
				servicePump(stations, &car, pumpLock)
			}
		}

		pumpLock.Unlock()
	}

}
