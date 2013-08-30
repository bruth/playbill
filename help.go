/*

   Package help handles rendering command usage and arguments
   for each command.

*/

package main

import (
	"fmt"
	"github.com/bruth/playbill/cli"
	"io"
	"os"
	"text/template"
)

// List of commands
var CmdList = []*cli.Cmd{
	cli.CastCmd,
	cli.InfoCmd,
	cli.ListCmd,
	cli.LogCmd,
	cli.RehearseCmd,
	cli.PerformCmd,
}

// Create a map of commands by name and aliases
func NewCmdMap(cmds ...*cli.Cmd) map[string]*cli.Cmd {
	cm := make(map[string]*cli.Cmd, len(cmds))
	for _, cmd := range cmds {
		cm[cmd.Name] = cmd
		// Associate aliases
		for _, alias := range cmd.Aliases {
			cm[alias] = cmd
		}
	}
	return cm
}

// Map (by name and aliases) of commands
var CmdMap = NewCmdMap(CmdList...)

var helpUsage = `usage: playbill <command> [options]

Playbill is a process management tool for ETL workflows.

The available commands are:{{range .Cmds}}
    {{.Name | printf "%-12s"}} {{.Short}}{{end}}

See 'playbill help <command>' for more information about a specific command.
`

type TemplateData struct {
	Cmds []*cli.Cmd
}

var HelpCmd = &cli.Cmd{
	Name: "help",

	UsageString: helpUsage,

	UsageHelper: func(c *cli.Cmd, w io.Writer) {
		t := template.Must(template.New("usage").Parse(c.UsageString))

		// Exexute the template rendering with supplied data
		t.Execute(w, TemplateData{
			CmdList,
		})
	},

	Run: func(c *cli.Cmd, args []string) {
		if len(args) == 0 {
			c.Usage()
		}

		cmd, ok := CmdMap[args[0]]

		if !ok {
			fmt.Println("unknown command:", args[0])
			// Unknown command
			os.Exit(127)
		}

		cmd.Usage()
	},
}
