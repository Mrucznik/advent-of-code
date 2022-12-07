package main

import "testing"

func Test_day3(t *testing.T) {
	type args struct {
		rucksack []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test day 3 part 2",
			args: args{
				rucksack: []string{
					"vJrwpWtwJgWrhcsFMMfFFhFp",
					"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
					"PmmdzqPrVvPwwTWBwg",
					"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
					"ttgJtRGJQctTZtZT",
					"CrZsJsPPZsGzwwsLwLmpwMDw",
				},
			},
			want: 70,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumPrioritiesOfTheSameItemsInRucksackCompartments(tt.args.rucksack); got != tt.want {
				t.Errorf("sumPrioritiesOfTheSameItemsInRucksackCompartments() = %v, want %v", got, tt.want)
			}
		})
	}
}
