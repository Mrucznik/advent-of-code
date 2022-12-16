package old

type Simulation struct {
	choice Choice
	valves map[string]Valve

	currentValve string
	pressure     int
}

func copyValves(valves map[string]Valve) map[string]Valve {
	result := map[string]Valve{}
	for k, v := range valves {
		result[k] = v
	}
	return result
}

func (s *Simulation) calculatePressure() int {
	pressure := 0
	for _, valve := range s.valves {
		if valve.open {
			pressure += valve.rate
		}
	}
	s.pressure = pressure
	return pressure
}

func (s *Simulation) currValePress() int {
	return s.valves[s.currentValve].rate
}
