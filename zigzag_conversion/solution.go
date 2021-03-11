package zigzag_conversion

import (
	"strings"
)

func convert(s string, numRows int) string {
	if len(s) <= 2 {
		return s
	}
	lines := getLines(s, numRows)

	columns := len(lines)/2*(numRows-1) + 1

	elements := make([][]string, numRows)
	for i, _ := range elements {
		elements[i] = make([]string, columns)
	}

	for i, line := range lines {
		if i%2 == 0 {
			fillVertical(elements, line, i/2*(numRows-1))
		} else {
			fillSlanted(elements, line, i/2*(numRows-1), numRows)
		}
	}

	var output string
	for _, column := range elements {
		output += strings.Join(column, "")
	}

	return output
}

func getLines(s string, numRows int) []string {
	var (
		lines  = make([]string, 0)
		length = len(s)
		start  = 0
		i      = 0
	)

	for start < length {
		end := start + numRows
		if end > length {
			end = length
		}
		lines = append(lines, s[start:end])

		i++
		start = i * (numRows - 1)
	}

	return lines
}

func fillVertical(elements [][]string, line string, start int) {
	for i, s := range line {
		elements[i][start] = string(s)
	}
}

func fillSlanted(elements [][]string, line string, start int, numRows int) {
	for i, s := range line {
		elements[numRows-i-1][start+i] = string(s)
	}
}
