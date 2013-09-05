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

// A Cmd is an implementation of a sub-command
type Cmd struct {
	// Primary name of the command
	Name string

	// Aliases for the command
	Aliases []string

	// Short is the short description shown in the 'help' command
	Short string

	// Template string containing the usage/documentation
	UsageString string

	// Executes the usage string as a template, passing in any data or helpers
	UsageHelper func(c *Cmd, w io.Writer)

	// Flag is a set of flags specific to this command.
	Flags *flag.FlagSet

	// Run runs the command.
	// The args are the arguments after the command name.
	Run func(c *Cmd, args []string)
}

func (c *Cmd) Usage() {
	stdout := bytes.Buffer{}

	if c.UsageHelper != nil {
		c.UsageHelper(c, &stdout)
	} else {
		t := template.Must(template.New("usage").Parse(c.UsageString))
		err := t.Execute(&stdout, nil)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println(strings.TrimSpace(stdout.String()))
	os.Exit(2)
}
