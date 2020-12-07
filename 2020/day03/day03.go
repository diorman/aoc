package day03

import (
	"bytes"
)

func Resolve(input []byte) ([]interface{}, error) {
	m := parseInput(input)
	return []interface{}{
		countTrees(m, 3, 1),
		countTrees(m, 1, 1) *
			countTrees(m, 3, 1) *
			countTrees(m, 5, 1) *
			countTrees(m, 7, 1) *
			countTrees(m, 1, 2),
	}, nil
}

func parseInput(input []byte) [][]int {
	var (
		m    [][]int
		rows = bytes.Split(input, []byte("\n"))
	)

	for _, row := range rows {
		r := make([]int, len(row))
		for i, c := range row {
			if c == '#' {
				r[i] = 1
			}
		}
		m = append(m, r)
	}

	return m
}

func countTrees(m [][]int, right, down int) int {
	var (
		column = right
		width  = len(m[0])
		total  = 0
	)

	for row := down; row < len(m); row += down {
		total += m[row][column]
		column += right
		if column >= width {
			column -= width
		}
	}

	return total
}
