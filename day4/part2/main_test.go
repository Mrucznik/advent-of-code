package main

import "testing"

func Test_getContainedPairCounts(t *testing.T) {
	type args struct {
		pairs []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{
				[]string{
					"2-4,6-8",
					"2-3,4-5",
					"5-7,7-9",
					"2-8,3-7",
					"6-6,4-6",
					"2-6,4-8",
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getContainedPairCounts(tt.args.pairs); got != tt.want {
				t.Errorf("getContainedPairCounts() = %v, want %v", got, tt.want)
			}
		})
	}
}
