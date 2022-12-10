package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fs := Filesystem{}
	fs.RecreateFilesystem(strings.Split(input, "\n"))
	dirsSize := getSumOfDirsBelowThreshold(&fs, 100_000)
	fmt.Println(dirsSize)

	fmt.Println("part2: ", part2(&fs))
}

func getSumOfDirsBelowThreshold(fs *Filesystem, threshold int) int {
	dirsSize := 0
	for _, dir := range fs.GetDirsBelowSize(threshold) {
		dirsSize += dir.GetSize()
	}
	return dirsSize
}

func part2(fs *Filesystem) int {
	threshold := (fs.root.GetSize() + 30000000) - 70_000_000
	fmt.Printf("treshold: %d\n", threshold)

	dirs := fs.GetDirsBiggerEqSize(threshold)
	minSize := 70_000_000
	for _, dir := range dirs {
		ds := dir.GetSize()
		if ds < minSize {
			fmt.Printf("dir %s suze %d\n", dir.name, ds)
			minSize = ds
		}
	}
	return minSize
}
