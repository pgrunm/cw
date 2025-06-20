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
- [ ] Print a list of all calendar weeks of a given year as a table
- [x] Print the start and end date if you enter a year and calendar week
- [x] Support for different output formats like JSON, ~~table~~ (maybe more to come)
- [ ] Support for different date formats

## Installation

TODO: Prepare installation manual.

## Usage

TODO: Prepara usage documentation.

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
