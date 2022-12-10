package main

import (
	"strconv"
	"strings"
)

type Filesystem struct {
	root *Dir
}

func (fs *Filesystem) RecreateFilesystem(statements []string) {
	fs.root = NewDir("/", nil)

	var currDir *Dir
	currDir = fs.root
	for _, statement := range statements {
		if statement[0] == '$' {
			commandStatement := strings.Split(statement, " ")
			cmd := commandStatement[1]
			var arg string
			if len(commandStatement) > 2 {
				arg = commandStatement[2]
			}

			// command
			switch cmd {
			case "cd":
				if arg == "/" {
					currDir = fs.root
				} else if arg == ".." {
					currDir = currDir.parent
				} else {
					currDir = currDir.AddDir(arg)
				}
			case "ls":
				continue
			}
		} else {
			// add file to dir
			fileStatement := strings.Split(statement, " ")
			if fileStatement[0] == "dir" {
				continue
			}
			size, err := strconv.Atoi(fileStatement[0])
			if err != nil {
				panic(err)
			}
			name := fileStatement[1]
			currDir.AddFile(name, size)
		}
	}
}

func (fs *Filesystem) GetDirsBelowSize(size int) []*Dir {
	dirs := []*Dir{fs.root}
	dirs = append(dirs, fs.root.GetDirs()...)
	var result []*Dir
	for _, dir := range dirs {
		if dir.GetSize() < size {
			result = append(result, dir)
		}
	}
	return result
}

func (fs *Filesystem) GetDirsBiggerEqSize(size int) []*Dir {
	dirs := []*Dir{fs.root}
	dirs = append(dirs, fs.root.GetDirs()...)
	var result []*Dir
	for _, dir := range dirs {
		if dir.GetSize() >= size {
			result = append(result, dir)
		}
	}
	return result
}

func (d *Dir) GetDirs() []*Dir {
	var dirs []*Dir
	dirs = append(dirs, d.dirs...)
	for _, dir := range d.dirs {
		dirs = append(dirs, dir.GetDirs()...)
	}
	return dirs
}
