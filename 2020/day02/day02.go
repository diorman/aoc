package day02

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type entry struct {
	min      int
	max      int
	letter   string
	password string
}

type policyFunc func(e entry) bool

var re = regexp.MustCompile(`(\d+)-(\d+)\s([a-z]):\s([a-z]+)`)

func Resolve(puzzlePath string) ([]interface{}, error) {
	entries, err := parsePuzzle(puzzlePath)
	if err != nil {
		return nil, err
	}

	return []interface{}{
		getTotal(entries, isValidPasswordPolicyA),
		getTotal(entries, isValidPasswordPolicyB),
	}, nil
}

func parsePuzzle(path string) ([]entry, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	entries := []entry{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parsedLine := re.FindStringSubmatch(scanner.Text())
		min, err := strconv.Atoi(parsedLine[1])
		if err != nil {
			return nil, fmt.Errorf("could not parse min in '%s': %w", scanner.Text(), err)
		}

		max, err := strconv.Atoi(parsedLine[2])
		if err != nil {
			return nil, fmt.Errorf("could not parse max in '%s': %w", scanner.Text(), err)
		}

		entries = append(entries, entry{min, max, parsedLine[3], parsedLine[4]})
	}

	return entries, scanner.Err()
}

func isValidPasswordPolicyA(e entry) bool {
	n := strings.Count(e.password, e.letter)
	return n >= e.min && n <= e.max
}

func isValidPasswordPolicyB(e entry) bool {
	l0 := string(e.password[e.min-1])
	l1 := string(e.password[e.max-1])
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
