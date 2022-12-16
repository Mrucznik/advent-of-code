package main

import (
	_ "embed"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

var valves = map[string]Valve{}

func main() {
	rows := strings.Split(input, "\n")

	for _, row := range rows {
		re := regexp.MustCompile("Valve (..) has flow rate=(\\d+); tunnels? leads? to valves? ((?:[A-Z]{2}[ ,]{0,2})+)")
		raw := re.FindStringSubmatch(row)

		valve := raw[1]
		flowRate, _ := strconv.Atoi(raw[2])
		connections := strings.Split(raw[3], ", ")

		valves[valve] = Valve{valve, flowRate, false, connections}
	}

	// 1. znajdź ścieżki, które przechodzą przez wszystkie node'y
	// 2. weź tylko takie ścieżki, które mają długość < 30

	// tylko 12 node'ów, więc można pomijać drogę

	// z każdym krokiem możemy wejść w mniejszą liczbę ścieżek, które są dla nas dobre

	// priorytetyzuj najlepsze

	// trzeba policzyć priorytety poszczególnych node'ów
	// by znać priorytet, trzeba mieć:
	// 1. najkrótszą ścieżkę
	// 2.

	// to musi być zagładzanie
	// w każdej rundzie, liczyć, ile możnaby zdobyć punktów, gdyby otworzyć dany zawór

	// good valves
	var goodValves []string
	for key, valve := range valves {
		if valve.rate > 0 {
			goodValves = append(goodValves, key)
		}
	}

	// znajdź liczbe punktów
	for i := 0; i < 30; i++ {
		maxPointsToGain = calculatePossiblePoints(valves, goodValves)
	}

	// znajdź koszt z AA do podanych ścieżek, nie więcej niż 30
	//simulations := []*Simulation{{
	//	currentValve: "AA",
	//	valves:       copyValves(valves),
	//}}
	//
	//// i can move / open tunnel
	//for i := 0; i < 12; i++ {
	//	// dodajemy nowe symulacje i ich wybory
	//	var newSims []*Simulation
	//	for _, simulation := range simulations {
	//		if simulation.choice != nil {
	//			simulation.choice.Do(simulation)
	//		}
	//
	//		// open simulation
	//		if simulation.currValePress() > 0 {
	//			openSim := Simulation{
	//				choice:       &Open{what: simulation.currentValve},
	//				valves:       copyValves(simulation.valves),
	//				currentValve: simulation.currentValve,
	//				pressure:     0,
	//			}
	//			newSims = append(newSims, &openSim)
	//		}
	//
	//		// move simulations
	//		for _, connection := range valves[simulation.currentValve].connections {
	//			choiceSim := Simulation{
	//				choice:       &Move{connection},
	//				valves:       copyValves(simulation.valves),
	//				currentValve: simulation.currentValve,
	//			}
	//			newSims = append(newSims, &choiceSim)
	//		}
	//	}
	//
	//	simulations = newSims
	//}
	//
	//for _, simulation := range simulations {
	//	simulation.choice.Do(simulation)
	//	simulation.calculatePressure()
	//}
	//sort.Slice(simulations, func(i, j int) bool {
	//	return simulations[i].pressure > simulations[j].pressure
	//})
	//
	//for i, simulation := range simulations[:10] {
	//	fmt.Printf("simulation %d: pressure %d\n", i, simulation.pressure)
	//}
	//
	//fmt.Println(len(simulations))
}

func calculatePossiblePoints(valves map[string]Valve, goodValves []string, moves int) int {
	max := 0
	for _, valve := range goodValves {
		if valves[valve].open {
			rate := valves[valve].rate
			if max > rate {
				max = rate
			}
		}
	}
	return max
}

func znajdzKosztDrogi(cofnij string, begin string, end string, limit int) (int, []string) {
	if limit == 0 {
		return 99999
	}

	for _, v := range valves[begin].connections {
		if v == cofnij {
			continue
		}
		if v == end {
			return 1
		}
	}

	koszty := []int{}
	for _, v := range valves[begin].connections {
		if v == cofnij {
			continue
		}
		koszty = append(koszty, znajdzKosztDrogi(begin, v, end, limit-1))
	}

	min := 99999999999999
	for _, koszt := range koszty {
		if koszt < min {
			min = koszt
		}
	}
	return min
}
