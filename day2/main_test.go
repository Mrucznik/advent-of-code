package main

import "testing"

func Test_calculateStrategyOutcome(t *testing.T) {
	type args struct {
		rows []string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name: "test",
			args: args{
				rows: []string{
					"A Y",
					"B X",
					"C Z",
				},
			},
			want:  12,
			want1: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := calculateStrategyOutcome(tt.args.rows)
			if got != tt.want {
				t.Errorf("calculateStrategyOutcome() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("calculateStrategyOutcome() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
