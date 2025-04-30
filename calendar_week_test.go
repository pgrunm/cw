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

func TestLastMonday(t *testing.T) {
	testCases := []struct {
		desc        string
		t, expected time.Time
	}{
		{
			desc:     "Get last Monay for current date",
			t:        time.Date(2025, 1, 31, 12, 30, 0, 0, time.UTC),
			expected: time.Date(2025, 1, 27, 12, 30, 0, 0, time.UTC),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := getLastMonday(tC.t)
			if actual.Day() != tC.expected.Day() {
				t.Errorf("Days are not equal, expected %d, got %d", tC.expected.Day(), actual.Day())
			}
		})
	}
}

func TestWeekOutput(t *testing.T) {
	testCases := []struct {
		desc     string
		params   calendarParams
		expected string
		time     time.Time
	}{
		{
			desc: "Check week output w summary",
			params: calendarParams{
				summary: true,
			},
			time:     time.Date(2025, 1, 27, 8, 30, 0, 0, time.UTC),
			expected: "It's currently calendar week 5 in 2025, which started on 2025-01-27 and will finish on 2025-02-03.\n",
		},
		{
			desc: "Check week output wo summary",
			params: calendarParams{
				summary: false,
				output:  "text",
			},
			time:     time.Date(2025, 1, 27, 8, 30, 0, 0, time.UTC),
			expected: "5\n",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual, _ := getWeekOutput(tC.params, tC.time)
			if actual != tC.expected {
				t.Errorf("Expected %s, got %s", tC.expected, actual)
			}
		})
	}
}
func TestParseRequestedDate(t *testing.T) {
	testCases := []struct {
		desc          string
		input         string
		expectedYear  int
		expectedMonth time.Month
		expectedDay   int
		expectError   bool
	}{
		{
			desc:          "Parse empty string (current date)",
			input:         "",
			expectedYear:  time.Now().Year(),
			expectedMonth: time.Now().Month(),
			expectedDay:   time.Now().Day(),
			expectError:   false,
		},
		{
			desc:          "Parse calendar week only (week 25 of current year)",
			input:         "25",
			expectedYear:  2025,
			expectedMonth: time.June,
			expectedDay:   16, // Assuming week 25 starts on June 17th
			expectError:   false,
		},
		{
			desc:          "Parse year with calendar week (week 25 of 2024)",
			input:         "25/2024",
			expectedYear:  2024,
			expectedMonth: time.June,
			expectedDay:   17, // Assuming week 25 starts on June 17th
			expectError:   false,
		},
		{
			desc:          "Parse specific date (2025-01-01)",
			input:         "2025-01-01",
			expectedYear:  2025,
			expectedMonth: time.January,
			expectedDay:   1,
			expectError:   false,
		},
		{
			desc:        "Parse invalid input",
			input:       "invalid",
			expectError: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual, err := parseRequestedDate(tC.input)
			if tC.expectError {
				if err == nil {
					t.Errorf("Expected an error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Did not expect an error but got: %v", err)
				}
				if actual.Year() != tC.expectedYear || actual.Month() != tC.expectedMonth || actual.Day() != tC.expectedDay {
					t.Errorf("Expected date %d-%02d-%02d, got %d-%02d-%02d",
						tC.expectedYear, tC.expectedMonth, tC.expectedDay,
						actual.Year(), actual.Month(), actual.Day())
				}
			}
		})
	}
}
