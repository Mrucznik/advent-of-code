package main

import (
	"reflect"
	"testing"
)

func Test_moveLeft(t *testing.T) {
	type args struct {
		rock [4]byte
		y    int
	}
	tests := []struct {
		name string
		args args
		want [4]byte
	}{
		{
			name: "normal move",
			args: args{
				rock: [4]byte{
					0b0_0000000,
					0b0_0000000,
					0b0_1100000 >> 2,
					0b0_1100000 >> 2},
				y: 5,
			},
			want: [4]byte{
				0b0_0000000,
				0b0_0000000,
				0b0_1100000 >> 1,
				0b0_1100000 >> 1},
		},
		{
			name: "no move",
			args: args{
				rock: [4]byte{
					0b0_0000000,
					0b0_0000000,
					0b0_1100000,
					0b0_1100000},
				y: 5,
			},
			want: [4]byte{
				0b0_0000000,
				0b0_0000000,
				0b0_1100000,
				0b0_1100000},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := moveLeft(tt.args.rock, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("moveLeft() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_moveRight(t *testing.T) {
	type args struct {
		rock [4]byte
		y    int
	}
	tests := []struct {
		name string
		args args
		want [4]byte
	}{
		{
			name: "normal move",
			args: args{
				rock: [4]byte{
					0b0_0000000,
					0b0_0000000,
					0b0_1100000 >> 2,
					0b0_1100000 >> 2},
				y: 5,
			},
			want: [4]byte{
				0b0_0000000,
				0b0_0000000,
				0b0_1100000 >> 3,
				0b0_1100000 >> 3},
		},
		{
			name: "no move",
			args: args{
				rock: [4]byte{
					0b0_0000000,
					0b0_0000000,
					0b0_1100000 >> 5,
					0b0_1100000 >> 5},
				y: 5,
			},
			want: [4]byte{
				0b0_0000000,
				0b0_0000000,
				0b0_1100000 >> 5,
				0b0_1100000 >> 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := moveRight(tt.args.rock, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("moveRight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fall(t *testing.T) {
	//space[1] = 0b11111111

	type args struct {
		rock [4]byte
		y    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "normal",
			args: args{
				rock: [4]byte{
					0b0_0000000,
					0b0_0000000,
					0b0_1100000 >> 2,
					0b0_1100000 >> 2},
				y: 5,
			},
			want: false,
		},
		{
			name: "no fall",
			args: args{
				rock: [4]byte{
					0b0_0000000,
					0b0_0000000,
					0b0_1100000 >> 2,
					0b0_1100000 >> 2},
				y: 0,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fall(tt.args.rock, tt.args.y); got != tt.want {
				t.Errorf("fall() = %v, want %v", got, tt.want)
			}
		})
	}
}
