package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
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

func parseInputFile(filename string) ([]entry, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
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

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("could not scan file: %w", err)
	}

	return entries, nil
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

func run(args []string, stdout io.Writer) error {
	var (
		flags = flag.NewFlagSet(args[0], flag.ExitOnError)
		test  = flags.Bool("t", false, "use input.test.txt")
	)

	if err := flags.Parse(args[1:]); err != nil {
		return err
	}

	filename := "./input.txt"
	if *test {
		filename = "./input.test.txt"
	}

	entries, err := parseInputFile(filename)
	if err != nil {
		return err
	}

	resultP1 := getTotal(entries, isValidPasswordPolicyA)
	resultP2 := getTotal(entries, isValidPasswordPolicyB)
	fmt.Fprintf(stdout, "part 1: %d\npart 2: %d\n", resultP1, resultP2)

	return nil
}

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
