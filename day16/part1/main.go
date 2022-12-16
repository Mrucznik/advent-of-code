package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// time needed: ~4h

//go:embed input.txt
var input string

type Valve struct {
	name        string
	rate        int
	connections []string
}

var valves []*Valve

func main() {
	rows := strings.Split(input, "\n")
	var beginValve int

	for i, row := range rows {
		re := regexp.MustCompile("Valve (..) has flow rate=(\\d+); tunnels? leads? to valves? ((?:[A-Z]{2}[ ,]{0,2})+)")
		raw := re.FindStringSubmatch(row)

		valve := raw[1]
		flowRate, _ := strconv.Atoi(raw[2])
		connections := strings.Split(raw[3], ", ")

		valves = append(valves, &Valve{valve, flowRate, connections})
		if valve == "AA" {
			beginValve = i
		}
	}

	// odpalamy symulacje
	var simulations []*Simulation
	simulations = append(simulations, &Simulation{
		opened:       map[int]struct{}{},
		currentValve: beginValve,
	})

	steps := 30
	for i := 0; i < steps; i++ {
		fmt.Printf("step %d, simulation counts: %d\n", i+1, len(simulations))
		for _, simulation := range simulations {
			// przeprowadź symulacje
			simulation.CollectPressure()
		}

		// find best simulation
		var maxSim *Simulation
		for _, simulation := range simulations {
			if maxSim == nil || simulation.pressure > maxSim.pressure {
				maxSim = simulation
			}
		}

		// głodzenie i umieranie
		newSims := make(map[string]*Simulation, len(simulations)*8)
		for _, simulation := range simulations {
			possGain := simulation.GetPossibleGain(steps - i)
			if possGain < maxSim.pressure && i > 2 {
				// die
				//fmt.Printf("i'm dying, possibble gain: %d, max pressure: %d\n", possGain, maxSim.pressure)
			} else {
				// create new choices
				for choice := 0; choice < simulation.getPossibleChoicesNumber(); choice++ {
					newSim := simulation.NextChoice(choice)
					newSims[newSim.Hash()] = newSim
				}
			}
		}

		result := make([]*Simulation, 0, len(newSims))
		for _, simulation := range newSims {
			result = append(result, simulation)
		}

		// replace sims, next step
		simulations = result
	}

	bestPressure := 0
	for _, simulation := range simulations {
		if simulation.pressure > bestPressure {
			bestPressure = simulation.pressure
		}
	}

	fmt.Printf("best pressure: %d\n", bestPressure)

}
