package structs

import (
	"math/rand"
)

type Car struct {
	Arrival_time int
	Car_type     string
	ID           int
}

func randomCarType() string {
	var CAR_TYPES = [3]string{"gas", "lpg", "electro"}
	return CAR_TYPES[rand.Intn(3)]
}
func NewCar(arrivalTimeMin, arrivalTimeMax, id int) *Car {
	return &Car{
		Arrival_time: getRandomDuration(arrivalTimeMin, arrivalTimeMax),
		Car_type:     randomCarType(),
	}
}

func NewCarAll(arrivalTimeMin, arrivalTimeMax, id int, carType string) *Car {
	return &Car{
		Arrival_time: getRandomDuration(arrivalTimeMin, arrivalTimeMax),
		Car_type:     carType,
	}
}
