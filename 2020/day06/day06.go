package day06

import (
	"bufio"
	"bytes"
	"os"
)

type group struct {
	size   int
	counts map[rune]int
}

func Resolve(puzzlePath string) ([]interface{}, error) {
	groups, err := parsePuzzle(puzzlePath)
	if err != nil {
		return nil, err
	}

	sum1 := 0
	sum2 := 0
	for _, g := range groups {
		sum1 += len(g.counts)
		for _, ac := range g.counts {
			if g.size == ac {
				sum2++
			}
		}
	}

	return []interface{}{
		sum1,
		sum2,
	}, nil
}

func newGroup(size int, buffer string) group {
	counts := make(map[rune]int)
	for _, r := range buffer {
		counts[r]++
	}
	return group{size, counts}
}

func parsePuzzle(parsePuzzle string) ([]group, error) {
	file, err := os.Open(parsePuzzle)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var (
		scanner   = bufio.NewScanner(file)
		groups    = []group{}
		groupSize int
		buffer    bytes.Buffer
	)

	for scanner.Scan() {
		if scanner.Text() != "" {
			groupSize++
			buffer.WriteString(scanner.Text())
			continue
		}

		groups = append(groups, newGroup(groupSize, buffer.String()))
		groupSize = 0
		buffer.Reset()
	}

	if buffer.Len() > 0 {
		groups = append(groups, newGroup(groupSize, buffer.String()))
	}

	return groups, scanner.Err()
}
