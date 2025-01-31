package main

import "time"

func getCalendarWeek(t time.Time) (year int, week int) {
	return time.Time.ISOWeek(t)
}
