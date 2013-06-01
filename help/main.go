/*

   Package help handles rendering command usage and arguments
   for each command.

*/

package help

import (
	"fmt"
	"io"
	"os"
	"playbill/cli"
	"playbill/commands"
	"text/template"
)

var tmpl = `usage: playbill command

Playbill is a management tool for ETL workflows.

The available commands are:{{range .Commands}}
    {{.Name | printf "%-12s"}} {{.Short}}{{end}}

See 'playbill help <command>' for more information about a specific command.
`

type TemplateData struct {
	Commands []*cli.Command
}

var Command = cli.Command{
	Name: "help",

	Template: tmpl,

	TemplateHelper: func(c *cli.Command, w io.Writer) {
		t := template.Must(template.New("usage").Parse(c.Template))

		// Exexute the template rendering with supplied data
		t.Execute(w, TemplateData{
			commands.List,
		})
	},

	Run: func(c *cli.Command, args []string) {
		if len(args) == 0 {
			c.Usage()
		}

		cmd, ok := commands.Map[args[0]]

		if !ok {
			fmt.Println("unknown command:", args[0])
			// Unknown command
			os.Exit(127)
		}

		cmd.Usage()
	},
}
