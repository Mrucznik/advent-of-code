package main

import "testing"

func Test_findMarker(t *testing.T) {

	type args struct {
		msg       string
		markerLen int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				msg:       "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
				markerLen: 4,
			},
			want: 7,
		},
		{
			name: "",
			args: args{
				msg:       "bvwbjplbgvbhsrlpgdmjqwftvncz",
				markerLen: 4,
			},
			want: 5,
		},
		{
			name: "",
			args: args{
				msg:       "nppdvjthqldpwncqszvftbrmjlhg",
				markerLen: 4,
			},
			want: 6,
		},
		{
			name: "",
			args: args{
				msg:       "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
				markerLen: 4,
			},
			want: 10,
		},
		{
			name: "",
			args: args{
				msg:       "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
				markerLen: 4,
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMarker(tt.args.msg, tt.args.markerLen); got != tt.want {
				t.Errorf("findMarker() = %v, want %v", got, tt.want)
			}
		})
	}
}
