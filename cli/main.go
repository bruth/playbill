/*

   Defines a Command struct for defining sub-commands and usage text.

*/

package cli

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
)

// A Command is an implementation of a sub-command
type Command struct {
	// Primary name of the command
	Name string

	// Aliases for the command
	Aliases []string

	// Short is the short description shown in the 'help' command
	Short string

	// Template instance containing the usage/documentation
	Template string

	// Executes the template passing in any data or helpers
	TemplateHelper func(c *Command, w io.Writer)

	// Flag is a set of flags specific to this command.
	Flags flag.FlagSet

	// Run runs the command.
	// The args are the arguments after the command name.
	Run func(c *Command, args []string)
}

func (c *Command) Usage() {
	stdout := bytes.Buffer{}

	if c.TemplateHelper != nil {
		c.TemplateHelper(c, &stdout)
	} else {
		t := template.Must(template.New("usage").Parse(c.Template))
		err := t.Execute(&stdout, nil)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println(strings.TrimSpace(stdout.String()))
	os.Exit(2)
}
