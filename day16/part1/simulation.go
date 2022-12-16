package main

import (
	"sort"
	"strconv"
	"strings"
)

type Simulation struct {
	currentValve int
	pressure     int

	opened map[int]struct{}
}

func (s *Simulation) Hash() string {
	sb := strings.Builder{}
	for i := range s.opened {
		sb.WriteString(strconv.Itoa(i))
	}
	sb.WriteString("_")
	sb.WriteString(strconv.Itoa(s.pressure))
	sb.WriteString(valves[s.currentValve].name)
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
	newOpened := make(map[int]struct{}, len(s.opened))
	for k, v := range s.opened {
		newOpened[k] = v
	}

	newSim := &Simulation{
		opened:       newOpened,
		pressure:     s.pressure,
		currentValve: -1,
	}
	if choice >= len(s.curr().connections) {
		// open current valve
		for i, valve := range valves {
			if s.curr().name == valve.name {
				newSim.currentValve = i
				newSim.opened[i] = struct{}{}
			}
		}

		//fmt.Printf("opening %s\n", newSim.currentValve.name)
	} else {
		// new choice
		chosenConnection := s.curr().connections[choice]
		for i, valve := range valves {
			if valve.name == chosenConnection {
				newSim.currentValve = i
				break
			}
		}
		//fmt.Printf("moving to %s\n", newSim.currentValve.name)
	}

	return newSim
}

func (s *Simulation) getPossibleChoicesNumber() int {
	if s.isOpen(s.currentValve) || valves[s.currentValve].rate == 0 {
		return len(s.curr().connections)
	} else {
		return len(s.curr().connections) + 1
	}
}

func (s *Simulation) getBestValves() []*Valve {
	var result []*Valve
	for i, valve := range valves {
		if !s.isOpen(i) && valve.rate > 0 {
			result = append(result, valve)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].rate > result[j].rate
	})
	return result
}

func (s *Simulation) getOpenValves() []*Valve {
	var result []*Valve
	for i := range s.opened {
		valve := valves[i]
		if valve.rate > 0 {
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

func (s *Simulation) isOpen(i int) bool {
	_, ok := s.opened[i]
	return ok
}

func (s *Simulation) curr() *Valve {
	return valves[s.currentValve]
}
