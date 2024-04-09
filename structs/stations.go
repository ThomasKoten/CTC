// station.go
package structs

import "time"

type Station struct {
	Serve_time_min int
	Serve_time_max int
	Station_type   string
	Station_queue  []CarQueue
	QueueTime      time.Duration
	CarsServed     int
	Big            time.Duration
}

// Function to create a new Station
func NewStation(serveTimeMin, serveTimeMax int, stationType string, queue []CarQueue) *Station {
	return &Station{
		Serve_time_min: serveTimeMin,
		Serve_time_max: serveTimeMax,
		Station_type:   stationType,
		Station_queue:  queue,
		CarsServed:     1,
		Big:            0,
	}
}

//Funkce vygeneruje auto, pošle ho natankovat nebo do fronty, pošle ho na kasu/frontu, uvolní thread
