package main

import (
	"reflect"
	"testing"
)

func Test_doRearrangment(t *testing.T) {
	type args struct {
		rows []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				[]string{
					"    [D]    ",
					"[N] [C]    ",
					"[Z] [M] [P]",
					" 1   2   3 ",
					"",
					"move 1 from 2 to 1",
					"move 3 from 1 to 3",
					"move 2 from 2 to 1",
					"move 1 from 1 to 2",
				},
			},
			want: "CMZ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doRearrangement(tt.args.rows); got != tt.want {
				t.Errorf("doRearrangement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createStacks(t *testing.T) {
	type args struct {
		rows []string
	}
	tests := []struct {
		name string
		args args
		want [3][]rune
	}{

		{
			name: "",
			args: args{
				[]string{
					"    [D]    ",
					"[N] [C]    ",
					"[Z] [M] [P]",
					" 1   2   3 ",
					"",
					"move 1 from 2 to 1",
					"move 3 from 1 to 3",
					"move 2 from 2 to 1",
					"move 1 from 1 to 2",
				},
			},
			want: [3][]rune{
				{'Z', 'N'},
				{'M', 'C', 'D'},
				{'P'},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createStacks(tt.args.rows); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createStacks() = %v, want %v", got, tt.want)
			}
		})
	}
}
