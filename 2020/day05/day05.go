package day05

import (
	"bufio"
	"math"
	"os"
	"sort"
)

var (
	totalRows    = 128
	totalColumns = 8
)

func parsePuzzle(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	ticketIDs := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		id := getTicketID(scanner.Text())
		ticketIDs = append(ticketIDs, id)
	}

	return ticketIDs, scanner.Err()
}

func calculatePosition(total int, s string) int {
	n := 0
	for i, r := range s {
		// 'B' and 'R' indicate n should point to the second half which is n + total/2^i
		// e.g. n + total/2, n + total/4, n + total/8
		if r == 'B' || r == 'R' {
			n += total / int(math.Pow(2, float64(i+1)))
		}
	}
	return n
}

func getTicketID(ticket string) int {
	row := calculatePosition(totalRows, ticket[0:7])
	column := calculatePosition(totalColumns, ticket[7:])
	return row*totalColumns + column
}

func Resolve(puzzlePath string) ([]interface{}, error) {
	ticketIDs, err := parsePuzzle(puzzlePath)
	if err != nil {
		return nil, err
	}

	sort.Ints(ticketIDs)

	mySeat := -1
	for i, id := range ticketIDs {
		nextTicketID := id + 1
		if ticketIDs[i+1] != nextTicketID {
			mySeat = nextTicketID
			break
		}
	}

	return []interface{}{
		ticketIDs[len(ticketIDs)-1],
		mySeat,
	}, nil
}
