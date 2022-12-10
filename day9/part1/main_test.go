package main

import "testing"

func Test_getDist(t *testing.T) {
	type args struct {
		tx int
		ty int
		hx int
		hy int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				tx: 10,
				ty: -5,
				hx: 15,
				hy: -10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDist(tt.args.tx, tt.args.ty, tt.args.hx, tt.args.hy); got != tt.want {
				t.Errorf("getDist() = %v, want %v", got, tt.want)
			}
		})
	}
}
