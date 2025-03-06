package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	// summary contains a boolean, whether to output a short summary.
	var summary bool
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
		},
		Action: func(cCtx *cli.Context) error {
			var formattedWeek string
			currentTime := time.Now()
			year, week := getCalendarWeek(currentTime)
			if summary {
				monday := getLastMonday(currentTime)
				formattedWeek = fmt.Sprintf("It's currently calendar week %d in %d, which started on %s and will finish on %s.",
					week, year, monday.Format(time.DateOnly), monday.AddDate(0, 0, 7).Format(time.DateOnly))
			} else {
				formattedWeek = fmt.Sprintf("%d", week)
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
