package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
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
	if strings.Contains(submittedDate, "/") {
		parts := strings.Split(submittedDate, "/")
		if len(parts) != 2 {
			return time.Time{}, fmt.Errorf("invalid week/year format")
		}
		week, err := strconv.Atoi(parts[0])
		if err != nil {
			return time.Time{}, err
		}
		year, err := strconv.Atoi(parts[1])
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

func getAllWeeksOfYear(currentTime time.Time) (any, any) {
	year := currentTime.Year()
	firstDay := time.Date(year, 1, 1, 0, 0, 0, 0, currentTime.Location())
	for firstDay.Weekday() != time.Monday {
		firstDay = firstDay.AddDate(0, 0, 1)
	}

	var weeks [][]time.Time
	week := []time.Time{}
	day := firstDay

	for day.Year() == year {
		week = append(week, day)
		if len(week) == 7 {
			weeks = append(weeks, week)
			week = []time.Time{}
		}
		day = day.AddDate(0, 0, 1)
	}
	if len(week) > 0 {
		weeks = append(weeks, week)
	}

	return weeks, nil
}

// printWeeksTable prints the weeks in a table format.
func printWeeksTable(weeks any, currentTime time.Time) {
	weeksSlice, ok := weeks.([][]time.Time)
	if !ok {
		fmt.Println("Invalid weeks data")
		return
	}
	_, currentISOWeek := currentTime.ISOWeek()
	fmt.Printf("Calendar Weeks for %d:\n", currentTime.Year())
	fmt.Println("+---------+-------+------------+------------+")
	fmt.Println("| Current | Week  |   Start    |    End     |")
	fmt.Println("+---------+-------+------------+------------+")
	for i, week := range weeksSlice {
		weekNum := i + 1
		start := week[0].Format("2006-01-02")
		end := week[len(week)-1].Format("2006-01-02")
		star := " "
		if weekNum == currentISOWeek {
			star = "*"
		}
		fmt.Printf("|    %s    | %2d    | %s | %s |\n", star, weekNum, start, end)
	}
	fmt.Println("+---------+-------+------------+------------+")
}
