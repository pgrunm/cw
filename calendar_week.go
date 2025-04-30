package main

import (
	"encoding/json"
	"fmt"
	"strconv"
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

func parseRequestedDate(submittedDate string) (time.Time, error) {
	if submittedDate == "" {
		return time.Now(), nil
	}

	// Parse the submitted calendar week from string to int
	week, err := strconv.Atoi(submittedDate)
	if err != nil {
		return time.Time{}, err
	}

	return time.Date(time.Now().Year(), 1, 1, 0, 0, 0, 0, time.Local).AddDate(0, 0, (week-1)*7), nil
}

func getWeekOutput(params calendarParams, currentTime time.Time) (formattedWeek string, err error) {
	year, week := getCalendarWeek(currentTime)
	monday := getLastMonday(currentTime)
	weekData := map[string]interface{}{
		"year":  year,
		"week":  week,
		"start": monday.Format(time.DateOnly),
		"end":   monday.AddDate(0, 0, 7).Format(time.DateOnly),
	}

	// If not summary is required, remove the data from the map.
	if !params.summary {
		delete(weekData, "year")
		delete(weekData, "start")
		delete(weekData, "end")
	}

	switch params.output {
	case "json":
		jsonData, jsonErr := json.Marshal(weekData)
		if jsonErr != nil {
			return "", fmt.Errorf("failed to marshal JSON: %v", err)
		}
		formattedWeek = fmt.Sprintf("%s\n", string(jsonData))
	default:
		if params.summary {
			formattedWeek = fmt.Sprintf("It's currently calendar week %d in %d, which started on %s and will finish on %s.\n",
				week, year, monday.Format(time.DateOnly), monday.AddDate(0, 0, 7).Format(time.DateOnly))
		} else {
			formattedWeek = fmt.Sprintf("%d\n", week)
		}
	}
	return formattedWeek, nil
}
