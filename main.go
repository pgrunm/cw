package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	// Define the flags for the CLI application
	params := calendarParams{}

	app := &cli.App{
		Name:  "cw",
		Usage: "Find the appropriate calendar week of a given date.",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "summary",
				Aliases:     []string{"s"},
				Value:       false,
				Usage:       "Print out the calendar week in a short summary.",
				Destination: &params.summary,
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
				Destination: &params.output,
			},
		},
		Action: func(cCtx *cli.Context) error {
			// Parse any entered calendar week as flag from commandline
			submittedDate := cCtx.Args().First()
			currentTime, err := parseRequestedDate(submittedDate)
			if err != nil {
				return fmt.Errorf("failed to parse the requested date: %v", err)
			}
			// Print the current calendar week
			formattedWeek, err := getWeekOutput(params, currentTime)
			if err != nil {
				return fmt.Errorf("failed to get week output: %v", err)
			}
			print(formattedWeek)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
