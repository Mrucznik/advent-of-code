package main

import "math/big"

func bigint(i int) *big.Int {
	return big.NewInt(int64(i))
}

func bg(i []int) []*big.Int {
	var bgs []*big.Int
	for _, in := range i {
		bgs = append(bgs, bigint(in))
	}
	return bgs
}

func mainInput() []*Monkey {
	monkeys := []*Monkey{{
		items: bg([]int{89, 73, 66, 57, 64, 80}),
		operation: func(old *big.Int) *big.Int {
			a, b := old, big.NewInt(3)
			return old.Mul(a, b)
		},
		testDivisableBy: 13,
		ifTrueMonkey:    6,
		ifFalseMonkey:   2,
	},
		{
			items: bg([]int{83, 78, 81, 55, 81, 59, 69}),
			operation: func(old *big.Int) *big.Int {
				a, b := old, big.NewInt(1)
				return old.Add(a, b)
			},
			testDivisableBy: 3,
			ifTrueMonkey:    7,
			ifFalseMonkey:   4,
		},
		{
			items: bg([]int{76, 91, 58, 85}),
			operation: func(old *big.Int) *big.Int {
				a, b := old, big.NewInt(13)
				return old.Mul(a, b)
			},
			testDivisableBy: 7,
			ifTrueMonkey:    1,
			ifFalseMonkey:   4,
		},
		{
			items: bg([]int{71, 72, 74, 76, 68}),
			operation: func(old *big.Int) *big.Int {
				a, b := old, old
				return old.Mul(a, b)
			},
			testDivisableBy: 2,
			ifTrueMonkey:    6,
			ifFalseMonkey:   0,
		},
		{
			items: bg([]int{98, 85, 84}),
			operation: func(old *big.Int) *big.Int {
				a, b := old, big.NewInt(7)
				return old.Add(a, b)
			},
			testDivisableBy: 19,
			ifTrueMonkey:    5,
			ifFalseMonkey:   7,
		},
		{
			items: bg([]int{78}),
			operation: func(old *big.Int) *big.Int {
				a, b := old, big.NewInt(8)
				return old.Add(a, b)
			},
			testDivisableBy: 5,
			ifTrueMonkey:    3,
			ifFalseMonkey:   0,
		},
		{
			items: bg([]int{86, 70, 60, 88, 88, 78, 74, 83}),
			operation: func(old *big.Int) *big.Int {
				a, b := old, big.NewInt(4)
				return old.Add(a, b)
			},
			testDivisableBy: 11,
			ifTrueMonkey:    1,
			ifFalseMonkey:   2,
		},
		{
			items: bg([]int{81, 58}),
			operation: func(old *big.Int) *big.Int {
				a, b := old, big.NewInt(5)
				return old.Add(a, b)
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
			items: bg([]int{79, 98}),
			operation: func(old *big.Int) *big.Int {
				a, b := old, big.NewInt(19)
				return old.Mul(a, b)
			},
			testDivisableBy: 23,
			ifTrueMonkey:    2,
			ifFalseMonkey:   3,
		},
		{
			items: bg([]int{54, 65, 75, 74}),
			operation: func(old *big.Int) *big.Int {
				a, b := old, big.NewInt(6)
				return old.Add(a, b)
			},
			testDivisableBy: 19,
			ifTrueMonkey:    2,
			ifFalseMonkey:   0,
		},
		{
			items: bg([]int{79, 60, 97}),
			operation: func(old *big.Int) *big.Int {
				a, b := old, old
				return old.Mul(a, b)
			},
			testDivisableBy: 13,
			ifTrueMonkey:    1,
			ifFalseMonkey:   3,
		},
		{
			items: bg([]int{74}),
			operation: func(old *big.Int) *big.Int {
				a, b := old, big.NewInt(3)
				return old.Add(a, b)
			},
			testDivisableBy: 17,
			ifTrueMonkey:    0,
			ifFalseMonkey:   1,
		},
	}
}
