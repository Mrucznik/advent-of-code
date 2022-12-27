package main

import (
	_ "embed"
	"fmt"
	"github.com/cespare/xxhash"
	"sort"
)

// time needed part 1: 2h
// time needed part 2: 12min

type Decision int

const (
	WAIT Decision = iota
	ORE
	CLAY
	OBSIDIAN
	GEODE
)

//go:embed input.txt
var input string

type Cost struct {
	ore, clay, obsidian byte
}

func (c Cost) afford(ore, clay, obsidian byte) bool {
	return c.ore <= ore && c.clay <= clay && c.obsidian <= obsidian
}

type Blueprint struct {
	oreRobot      Cost
	clayRobot     Cost
	obsidianRobot Cost
	geodeRobot    Cost
}

type Simulation struct {
	blueprint *Blueprint

	oreRobots      byte
	clayRobots     byte
	obsidianRobots byte
	geodeRobots    byte

	ore, clay, obsidian, geodes byte

	decision Decision
}

func (s *Simulation) hash() uint64 {
	h := xxhash.New()
	_, err := h.Write([]byte{s.oreRobots, s.clayRobots, s.obsidianRobots, s.geodeRobots, s.ore, s.clay, s.obsidian, s.geodes})
	if err != nil {
		panic(err)
	}
	return h.Sum64()
}

func NewSimulation(blueprint *Blueprint) *Simulation {
	return &Simulation{blueprint: blueprint, oreRobots: 1}
}

func (s *Simulation) decide() []*Simulation {
	decisions := []*Simulation{}

	waitSim := *s
	waitSim.decision = WAIT
	decisions = append(decisions, &waitSim)

	if s.blueprint.geodeRobot.afford(s.ore, s.clay, s.obsidian) {
		d := *s
		d.decision = GEODE
		decisions = append(decisions, &d)
	}
	if s.blueprint.obsidianRobot.afford(s.ore, s.clay, s.obsidian) {
		d := *s
		d.decision = OBSIDIAN
		decisions = append(decisions, &d)
	}
	if s.blueprint.clayRobot.afford(s.ore, s.clay, s.obsidian) {
		d := *s
		d.decision = CLAY
		decisions = append(decisions, &d)
	}
	if s.blueprint.oreRobot.afford(s.ore, s.clay, s.obsidian) {
		d := *s
		d.decision = ORE
		decisions = append(decisions, &d)
	}

	return decisions
}

func (s *Simulation) tick() {
	// after decision
	s.collect()
	s.build()
}

func (s *Simulation) collect() {
	s.ore += s.oreRobots
	s.clay += s.clayRobots
	s.obsidian += s.obsidianRobots
	s.geodes += s.geodeRobots
}

func (s *Simulation) build() {
	switch s.decision {
	case ORE:
		s.removeMaterials(s.blueprint.oreRobot)
		s.oreRobots++
	case CLAY:
		s.removeMaterials(s.blueprint.clayRobot)
		s.clayRobots++
	case OBSIDIAN:
		s.removeMaterials(s.blueprint.obsidianRobot)
		s.obsidianRobots++
	case GEODE:
		s.removeMaterials(s.blueprint.geodeRobot)
		s.geodeRobots++
	}
}

func (s *Simulation) removeMaterials(cost Cost) {
	s.ore -= cost.ore
	s.clay -= cost.clay
	s.obsidian -= cost.obsidian
}

// we can construct, or not construct

func main() {
	result := 1
	for _, blueprint := range mainBlueprints[:3] {
		bpResult := evaluateBlueprint(blueprint)
		fmt.Println("---- bpResult", bpResult)
		result *= bpResult
	}

	fmt.Println(result)
}

func evaluateBlueprint(blueprint *Blueprint) int {
	fmt.Println("evaluate blueprint", blueprint)
	simulations := map[uint64]*Simulation{}
	beginSim := NewSimulation(blueprint)

	simulations[beginSim.hash()] = beginSim
	for i := 0; i < 32; i++ {
		fmt.Println("step", i)

		// create simulations based on available decisions
		sims := []*Simulation{}
		for _, simulation := range simulations {
			for _, newSim := range simulation.decide() {
				sims = append(sims, newSim)
			}
		}

		// tick & remove duplicates
		deduplicated := map[uint64]*Simulation{}
		for _, simulation := range sims {
			simulation.tick()
			deduplicated[simulation.hash()] = simulation
		}

		// find potential gain
		minimalPotentialGain := int(0)
		for _, simulation := range deduplicated {
			potentialGain := int(simulation.geodes) + int(simulation.geodeRobots)*(32-i)
			if potentialGain > minimalPotentialGain {
				minimalPotentialGain = potentialGain
			}
		}

		// remove simulations that has no potential to get more geodes
		for key, simulation := range deduplicated {
			potentialGain := int(simulation.geodes)
			for j := 0; j < 32-i; j++ {
				potentialGain += int(simulation.geodeRobots) + j
			}

			if potentialGain < minimalPotentialGain {
				delete(deduplicated, key)
			}
		}

		fmt.Println("sims", len(deduplicated))
		simulations = deduplicated
	}

	// find max
	simsSlice := make([]*Simulation, 0, len(simulations))
	for _, simulation := range simulations {
		simsSlice = append(simsSlice, simulation)
	}
	sort.Slice(simsSlice, func(i, j int) bool {
		return simsSlice[i].geodes > simsSlice[j].geodes
	})

	return int(simsSlice[0].geodes)
}
