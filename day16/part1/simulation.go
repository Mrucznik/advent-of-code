package main

import (
	"sort"
	"strconv"
	"strings"
)

type Choices int

type Simulation struct {
	valves []Valve

	currentValve *Valve
	pressure     int
}

func (s *Simulation) Hash() string {
	sb := strings.Builder{}
	for i, valve := range s.valves {
		if valve.open {
			sb.WriteString(strconv.Itoa(i))
		}
	}
	sb.WriteString("_")
	sb.WriteString(strconv.Itoa(s.pressure))
	sb.WriteString(s.currentValve.name)
	return sb.String()
}

func (s *Simulation) CollectPressure() {
	opened := s.getOpenValves()
	if len(opened) == 0 {
		return
	}

	gain := 0
	//fmt.Print("Valves ")
	for _, valve := range opened {
		gain += valve.rate
	}
	s.pressure += gain
	//fmt.Printf(" pressure %d\n", s.pressure)
}

func (s *Simulation) NextChoice(choice int) *Simulation {
	newValves := make([]Valve, len(s.valves))
	for i := 0; i < len(s.valves); i++ {
		cpy := s.valves[i]
		newValves[i] = cpy
	}

	newSim := &Simulation{
		valves:       newValves,
		pressure:     s.pressure,
		currentValve: nil,
	}
	if choice >= len(s.currentValve.connections) {
		// open current valve
		for i, valve := range s.valves {
			if s.currentValve.name == valve.name {
				newSim.currentValve = &newSim.valves[i]
			}
		}

		newSim.currentValve.open = true
		//fmt.Printf("opening %s\n", newSim.currentValve.name)
	} else {
		// new choice
		chosenConnection := s.currentValve.connections[choice]
		var chosen *Valve
		for i, valve := range newValves {
			if valve.name == chosenConnection {
				chosen = &newValves[i]
			}
		}

		newSim.currentValve = chosen
		//fmt.Printf("moving to %s\n", newSim.currentValve.name)
	}

	return newSim
}

func (s *Simulation) getPossibleChoicesNumber() int {
	if s.currentValve.open {
		return len(s.currentValve.connections)
	} else {
		return len(s.currentValve.connections) + 1
	}
}

func (s *Simulation) getBestValves() []Valve {
	var result []Valve
	for _, valve := range s.valves {
		if !valve.open && valve.rate > 0 {
			result = append(result, valve)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].rate > result[j].rate
	})
	return result
}

func (s *Simulation) getOpenValves() []Valve {
	var result []Valve
	for _, valve := range s.valves {
		if valve.open && valve.rate > 0 {
			result = append(result, valve)
		}
	}
	return result
}

func (s *Simulation) GetPossibleGain(stepsToEnd int) int {
	gain := s.getOpenedValvesGain(stepsToEnd)

	return gain + s.naivePossibleGain(stepsToEnd) + s.pressure
}

func (s *Simulation) getOpenedValvesGain(stepsToEnd int) int {
	gain := 0
	for _, valve := range s.getOpenValves() {
		gain += valve.rate * stepsToEnd
	}
	return gain
}

func (s *Simulation) naivePossibleGain(stepsToEnd int) int {
	possibleGain := 0
	turns := 2
	for _, valve := range s.getBestValves() {
		stepsLeft := stepsToEnd - turns
		if stepsLeft <= 0 {
			break
		}
		possibleGain += valve.rate * stepsLeft
		turns += 2
	}
	return possibleGain
}

func (s *Simulation) getGain() int {
	return s.getOpenedValvesGain(1)
}
