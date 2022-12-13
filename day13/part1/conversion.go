package main

import (
	"fmt"
	"strconv"
	"strings"
)

func convert(line string) ([]interface{}, int) {
	var result []interface{}
	var buf []rune
	for i := 0; i < len(line); i++ {
		char := rune(line[i])
		if char == '[' {
			blob, processed := convert(line[i+1:])
			result = append(result, blob)
			i += processed
		} else if char == ']' {
			if buf != nil {
				number, err := strconv.Atoi(string(buf))
				if err != nil {
					panic(err)
				}
				result = append(result, number)
				buf = nil
			}
			return result, i + 1
		} else if char == ',' {
			if buf != nil {
				number, err := strconv.Atoi(string(buf))
				if err != nil {
					panic(err)
				}
				result = append(result, number)
				buf = nil
			}
		} else {
			buf = append(buf, char)
		}
	}
	return result, len(line)
}

func check(lines [][]interface{}) {
	rows := strings.Split(input, "\n")

	j := 0
	for i := 0; i < len(rows); i++ {
		if len(rows[i]) == 0 {
			continue
		}
		line := fmt.Sprintf("%+v", lines[j])
		line = strings.ReplaceAll(line, " ", ",")

		if rows[i] != line {
			panic(fmt.Sprintf("%s %d != %s %d", rows[i], i, line, j))
		}
		j++
	}
}
