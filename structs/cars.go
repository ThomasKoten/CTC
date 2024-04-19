package structs

import (
	"math/rand"
	"time"
)

type Car struct {
	ID            int
	Arrival_delay int
	Arrival_time  time.Time
	Car_type      string
	Queued        bool
}

func randomCarType() string {
	var CAR_TYPES = [3]string{"gas", "lpg", "electro"}
	return CAR_TYPES[rand.Intn(3)]
}
func NewCar(delay, id int) *Car {
	return &Car{
		ID:            id,
		Arrival_delay: delay,
		Car_type:      "gas",
		Queued:        false,
	}
}

// type Car struct {
// 	ID           int
// 	Arrival_time int
// 	Arrival      time.Time
// 	Car_type     string
// 	Serviced     bool
// 	QueueTime    time.Duration
// }

// func randomCarType() string {
// 	var CAR_TYPES = [3]string{"gas", "lpg", "electro"}
// 	return CAR_TYPES[rand.Intn(3)]
// }
// func NewCar(arrivalTime, id int) *Car {
// 	return &Car{
// 		ID:           id,
// 		Arrival_time: arrivalTime,
// 		Car_type:     randomCarType(),
// 		Serviced:     false,
// 	}
// }

// func NewCarAll(arrivalTimeMin, arrivalTimeMax, id int, carType string) *Car {
// 	return &Car{
// 		ID:           id,
// 		Arrival_time: GetRandomDuration(arrivalTimeMin, arrivalTimeMax),
// 		Car_type:     carType,
// 		Serviced:     false,
// 	}
// }
