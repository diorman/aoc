package day05

import (
	"bytes"
	"math"
	"sort"
)

var (
	totalRows    = 128
	totalColumns = 8
)

func Resolve(input []byte) ([]interface{}, error) {
	ticketIDs := parseInput(input)
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

func calculatePosition(total int, s []byte) int {
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

func getTicketID(ticket []byte) int {
	row := calculatePosition(totalRows, ticket[0:7])
	column := calculatePosition(totalColumns, ticket[7:])
	return row*totalColumns + column
}

func parseInput(input []byte) []int {
	ticketIDs := []int{}
	for _, t := range bytes.Split(input, []byte{'\n'}) {
		id := getTicketID(t)
		ticketIDs = append(ticketIDs, id)
	}
	return ticketIDs
}
