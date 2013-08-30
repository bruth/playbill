package cli

var listUsage = `
usage: playbill list

List all plays by name.
`

var ListCmd = &Cmd{
	Name:        "list",
	Aliases:     []string{"ls"},
	Short:       "List all plays by name",
	UsageString: listUsage,
}
