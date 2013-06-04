package log

import (
	"playbill/cli"
)

var tmpl = `
usage: playbill log [<play>] [<act>] [<scene>]

View the detailed activity log of runs. A play, act, or scene can
be specified to limit the scope of the log.
`

var Command = &cli.Command{
	Name:     "log",
	Short:    "View the activity log",
	Template: tmpl,
}
