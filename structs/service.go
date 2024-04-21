package main

//Auto přijede začne tankovat určitou dobu nebo si stoupne do fronty/čeká
//Dokud je někdo v určité frontě bude se přičítat čas
import (
	"fmt"
	"math/rand"
	"time"
)

func (s *Station) Serve(car *Car) {
	defer s.queueMutex.Unlock()
	servedTime := time.Duration(GetRandomDuration(s.Serve_time_min, s.Serve_time_max))
	time.Sleep(servedTime * time.Millisecond)
	s.queueMutex.Lock()
}

func FindShortQueue(pumps []*Station) int {
	bestPump := 0
	shortestQueue := len(pumps[0].WaitQueue)
	for id := 1; id < len(pumps); id++ {
		tempQueue := len(pumps[id].WaitQueue)
		fmt.Println(tempQueue)
		if tempQueue < shortestQueue {
			shortestQueue = tempQueue
			bestPump = id
		}
	}

	return bestPump
}

func RemoveIndex(queue []*Car, carID int) []*Car {
	ret := make([]*Car, 0)
	for index, car := range queue {
		if carID == car.ID {
			ret = append(ret, queue[:index]...)
			return append(ret, queue[index+1:]...)
		}

	}
	fmt.Printf("Auto %d nenalezeno.\n", carID)
	return nil

}

func (sta *Station) fillUpAndLeave() {
	for {
		car := <-sta.WaitQueue
		car.RefuelTime = time.Now()
		// fmt.Printf("Auto %d tankuje na stanici typu %d. Čekal %s\n", car.ID, sta.Pump_type, car.RefuelTime)
		sta.Serve(&car)
		car.ExitTime = time.Now()
		go updateStats(car)
		// fmt.Printf("Auto %d bylo obslouženo na stanici typu %d\n", car.ID, sta.Pump_type)

	}
}

func GetRandomDuration(min, max int) int {
	return (rand.Intn((max+1)-min) + min)
}
