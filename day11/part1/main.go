package main

import (
	_ "embed"
	"fmt"
	"math"
	"sort"
)

// part 1 time needed: 32min

//go:embed input.txt
var input string

type Monkeys []*Monkey

var monkeys Monkeys

func (m Monkeys) Len() int {
	return len(m)
}

func (m Monkeys) Less(i, j int) bool {
	return m[i].inspections > m[j].inspections
}

func (m Monkeys) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

type Monkey struct {
	items           []int
	operation       func(old int) int
	testDivisableBy int
	ifTrueMonkey    int
	ifFalseMonkey   int
	inspections     int
}

func isDivisable(a int, b int) bool {
	return a%b == 0
}

func (m *Monkey) round() {
	for i := range m.items {
		worryLvl := m.inspectItem(i)

		if isDivisable(worryLvl, m.testDivisableBy) {
			m.throwTo(i, m.ifTrueMonkey)
		} else {
			m.throwTo(i, m.ifFalseMonkey)
		}
	}

	m.items = []int{}
}

func (m *Monkey) throwTo(itemNr, monkeyNr int) {
	monkeys[monkeyNr].items = append(monkeys[monkeyNr].items, m.items[itemNr])
}

func (m *Monkey) inspectItem(nr int) int {
	newWorryLevel := m.operation(m.items[nr])
	newWorryLevel = int(math.Floor(float64(newWorryLevel) / 3))
	m.items[nr] = newWorryLevel
	m.inspections++
	return newWorryLevel
}

func main() {
	monkeys = mainInput()

	for i := 0; i < 20; i++ {
		fmt.Println(i)
		for _, monkey := range monkeys {
			monkey.round()
		}
	}

	sort.Sort(monkeys)

	fmt.Printf("m1 %d m2 %d\n", monkeys[0].inspections, monkeys[1].inspections)
	fmt.Println(monkeys[0].inspections * monkeys[1].inspections)

}
