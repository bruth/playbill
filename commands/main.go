package commands

import (
	"playbill/cast"
	"playbill/cli"
	"playbill/info"
	"playbill/list"
	"playbill/log"
	"playbill/perform"
	"playbill/rehearse"
)

type CommandMap map[string]*cli.Command

// Create a map of commands by name and aliases
func newCommandMap(cmds ...*cli.Command) CommandMap {
	cm := make(CommandMap, len(cmds))
	for _, cmd := range cmds {
		cm[cmd.Name] = cmd
		// Associate aliases
		for _, alias := range cmd.Aliases {
			cm[alias] = cmd
		}
	}
	return cm
}

// List of commands
var List = []*cli.Command{
	cast.Command,
	info.Command,
	list.Command,
	log.Command,
	rehearse.Command,
	perform.Command,
}

// Map (by name and aliases) of commands
var Map = newCommandMap(List...)
