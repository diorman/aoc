package day01

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Resolve(puzzlePath string) ([]interface{}, error) {
	entries, err := parsePuzzle(puzzlePath)
	if err != nil {
		return nil, err
	}

	return []interface{}{
		calculateMagicNumber(entries, []int{}, 2, 0),
		calculateMagicNumber(entries, []int{}, 3, 0),
	}, nil
}

func parsePuzzle(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	entries := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		expense, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("could not parse entry '%s': %w", scanner.Text(), err)
		}
		entries = append(entries, expense)
	}

	return entries, scanner.Err()
}

func sum(entries, indexes []int) int {
	total := 0
	for _, index := range indexes {
		total += entries[index]
	}
	return total
}

func multiply(entries, indexes []int) int {
	total := 1
	for _, index := range indexes {
		total *= entries[index]
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
