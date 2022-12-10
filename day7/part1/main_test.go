package main

import (
	"fmt"
	"testing"
)

func Test_getSumOfDirsBelowThreshold(t *testing.T) {
	ff := Filesystem{}
	ff.RecreateFilesystem([]string{
		"$ cd /",
		"$ ls",
		"dir a",
		"14848514 b.txt",
		"8504156 c.dat",
		"dir d",
		"$ cd a",
		"$ ls",
		"dir e",
		"29116 f",
		"2557 g",
		"62596 h.lst",
		"$ cd e",
		"$ ls",
		"584 i",
		"$ cd ..",
		"$ cd ..",
		"$ cd d",
		"$ ls",
		"4060174 j",
		"8033020 d.log",
		"5626152 d.ext",
		"7214296 k",
	})

	type args struct {
		fs        *Filesystem
		threshold int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				fs:        &ff,
				threshold: 100_000,
			},
			want: 95437,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSumOfDirsBelowThreshold(tt.args.fs, tt.args.threshold); got != tt.want {
				t.Errorf("getSumOfDirsBelowThreshold() = %v, want %v", got, tt.want)
				fmt.Printf("%+v\n", tt.args.fs)
			}
		})
	}
}
