// station.go
package structs

type Station struct {
	Serve_time_min int
	Serve_time_max int
	Station_type   string
	Station_queue  []CarQueue
}

// Function to create a new Station
func NewStation(serveTimeMin, serveTimeMax int, stationType string, queue []CarQueue) *Station {
	return &Station{
		Serve_time_min: serveTimeMin,
		Serve_time_max: serveTimeMax,
		Station_type:   stationType,
		Station_queue:  queue,
	}
}

//Funkce vygeneruje auto, pošle ho natankovat nebo do fronty, pošle ho na kasu/frontu, uvolní thread
