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

	// Slice of arguments for the subcommand
	args := os.Args[2]

	// The help command is not in the commands.Map, so this must be
	// check manually
	if name == HelpCmd.Name {
		HelpCmd.Run(HelpCmd, args)
	}

	// Ensure the command exists
	cmd, ok := CmdMap[name]

	if !ok {
		HelpCmd.Run(HelpCmd, args)
		// Exit code that denotes an unknown command
		// http://tldp.org/LDP/abs/html/exitcodes.html
		os.Exit(127)
	}

	// Run function not defined, fallback to usage
	if cmd.Run == nil {
		cmd.Usage()
	}

	// If flags are defined for the command, parse the arguments
	if cmd.Flags != nil {
		if err := cmd.Flags.Parse(args); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	cmd.Run(cmd, args)
}
