# cw

[![Go Report Card](https://goreportcard.com/badge/github.com/pgrunm/cw)](https://goreportcard.com/report/github.com/pgrunm/cw)
[![OpenSSF Scorecard](https://api.scorecard.dev/projects/github.com/pgrunm/cw/badge)](https://scorecard.dev/viewer/?uri=github.com/pgrunm/cw)

Aren't you also annoyed by people, who say `XYZ will take place in calendar week XX`? cw comes to solve this issue! It is a commandline tool for finding out the calendar week of a date. It's completely written in Go.

<!-- Howto write a good readme, see
https://github.com/create-go-app/cli#readme
https://github.com/ergochat/ergo/blob/master/README.md
Nice and clean: https://github.com/urfave/cli
https://github.com/patrickhener/goshs
https://github.com/mr-karan/doggo
-->

## Features

- [x] Get the current calendar week without specifying any arguments.
- [x] Get the calendar week of a date.
- [x] When entering a week as flag, get start and end date.
- [x] Print a list of all calendar weeks of a given year as a table
- [x] Print the start and end date if you enter a year and calendar week
- [x] Support for different output formats like JSON, ~~table~~ (maybe more to come)
- [ ] Support for different date formats

## Installation

### Install as Go binary

Simply run `go install github.com/pgrunm/cw@latest` to install the program directly.

### Download binary package

You can also download a compiled package from the [releases page](https://github.com/pgrunm/cw/releases).

## Usage

cw supports a variety of different commandline parameters. You can see them by running `cw -h` or below:

```txt
NAME:
   cw - Find the appropriate calendar week of a given date.

USAGE:
   cw [global options] command [command options]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --summary, -s             Print out the calendar week in a short summary. (default: false)
   --table, -t               Print out the calendar weeks in a table format. (default: false)
   --output value, -o value  Prints the requested calendar week as JSON.
   --help, -h                show help
```

## Development

Development takes place in main branch, so it can be ahead of other branches. In case you'd like to access a specific version, please use the according tag or download any desired release.

1. Clone the repository with `git clone https://github.com/pgrunm/cw`
2. Run `go get .`.
3. Start developing.
<!-- 
- [Golangci-lint golden config](https://gist.github.com/maratori/47a4d00457a92aa426dbd48a18776322)
- [Project layout](https://github.com/golang-standards/project-layout)
- [OpenSSF Score Card](https://github.com/marketplace/actions/ossf-scorecard-action)
-->
