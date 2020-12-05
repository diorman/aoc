package day05

import "testing"

func TestGetTicketID(t *testing.T) {
	tests := []struct {
		ticket string
		id     int
	}{
		{"BFFFBBFRRR", 567},
		{"FFFBBBFRRR", 119},
		{"BBFFBBFRLL", 820},
	}

	for _, tt := range tests {
		id := getTicketID(tt.ticket)
		if id != tt.id {
			t.Errorf("got %v, want %v", id, tt.id)
		}
	}
}
