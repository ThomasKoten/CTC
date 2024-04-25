package main

//Auto přijede začne tankovat určitou dobu nebo si stoupne do fronty/čeká
//Dokud je někdo v určité frontě bude se přičítat čas
import (
	"math/rand"
	"sync"
	"time"
)

var serviceLock sync.Mutex

func (s *Station) Serve(car *Car) {
	// defer s.queueMutex.Unlock()
	randomCarType()
	servedTime := GetRandomDurationMS(s.Serve_time_min, s.Serve_time_max)
	time.Sleep(servedTime)
	// s.queueMutex.Lock()
}

func FindShortQueue(stations []*Station) int {
	bestStation := 0
	shortestQueue := len(stations[0].WaitQueue)
	if stations[0].Occupied {
		shortestQueue++
	}
	for id := 1; id < len(stations); id++ {
		tempQueue := len(stations[id].WaitQueue)
		if stations[id].Occupied {
			tempQueue++
		}
		if tempQueue < shortestQueue {
			shortestQueue = tempQueue
			bestStation = id
		}
	}
	return bestStation
}

func (sta *Station) fillUpAndGoPay(registers []*Station) {
	for {
		car := <-sta.WaitQueue
		car.RefuelTime = time.Now()
		sta.Occupied = true
		// fmt.Printf("Auto %d tankuje na stanici typu %d. Čekal %s\n", car.ID, sta.Type, car.QueueUpForRegister.Sub(car.RefuelTime))
		sta.Serve(&car)
		serviceLock.Lock()
		queueIndex := FindShortQueue(registers)
		// fmt.Println(queueIndex, len(registers[queueIndex].WaitQueue))
		serviceLock.Unlock()
		car.QueueUpForRegister = time.Now()
		registers[queueIndex].WaitQueue <- car
		sta.Occupied = false
		// fmt.Printf("Auto %d bylo obslouženo na stanici typu %d\n", car.ID, sta.Pump_type)

	}
}

func (station *Station) payAndLeave() {
	for {
		car := <-station.WaitQueue //Occupied and stuff
		station.Occupied = true
		car.PayTime = time.Now()
		payingTime := GetRandomDurationMS(station.Serve_time_min, station.Serve_time_max)
		// fmt.Printf("Car %d about to pay at reg quue %d\n", car.ID, len(station.WaitQueue))
		time.Sleep(payingTime)
		car.ExitTime = time.Now()
		go updateStats(car)
		station.Occupied = false

	}
}

func GetRandomDurationMS(min, max int) time.Duration {
	return time.Duration(rand.Intn(max+1-min)+min) * time.Millisecond
}
