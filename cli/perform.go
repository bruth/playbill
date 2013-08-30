package cli

var performUsage = `
usage: playbill perform <play> [<act>] [<scene>]
`

var PerformCmd = &Cmd{
	Name:        "perform",
	Aliases:     []string{"run"},
	Short:       "Performs/executes a run",
	UsageString: performUsage,
}
