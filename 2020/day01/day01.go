package day01

import (
	"bytes"
	"strconv"
)

func Resolve(input []byte) ([]interface{}, error) {
	entries, err := parseInput(input)
	if err != nil {
		return nil, err
	}

	return []interface{}{
		calculateMagicNumber(entries, []int{}, 2, 0),
		calculateMagicNumber(entries, []int{}, 3, 0),
	}, nil
}

func parseInput(input []byte) ([]int, error) {
	var (
		lines   = bytes.Split(input, []byte("\n"))
		entries = []int{}
	)

	for _, line := range lines {
		expense, err := strconv.Atoi(string(line))
		if err != nil {
			return nil, err
		}
		entries = append(entries, expense)
	}

	return entries, nil
}

func multiply(entries, indexes []int) int {
	total := 1
	for _, index := range indexes {
		total *= entries[index]
	}
	return total
}

func sum(entries, indexes []int) int {
	total := 0
	for _, index := range indexes {
		total += entries[index]
	}
	return total
}

func calculateMagicNumber(entries, indexes []int, n, startIndex int) int {
	for i := startIndex; i < len(entries); i++ {
		idxs := append(indexes, i)

		if len(idxs) < n {
			res := calculateMagicNumber(entries, idxs, n, startIndex+1)
			if res != -1 {
				return res
			}
			continue
		}

		if sum(entries, idxs) == 2020 {
			return multiply(entries, idxs)
		}
	}

	return -1
}
