package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func parseInputFile(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
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

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("could not scan file: %w", err)
	}

	return m, nil
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

func calculateAllSlopesMagicNumber(m [][]int) int {
	return countTrees(m, 1, 1) *
		countTrees(m, 3, 1) *
		countTrees(m, 5, 1) *
		countTrees(m, 7, 1) *
		countTrees(m, 1, 2)
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

	m, err := parseInputFile(filename)
	if err != nil {
		return err
	}

	resultP1 := countTrees(m, 3, 1)
	resultP2 := calculateAllSlopesMagicNumber(m)
	fmt.Fprintf(stdout, "part 1: %d\npart 2: %d\n", resultP1, resultP2)

	return nil
}

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
