package main

import (
	"testing"
	"time"

	"github.com/nitin-dixit/snippetBox/internal/assert"
)

func TestHumanDate(t *testing.T) {
	// Define zones for testing
	ist := time.FixedZone("IST", 19800) // UTC + 5:30
	cet := time.FixedZone("CET", 3600)  // UTC + 1:00

	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "IST",
			tm:   time.Date(2024, 3, 17, 15, 45, 0, 0, ist),
			want: "17 Mar 2024 at 15:45",
		},
		{
			name: "Empty", // Fixed typo "Emtpy"
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET to IST",
			// 11:15 CET + 4.5 hours = 15:45 IST
			tm:   time.Date(2024, 3, 17, 11, 15, 0, 0, cet),
			want: "17 Mar 2024 at 15:45",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hd := humanDate(tt.tm)
			assert.Equal(t, hd, tt.want)
		})
	}
}
