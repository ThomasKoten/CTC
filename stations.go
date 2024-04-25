package main

const (
	Gas = iota
	LPG
	Electric
	Diesel
)

type Station struct {
	Serve_time_min int
	Serve_time_max int
	Type           int
	// Pump_count     int
	// ServicePlace   []*Car
	WaitQueue chan Car
	Occupied  bool
}

type GasStation struct {
	Pumps     map[int][]*Station
	Registers []*Station
}

// Function to create a new Station
func NewPump(serveTimeMin, serveTimeMax int, fuelType int, maxQueueLength int) *Station {
	return &Station{
		Serve_time_min: serveTimeMin,
		Serve_time_max: serveTimeMax,
		Type:           fuelType,
		WaitQueue:      make(chan Car, maxQueueLength),
		Occupied:       false,
	}
}

func NewRegister(serveTimeMin, serveTimeMax int, maxQueueLength int) *Station {
	return &Station{
		Serve_time_min: serveTimeMin,
		Serve_time_max: serveTimeMax,
		Type:           8,
		WaitQueue:      make(chan Car, maxQueueLength),
		Occupied:       false,
	}
}

func (gs *GasStation) AddPumps(pumpType int, nPumps, minTime, maxTime, maxQueue int) {
	gs.Pumps[pumpType] = make([]*Station, nPumps)
	for i := 0; i < nPumps; i++ {
		gs.Pumps[pumpType][i] = NewPump(minTime, maxTime, pumpType, maxQueue)
		go gs.Pumps[pumpType][i].fillUpAndGoPay(gs.Registers)
	}
}

func (gs *GasStation) AddRegisters(nRegisters, minTime, maxTime, maxQueue int) {
	gs.Registers = make([]*Station, nRegisters)
	for i := 0; i < nRegisters; i++ {
		gs.Registers[i] = NewRegister(minTime, maxTime, maxQueue)
		go gs.Registers[i].payAndLeave()
	}
}
