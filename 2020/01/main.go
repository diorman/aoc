package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

func parseInputFile(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
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

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("could not scan file: %w", err)
	}

	return entries, nil
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

	resultP1 := calculateMagicNumber(entries, []int{}, 2, 0)
	resultP2 := calculateMagicNumber(entries, []int{}, 3, 0)
	fmt.Fprintf(stdout, "part 1: %d\npart 2: %d\n", resultP1, resultP2)

	return nil
}

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
