// station.go
package structs

import (
	"sync"
	"time"
)

type Station struct {
	Serve_time_min int
	Serve_time_max int
	Pump_type      string
	Pump_count     int
	ServicePlace   []*Car
	WaitQueue      []*Car
	CarsServed     int
	TotalQueueTime time.Duration
	queueMutex     sync.Mutex
}

// Function to create a new Station
func NewStation(serveTimeMin, serveTimeMax int, pumpType string, pumpCount int) *Station {
	return &Station{
		Serve_time_min: serveTimeMin,
		Serve_time_max: serveTimeMax,
		Pump_type:      pumpType,
		Pump_count:     pumpCount,
		ServicePlace:   make([]*Car, 0),
		WaitQueue:      make([]*Car, 0),
		CarsServed:     0,
		TotalQueueTime: 0,
	}
}

// type Station struct {
// 	Serve_time_min int
// 	Serve_time_max int
// 	Station_type   string
// 	Station_queue  []CarQueue
// 	QueueTime      time.Duration
// 	CarsServed     int
// 	Big            time.Duration
// }

// // Function to create a new Station
// func NewStation(serveTimeMin, serveTimeMax int, stationType string, queue []CarQueue) *Station {
// 	return &Station{
// 		Serve_time_min: serveTimeMin,
// 		Serve_time_max: serveTimeMax,
// 		Station_type:   stationType,
// 		Station_queue:  queue,
// 		CarsServed:     1,
// 		Big:            0,
// 	}
// }

//Funkce vygeneruje auto, pošle ho natankovat nebo do fronty, pošle ho na kasu/frontu, uvolní thread
