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
}

func getSumOfDirsBelowThreshold(fs *Filesystem, threshold int) int {
	dirsSize := 0
	for _, dir := range fs.GetDirsBelowSize(threshold) {
		dirsSize += dir.GetSize()
	}
	return dirsSize
}
