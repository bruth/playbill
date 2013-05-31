package cast

import (
	"playbill/cli"
)

var tmpl = `
usage: playbill cast

Casts a new production. If an existing production has the same
name a warning will be displayed.
`

var Command = &cli.Command{
	Name:     "cast",
	Short:    "Cast a new production",
	Aliases:  []string{"init", "new"},
	Template: tmpl,
}
