package day05

import "testing"

func TestGetTicketID(t *testing.T) {
	tests := []struct {
		ticket []byte
		id     int
	}{
		{[]byte("BFFFBBFRRR"), 567},
		{[]byte("FFFBBBFRRR"), 119},
		{[]byte("BBFFBBFRLL"), 820},
	}

	for _, tt := range tests {
		id := getTicketID(tt.ticket)
		if id != tt.id {
			t.Errorf("got %v, want %v", id, tt.id)
		}
	}
}
