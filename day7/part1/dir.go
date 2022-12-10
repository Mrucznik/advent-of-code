package main

import "fmt"

type Dir struct {
	name   string
	parent *Dir
	files  []*File
	dirs   []*Dir
}

func NewDir(name string, parent *Dir) *Dir {
	return &Dir{name: name, parent: parent, files: []*File{}, dirs: []*Dir{}}
}

func (d *Dir) GetSize() int {
	size := 0
	for _, dir := range d.dirs {
		size += dir.GetSize()
	}
	for _, file := range d.files {
		size += file.GetSize()
	}
	fmt.Printf("dir %s size: %d\n", d.name, size)
	return size
}

func (d *Dir) AddFile(name string, size int) {
	d.files = append(d.files, NewFile(name, size))
	fmt.Printf("adding file %s size %d\n", name, size)
}

func (d *Dir) AddDir(name string) *Dir {
	newDir := NewDir(name, d)
	d.dirs = append(d.dirs, newDir)
	return newDir
}

func (d *Dir) ChangeDir(name string) *Dir {
	for _, dir := range d.dirs {
		if dir.name == "name" {
			return dir
		}
	}
	return nil
}
