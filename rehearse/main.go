package rehearse

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"playbill/cli"
)

var Command = &cli.Command{
	Name: "rehearse",

	Short: "Runs a rehearsal of all or some of the scenes",

	Run: func(c *cli.Command, args []string) {
		var stdout bytes.Buffer

		c.Flags.Parse(args)
		args = c.Flags.Args()

		if len(args) == 0 {
			c.Usage()
		}

		cmd := exec.Command(args[0], args[1:]...)

		// Redirect command's stdout/err to be local stdout
		cmd.Stdout = &stdout
		cmd.Stderr = &stdout

		err := cmd.Run()

		fmt.Print(stdout.String())

		if err != nil {
			os.Exit(1)
		}
	},
}
