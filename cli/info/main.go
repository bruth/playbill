package info

import (
	"playbill/cli"
)

var tmpl = `
usage: playbill info <play>

Prints info about a particular ETL workflow including:

    - description
    - list of acts and scenes
    - pass/fail runs
    - current progress (if in a run)
`

var Command = &cli.Command{
	Name:     "info",
	Aliases:  []string{"stat"},
	Short:    "Prints info about an ETL workflow",
	Template: tmpl,
}
