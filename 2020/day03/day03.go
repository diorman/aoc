package day03

import (
	"bufio"
	"os"
)

func Resolve(puzzlePath string) ([]interface{}, error) {
	m, err := parsePuzzle(puzzlePath)
	if err != nil {
		return nil, err
	}

	return []interface{}{
		countTrees(m, 3, 1),
		countTrees(m, 1, 1) *
			countTrees(m, 3, 1) *
			countTrees(m, 5, 1) *
			countTrees(m, 7, 1) *
			countTrees(m, 1, 2),
	}, nil
}

func parsePuzzle(puzzlePath string) ([][]int, error) {
	file, err := os.Open(puzzlePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	m := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := make([]int, len(scanner.Text()))
		for i, c := range scanner.Text() {
			if c == '#' {
				row[i] = 1
			}
		}

		m = append(m, row)
	}

	return m, scanner.Err()
}

func countTrees(m [][]int, right, down int) int {
	var (
		x     = right
		width = len(m[0])
		total = 0
	)

	for y := down; y < len(m); y += down {
		total += m[y][x]

		x += right
		if x >= width {
			x -= width
		}
	}

	return total
}
