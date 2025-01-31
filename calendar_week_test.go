package main

import (
	"testing"
	"time"
)

func TestWeekFunction(t *testing.T) {
	testCases := []struct {
		desc       string
		t          time.Time
		year, week int
	}{
		{
			desc: "Check week for 24.12.2024",
			t:    time.Date(2024, 12, 24, 0, 0, 0, 0, time.UTC),
			week: 52,
			year: 2024,
		},
		{
			desc: "Check week of 1.7.2025",
			t:    time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC),
			week: 26,
			year: 2025,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			year, week := getCalendarWeek(tC.t)
			if year != tC.year && week != tC.week {
				t.Errorf("Expected week %d for %d, got %d", tC.week, tC.year, week)
			}
		})
	}
}
