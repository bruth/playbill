/*

   Entry point for Playbill CLI

*/

package main

import (
	"fmt"
	"os"
	"playbill/commands"
	"playbill/help"
)

func main() {

	// Special case, fallback to help
	if len(os.Args) < 2 {
		help.Command.Usage()
	}

	name := os.Args[1]

	// The help command is not in the commands.Map, so this must be
	// check manually
	if name == help.Command.Name {
		help.Command.Run(help.Command, os.Args[2:])
	}

	args := os.Args[2:]

	cmd, ok := commands.Map[name]

	if !ok {
		fmt.Println("unknown command:", name)
		// Unknown command
		os.Exit(127)
	}

	// Run function not defined, fallback to usage
	if cmd.Run == nil {
		cmd.Usage()
	}

	cmd.Run(cmd, args)
}
