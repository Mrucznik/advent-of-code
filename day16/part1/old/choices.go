package old

type Choice interface {
	Do(s *Simulation)
}

type Move struct {
	where string
}

func (c *Move) Do(s *Simulation) {
	s.currentValve = c.where
}

type Open struct {
	what string
}

func (c *Open) Do(s *Simulation) {
	valve := s.valves[c.what]
	valve.open = true
	s.valves[c.what] = valve
}
