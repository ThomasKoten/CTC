package main

import (
	"gas_station/structs" // assuming "gas_station" is the name of your module
	"math/rand"
	"time"
)

func main() {
	// stations := []*structs.Station{}
	cars := make([]*structs.Car, 10)

	gasPumps := structs.NewStation(3, 5, "gas", 3)
	lpgPumps := structs.NewStation(4, 6, "lpg", 2)
	electroPumps := structs.NewStation(2, 4, "electro", 1)

	pumps := []*structs.Station{gasPumps, lpgPumps, electroPumps}

	for i := 0; i < 10; i++ {
		cars[i] = generateCar(i + 1)
	}

	for _, car := range cars {
		arrivalTime := time.Duration(car.Arrival_time) * time.Millisecond
		time.Sleep(arrivalTime)
		pump := structs.GetPumpForCar(car, pumps)
		pump.ProcessCarArrival(car)
	}
}

// Funkce pro generování náhodného času s daným průměrem a rozptylem
func randomArrivalTime(mean, stddev int) int {
	source := rand.NewSource(time.Now().UnixNano())
	generator := rand.New(source)
	return int(generator.NormFloat64()*float64(stddev) + float64(mean))
}

// Funkce pro generování auta s náhodným příjezdovým časem a typem paliva
func generateCar(id int) *structs.Car {
	arrivalTime := randomArrivalTime(1000, 300) // Průměrný příjezdový čas 1000 ms, rozptyl 300 ms
	return structs.NewCar(arrivalTime, id)
}
