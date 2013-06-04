/*

   Package help handles rendering command usage and arguments
   for each command.

*/

package main

import (
	"fmt"
	"io"
	"os"
	"playbill/cli"
	"playbill/cli/cast"
	"playbill/cli/info"
	"playbill/cli/list"
	"playbill/cli/log"
	"playbill/cli/perform"
	"playbill/cli/rehearse"
	"text/template"
)

type commandList []*cli.Command
type commandMap map[string]*cli.Command

// List of commands
var CommandList = commandList{
	cast.Command,
	info.Command,
	list.Command,
	log.Command,
	rehearse.Command,
	perform.Command,
}

// Create a map of commands by name and aliases
func NewCommandMap(cmds ...*cli.Command) commandMap {
	cm := make(commandMap, len(cmds))
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
var CommandMap = NewCommandMap(CommandList...)

var tmpl = `usage: playbill command

Playbill is a management tool for ETL workflows.

The available commands are:{{range .Commands}}
    {{.Name | printf "%-12s"}} {{.Short}}{{end}}

See 'playbill help <command>' for more information about a specific command.
`

type TemplateData struct {
	Commands []*cli.Command
}

var HelpCommand = &cli.Command{
	Name: "help",

	Template: tmpl,

	TemplateHelper: func(c *cli.Command, w io.Writer) {
		t := template.Must(template.New("usage").Parse(c.Template))

		// Exexute the template rendering with supplied data
		t.Execute(w, TemplateData{
			CommandList,
		})
	},

	Run: func(c *cli.Command, args []string) {
		if len(args) == 0 {
			c.Usage()
		}

		cmd, ok := CommandMap[args[0]]

		if !ok {
			fmt.Println("unknown command:", args[0])
			// Unknown command
			os.Exit(127)
		}

		cmd.Usage()
	},
}
