package list

import (
	"playbill/cli"
)

var tmpl = `
usage: playbill list

List all plays by name.
`

var Command = &cli.Command{
	Name:     "list",
	Aliases:  []string{"ls"},
	Short:    "List all plays by name",
	Template: tmpl,
}
