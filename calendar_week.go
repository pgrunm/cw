package main

import (
	"time"
)

type calendarParams struct {
	// summary contains a boolean, whether to output a short summary.
	summary bool
	// output contains the desired output format.
	output string
}

func getCalendarWeek(t time.Time) (year int, week int) {
	return time.Time.ISOWeek(t)
}

// Returns the date of the last monday for t.
func getLastMonday(t time.Time) time.Time {
	return t.AddDate(0, 0, (-int(t.Weekday()) + 1))
}
