package main

type FuelType int

const (
	Gas FuelType = iota + 1
	LPG
	Electric
	Diesel
)

type Station struct {
	ServeTimeMin int
	ServeTimeMax int
	Type         FuelType
	WaitQueue    chan Car
	Occupied     bool
}

type GasStation struct {
	Pumps     map[FuelType][]*Station
	Registers []*Station
}

// Function to create a new GasStation
func newGasStation() *GasStation {
	return &GasStation{Pumps: make(map[FuelType][]*Station)}
}

// Function to create a new Pump
func newFuelPump(serveTimeMin, serveTimeMax int, fuelType FuelType, maxQueueLength int) *Station {
	return &Station{
		ServeTimeMin: serveTimeMin,
		ServeTimeMax: serveTimeMax,
		Type:         fuelType,
		WaitQueue:    make(chan Car, maxQueueLength), // Channel for cars waiting at the pump
		Occupied:     false,
	}
}

// Function to create a new Register
func newRegister(serveTimeMin, serveTimeMax int, maxQueueLength int) *Station {
	return &Station{
		ServeTimeMin: serveTimeMin,
		ServeTimeMax: serveTimeMax,
		WaitQueue:    make(chan Car, maxQueueLength), // Channel for cars waiting at the register
		Occupied:     false,
	}
}

// Add pumps to a GasStation
func (gs *GasStation) AddFuelPumps(pumpType FuelType, nPumps, minTime, maxTime, maxQueue int) {
	gs.Pumps[pumpType] = make([]*Station, nPumps)
	for i := 0; i < nPumps; i++ {
		gs.Pumps[pumpType][i] = newFuelPump(minTime, maxTime, pumpType, maxQueue)
		go gs.Pumps[pumpType][i].refuelAndGoPay(gs.Registers)
	}
}

// Add registers to a GasStation
func (gs *GasStation) AddRegisters(nRegisters, minTime, maxTime, maxQueue int) {
	gs.Registers = make([]*Station, nRegisters)
	for i := 0; i < nRegisters; i++ {
		gs.Registers[i] = newRegister(minTime, maxTime, maxQueue)
		go gs.Registers[i].payAndLeave()
	}
}
