package day06

import (
	"bytes"
)

type group struct {
	size   int
	counts map[byte]int
}

func Resolve(input []byte) ([]interface{}, error) {
	var (
		sum1   = 0
		sum2   = 0
		groups = parseInput(input)
	)

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

func parseInput(input []byte) []group {
	groups := []group{}
	for _, groupBytes := range bytes.Split(input, []byte("\n\n")) {
		counts := make(map[byte]int)
		size := 1
		for _, b := range groupBytes {
			if b != '\n' {
				counts[b]++
			} else {
				size++
			}
		}
		groups = append(groups, group{size, counts})
	}
	return groups
}
