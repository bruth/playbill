/*

   Entry point for Playbill CLI

*/

package main

import (
	"fmt"
	"os"
)

func main() {

	// Special case, fallback to help
	if len(os.Args) < 2 {
		HelpCmd.Usage()
	}

	name := os.Args[1]

	// The help command is not in the commands.Map, so this must be
	// check manually
	if name == HelpCmd.Name {
		HelpCmd.Run(HelpCmd, os.Args[2:])
	}

	args := os.Args[2:]

	cmd, ok := CmdMap[name]

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
