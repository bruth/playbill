package cli

var castUsage = `
usage: playbill cast

Casts a new production. If an existing production has the same
name a warning will be displayed.
`

var CastCmd = &Cmd{
	Name:        "cast",
	Short:       "Cast a new production",
	Aliases:     []string{"init", "new"},
	UsageString: castUsage,
}
