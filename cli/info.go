package cli

var infoUsage = `
usage: playbill info <play>

Prints info about a particular ETL workflow including:

    - description
    - list of acts and scenes
    - pass/fail runs
    - current progress (if in a run)
`

var InfoCmd = &Cmd{
	Name:        "info",
	Short:       "Prints info about an ETL workflow",
	UsageString: infoUsage,
}
