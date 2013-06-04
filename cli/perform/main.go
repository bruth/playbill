package perform

import (
	"playbill/cli"
)

var tmpl = `
usage: playbill perform <play> [<act>] [<scene>]
`

var Command = &cli.Command{
	Name:     "perform",
	Aliases:  []string{"run"},
	Short:    "Performs/executes a run",
	Template: tmpl,
}
