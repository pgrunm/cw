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
	// table contains a boolean, whether to output the calendar week in table format.
	table bool
}

func getCalendarWeek(t time.Time) (year int, week int) {
	return time.Time.ISOWeek(t)
}

// Returns the date of the last monday for t.
func getLastMonday(t time.Time) time.Time {
	return t.AddDate(0, 0, (-int(t.Weekday()) + 1))
}

// parseRequestedDate parses the submitted date string and returns the corresponding time.Time.
// This can be either a calenderweek e.g. `20`, a year with a calendarweek e.g. `24/2023`
// or an empty string or a date in the format `YYYY-MM-DD`.
func parseRequestedDate(submittedDate string) (time.Time, error) {
	if submittedDate == "" {
		return time.Now(), nil
	}

	// Check if the submitted date is in the format YYYY-MM-DD
	if parsedDate, err := time.Parse("2006-01-02", submittedDate); err == nil {
		return parsedDate, nil
	}

	// Check if the submitted date is in the format week/year (e.g., 24/2023)
	if parts := len(submittedDate); parts == 7 && submittedDate[2] == '/' {
		week, err := strconv.Atoi(submittedDate[:2])
		if err != nil {
			return time.Time{}, err
		}
		year, err := strconv.Atoi(submittedDate[3:])
		if err != nil {
			return time.Time{}, err
		}
		// Calculate the first day of the given week/year
		firstDay := time.Date(year, 1, 1, 0, 0, 0, 0, time.Local)
		isoYear, isoWeek := firstDay.ISOWeek()
		for isoYear < year || isoWeek < week {
			firstDay = firstDay.AddDate(0, 0, 1)
			isoYear, isoWeek = firstDay.ISOWeek()
		}
		return firstDay, nil
	}

	// Parse the submitted calendar week from string to int
	week, err := strconv.Atoi(submittedDate)
	if err != nil {
		return time.Time{}, err
	}

	// Calculate the first day of the given week in the current year
	currentYear := time.Now().Year()
	firstDay := time.Date(currentYear, 1, 1, 0, 0, 0, 0, time.Local)
	isoYear, isoWeek := firstDay.ISOWeek()
	for isoYear < currentYear || isoWeek < week {
		firstDay = firstDay.AddDate(0, 0, 1)
		isoYear, isoWeek = firstDay.ISOWeek()
	}
	return firstDay, nil
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
			_, actualWeek := time.Now().ISOWeek()
			// Check if the current week is the same as the requested week. If not choose the appropriate time for output.
			switch {
			case week == actualWeek:
				formattedWeek = fmt.Sprintf("It's currently calendar week %d in %d, which started on %s and will finish on %s.\n",
					week, year, monday.Format(time.DateOnly), monday.AddDate(0, 0, 7).Format(time.DateOnly))
			case week > actualWeek:
				formattedWeek = fmt.Sprintf("Calendar week %d in %d will start on %s and finish on %s.\n",
					week, year, monday.Format(time.DateOnly), monday.AddDate(0, 0, 7).Format(time.DateOnly))
			case week < actualWeek:
				formattedWeek = fmt.Sprintf("Calendar week %d in %d started on %s and finished on %s.\n",
					week, year, monday.Format(time.DateOnly), monday.AddDate(0, 0, 7).Format(time.DateOnly))
			}
		} else {
			formattedWeek = fmt.Sprintf("%d\n", week)
		}
	}
	return formattedWeek, nil
}
