package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "cw",
		Usage: "Find the appropriate calendar week of a given date.",
		Action: func(*cli.Context) error {

			currentTime := time.Now()
			year, week := getCalendarWeek(currentTime)

			// Print the current calendar week
			formattedWeek := fmt.Sprintf("It's currently calendar week %d in %d.", week, year)
			fmt.Println(formattedWeek)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
