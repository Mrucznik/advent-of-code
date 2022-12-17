package main

import (
	"bytes"
	"encoding/binary"
	"github.com/cespare/xxhash"
	"sort"
)

type Simulation struct {
	currentValve    int8
	currentElephant int8
	pressure        int16

	opened map[int8]struct{}
}

func (s *Simulation) Hash() uint64 {
	sb := xxhash.New()
	buff := new(bytes.Buffer)

	binary.Write(buff, binary.BigEndian, int64(s.pressure))

	if len(s.opened) == len(s.getBestValves()) {
		sb.Write(buff.Bytes())
		return sb.Sum64()
	}

	for i := range s.opened {
		sb.Write([]byte(valves[i].name))
	}
	binary.Write(buff, binary.BigEndian, int64(s.currentValve))
	binary.Write(buff, binary.BigEndian, int64(s.currentElephant))

	sb.Write(buff.Bytes())
	return sb.Sum64()
}

func (s *Simulation) CollectPressure() {
	opened := s.getOpenValves()
	if len(opened) == 0 {
		return
	}

	gain := int16(0)
	//fmt.Print("Valves ")
	for _, valve := range opened {
		gain += valve.rate
	}
	s.pressure += gain
	//fmt.Printf(" pressure %d\n", s.pressure)
}

func (s *Simulation) NextChoice(choice, elephantChoice int) *Simulation {
	newOpened := make(map[int8]struct{}, len(s.opened))
	for k, v := range s.opened {
		newOpened[k] = v
	}

	newSim := &Simulation{
		opened:          newOpened,
		pressure:        s.pressure,
		currentValve:    -1,
		currentElephant: -1,
	}
	if choice >= len(s.curr().connections) {
		// open current valve
		for i, valve := range valves {
			if s.curr().name == valve.name {
				newSim.currentValve = int8(i)
				newSim.opened[int8(i)] = struct{}{}
			}
		}

		//fmt.Printf("opening %s\n", newSim.currentValve.name)
	} else {
		// new choice
		chosenConnection := s.curr().connections[choice]
		for i, valve := range valves {
			if valve.name == chosenConnection {
				newSim.currentValve = int8(i)
				break
			}
		}
		//fmt.Printf("moving to %s\n", newSim.currentValve.name)
	}

	// elephant
	if elephantChoice >= len(s.currElephant().connections) {
		// open current valve
		for i, valve := range valves {
			if s.currElephant().name == valve.name {
				newSim.currentElephant = int8(i)
				newSim.opened[int8(i)] = struct{}{}
			}
		}

		//fmt.Printf("opening %s\n", newSim.currentValve.name)
	} else {
		// new choice
		chosenConnection := s.currElephant().connections[elephantChoice]
		for i, valve := range valves {
			if valve.name == chosenConnection {
				newSim.currentElephant = int8(i)
				break
			}
		}
		//fmt.Printf("moving to %s\n", newSim.currentValve.name)
	}

	return newSim
}

func (s *Simulation) getPossibleChoicesNumber() int {
	if len(s.opened) == len(s.getBestValves()) {
		return 1
	}

	if s.isOpen(s.currentValve) || valves[s.currentValve].rate == 0 {
		return len(s.curr().connections)
	} else {
		return len(s.curr().connections) + 1
	}
}

func (s *Simulation) getPossibleElephantChoicesNumber() int {
	if len(s.opened) == len(s.getBestValves()) {
		return 1
	}

	if s.isOpen(s.currentElephant) || valves[s.currentElephant].rate == 0 {
		return len(s.currElephant().connections)
	} else {
		return len(s.currElephant().connections) + 1
	}
}

func (s *Simulation) getBestValves() []*Valve {
	var result []*Valve
	for i, valve := range valves {
		if !s.isOpen(int8(i)) && valve.rate > 0 {
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

func (s *Simulation) GetPossibleGain(stepsToEnd int8) int16 {
	gain := s.getOpenedValvesGain(stepsToEnd)

	return gain + s.naivePossibleGain(stepsToEnd) + s.pressure
}

func (s *Simulation) getOpenedValvesGain(stepsToEnd int8) int16 {
	var gain int16
	for _, valve := range s.getOpenValves() {
		gain += valve.rate * int16(stepsToEnd)
	}
	return gain
}

func (s *Simulation) naivePossibleGain(stepsToEnd int8) int16 {
	possibleGain := int16(0)
	turns := int8(2)
	for _, valve := range s.getBestValves() {
		stepsLeft := stepsToEnd - turns
		if stepsLeft <= 0 {
			break
		}
		possibleGain += valve.rate * int16(stepsLeft)
		turns += 2
	}
	return possibleGain
}

func (s *Simulation) getGain() int16 {
	return s.getOpenedValvesGain(1)
}

func (s *Simulation) isOpen(i int8) bool {
	_, ok := s.opened[i]
	return ok
}

func (s *Simulation) curr() *Valve {
	return valves[s.currentValve]
}

func (s *Simulation) currElephant() *Valve {
	return valves[s.currentElephant]
}
