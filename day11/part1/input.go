package main

func mainInput() []*Monkey {
	monkeys := []*Monkey{{
		items: []int{89, 73, 66, 57, 64, 80},
		operation: func(old int) int {
			return old * 3
		},
		testDivisableBy: 13,
		ifTrueMonkey:    6,
		ifFalseMonkey:   2,
	},
		{
			items: []int{83, 78, 81, 55, 81, 59, 69},
			operation: func(old int) int {
				return old + 1
			},
			testDivisableBy: 3,
			ifTrueMonkey:    7,
			ifFalseMonkey:   4,
		},
		{
			items: []int{76, 91, 58, 85},
			operation: func(old int) int {
				return old * 13
			},
			testDivisableBy: 7,
			ifTrueMonkey:    1,
			ifFalseMonkey:   4,
		},
		{
			items: []int{71, 72, 74, 76, 68},
			operation: func(old int) int {
				return old * old
			},
			testDivisableBy: 2,
			ifTrueMonkey:    6,
			ifFalseMonkey:   0,
		},
		{
			items: []int{98, 85, 84},
			operation: func(old int) int {
				return old + 7
			},
			testDivisableBy: 19,
			ifTrueMonkey:    5,
			ifFalseMonkey:   7,
		},
		{
			items: []int{78},
			operation: func(old int) int {
				return old + 8
			},
			testDivisableBy: 5,
			ifTrueMonkey:    3,
			ifFalseMonkey:   0,
		},
		{
			items: []int{86, 70, 60, 88, 88, 78, 74, 83},
			operation: func(old int) int {
				return old + 4
			},
			testDivisableBy: 11,
			ifTrueMonkey:    1,
			ifFalseMonkey:   2,
		},
		{
			items: []int{81, 58},
			operation: func(old int) int {
				return old + 5
			},
			testDivisableBy: 17,
			ifTrueMonkey:    3,
			ifFalseMonkey:   5,
		},
	}
	return monkeys
}

func testInput() []*Monkey {
	return []*Monkey{
		{
			items: []int{79, 98},
			operation: func(old int) int {
				return old * 19
			},
			testDivisableBy: 23,
			ifTrueMonkey:    2,
			ifFalseMonkey:   3,
		},
		{
			items: []int{54, 65, 75, 74},
			operation: func(old int) int {
				return old + 6
			},
			testDivisableBy: 19,
			ifTrueMonkey:    2,
			ifFalseMonkey:   0,
		},
		{
			items: []int{79, 60, 97},
			operation: func(old int) int {
				return old * old
			},
			testDivisableBy: 13,
			ifTrueMonkey:    1,
			ifFalseMonkey:   3,
		},
		{
			items: []int{74},
			operation: func(old int) int {
				return old + 3
			},
			testDivisableBy: 17,
			ifTrueMonkey:    0,
			ifFalseMonkey:   1,
		},
	}
}
