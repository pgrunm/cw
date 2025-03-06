package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	// summary contains a boolean, whether to output a short summary.
	var summary bool
	// output contains the desired output format.
	var output string

	app := &cli.App{
		Name:  "cw",
		Usage: "Find the appropriate calendar week of a given date.",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "summary",
				Aliases:     []string{"s"},
				Value:       false,
				Usage:       "Print out the calendar week in a short summary.",
				Destination: &summary,
			},

			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "Prints the requested calendar week as JSON.",
				Action: func(cCtx *cli.Context, s string) error {
					if s != "json" && s != "yaml" {
						return fmt.Errorf("Output format must be 'json', got '%s'.", s)
					}
					return nil
				},
				Destination: &output,
			},
		},
		Action: func(cCtx *cli.Context) error {
			var formattedWeek string
			currentTime := time.Now()
			year, week := getCalendarWeek(currentTime)
			monday := getLastMonday(currentTime)
			weekData := map[string]interface{}{
				"year":  year,
				"week":  week,
				"start": monday.Format(time.DateOnly),
				"end":   monday.AddDate(0, 0, 7).Format(time.DateOnly),
			}

			// If not summary is required, remove the data from the map.
			if !summary {
				delete(weekData, "year")
				delete(weekData, "start")
				delete(weekData, "end")
			}

			switch output {
			case "json":
				jsonData, err := json.Marshal(weekData)
				if err != nil {
					return fmt.Errorf("failed to marshal JSON: %v", err)
				}
				formattedWeek = string(jsonData)
			default:
				if summary {
					formattedWeek = fmt.Sprintf("It's currently calendar week %d in %d, which started on %s and will finish on %s.",
						week, year, monday.Format(time.DateOnly), monday.AddDate(0, 0, 7).Format(time.DateOnly))
				} else {
					formattedWeek = fmt.Sprintf("%d", week)
				}
			}
			// Print the current calendar week
			fmt.Println(formattedWeek)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
