package main

import "testing"

func TestFilesystem_RecreateFilesystem(t *testing.T) {
	type fields struct {
		root *Dir
	}
	type args struct {
		statements []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "test",
			fields: fields{
				root: nil,
			},
			args: args{
				statements: []string{
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
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &Filesystem{
				root: tt.fields.root,
			}
			fs.RecreateFilesystem(tt.args.statements)
		})
	}
}
