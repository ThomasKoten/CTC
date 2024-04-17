package structs

import (
	"math/rand"
)

type Car struct {
	ID           int
	Arrival_time int
	Car_type     string
}

func randomCarType() string {
	var CAR_TYPES = [3]string{"gas", "lpg", "electro"}
	return CAR_TYPES[rand.Intn(3)]
}
func NewCar(arrivalTime, id int) *Car {
	return &Car{
		ID:           id,
		Arrival_time: arrivalTime,
		Car_type:     randomCarType(),
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
