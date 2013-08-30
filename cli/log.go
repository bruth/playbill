package cli

var logUsage = `
usage: playbill log [<play>] [<act>] [<scene>]

View the detailed activity log of runs. A play, act, or scene can
be specified to limit the scope of the log.
`

var LogCmd = &Cmd{
	Name:        "log",
	Short:       "View the activity log",
	UsageString: logUsage,
}
