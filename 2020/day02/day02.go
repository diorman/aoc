package day02

import (
	"bytes"
	"fmt"
)

type entry struct {
	min      int
	max      int
	letter   byte
	password []byte
}

type policyFunc func(e entry) bool

func Resolve(input []byte) ([]interface{}, error) {
	entries, err := parseInput(input)
	if err != nil {
		return nil, err
	}

	return []interface{}{
		getTotal(entries, isValidPasswordPolicyA),
		getTotal(entries, isValidPasswordPolicyB),
	}, nil
}

func parseInput(input []byte) ([]entry, error) {
	var (
		lines   = bytes.Split(input, []byte("\n"))
		entries = []entry{}
	)

	for _, line := range lines {
		e := entry{}
		_, err := fmt.Sscanf(string(line), "%d-%d %c: %s", &e.min, &e.max, &e.letter, &e.password)
		if err != nil {
			return nil, err
		}
		entries = append(entries, e)
	}
	return entries, nil
}

func isValidPasswordPolicyA(e entry) bool {
	n := bytes.Count(e.password, []byte{e.letter})
	return n >= e.min && n <= e.max
}

func isValidPasswordPolicyB(e entry) bool {
	l0 := e.password[e.min-1]
	l1 := e.password[e.max-1]
	return (l0 == e.letter) != (l1 == e.letter)
}

func getTotal(entries []entry, p policyFunc) int {
	total := 0
	for _, e := range entries {
		if p(e) {
			total++
		}
	}
	return total
}
