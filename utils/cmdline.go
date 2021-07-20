package utils

import (
	"fmt"
)

type CmdlineHandlerFunc func(args []string) error

type CommandDetails struct {
	Handler     CmdlineHandlerFunc
	Description string
}

type CommandMap map[string]CommandDetails

func HandleCmdline(args []string, commands *CommandMap) error {
	// Missing command
	if len(args) < 2 {
		return fmt.Errorf("no command specified")
	}

	// Call the command's registered function
	for command, cmdDetails := range *commands {
		if command == args[1] {
			return cmdDetails.Handler(args[1:])
		}
	}

	// Error - Unrecongised command
	return fmt.Errorf("command %v not recognised", args[1])
}
