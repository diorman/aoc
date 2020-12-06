package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/diorman/aoc/2020/day01"
	"github.com/diorman/aoc/2020/day02"
	"github.com/diorman/aoc/2020/day03"
	"github.com/diorman/aoc/2020/day04"
	"github.com/diorman/aoc/2020/day05"
	"github.com/diorman/aoc/2020/day06"
)

type resolverFunc func(string) ([]interface{}, error)

var resolvers = []resolverFunc{
	day01.Resolve,
	day02.Resolve,
	day03.Resolve,
	day04.Resolve,
	day05.Resolve,
	day06.Resolve,
}

func run(args []string, stdout io.Writer) error {
	var (
		flags    = flag.NewFlagSet(args[0], flag.ExitOnError)
		dayFlag  = flags.Int("d", 0, "day to run")
		testFlag = flags.Bool("t", false, "use test puzzle")
	)

	if err := flags.Parse(args[1:]); err != nil {
		return err
	}

	if *dayFlag < 1 || *dayFlag > len(resolvers) {
		return fmt.Errorf("invalid day! ho ho ho")
	}

	puzzlePath := fmt.Sprintf("./puzzles/%02d", *dayFlag)
	if *testFlag {
		puzzlePath = fmt.Sprintf("%s_test", puzzlePath)
	}

	results, err := resolvers[*dayFlag-1](puzzlePath)
	if err != nil {
		return err
	}

	for i, result := range results {
		fmt.Fprintf(stdout, "part %d: %v\n", i+1, result)
	}

	return nil
}

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
