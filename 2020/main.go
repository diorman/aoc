package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/diorman/aoc/2020/day01"
	"github.com/diorman/aoc/2020/day02"
	"github.com/diorman/aoc/2020/day03"
	"github.com/diorman/aoc/2020/day04"
	"github.com/diorman/aoc/2020/day05"
	"github.com/diorman/aoc/2020/day06"
	"github.com/diorman/aoc/2020/day07"
)

type resolverFunc func([]byte) ([]interface{}, error)

var resolvers = []resolverFunc{
	day01.Resolve,
	day02.Resolve,
	day03.Resolve,
	day04.Resolve,
	day05.Resolve,
	day06.Resolve,
	day07.Resolve,
}

func run(args []string, stdout io.Writer) error {
	var (
		flags        = flag.NewFlagSet(args[0], flag.ExitOnError)
		dayFlag      = flags.Int("d", 0, "day to run")
		testFlag     = flags.Bool("t", false, "use test puzzle")
		inputPathFmt string
	)

	if err := flags.Parse(args[1:]); err != nil {
		return err
	}

	if *dayFlag < 1 || *dayFlag > len(resolvers) {
		return fmt.Errorf("invalid day! ho ho ho")
	}

	if *testFlag {
		inputPathFmt = "./day%02d/day%02d_test.in"
	} else {
		inputPathFmt = "./day%02d/day%02d.in"
	}

	inputPath := fmt.Sprintf(inputPathFmt, *dayFlag, *dayFlag)
	input, err := ioutil.ReadFile(inputPath)
	if err != nil {
		return err
	}

	// remove trailing line breaks added by editor on save
	input = bytes.TrimSuffix(input, []byte{'\n'})

	results, err := resolvers[*dayFlag-1](input)
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
