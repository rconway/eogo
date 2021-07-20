package utils

import (
	"fmt"
	"io"
)

type CmdlineHandlerFunc func(args []string, w io.Writer) (CommandExecutor, error)

type CommandDetails struct {
	Handler     CmdlineHandlerFunc
	Description string
}

type CommandMap map[string]CommandDetails

type UsageFunction func(args []string, commands *CommandMap, w io.Writer)

type CommandExecutor interface {
	Execute() error
}

func HandleCmdline(args []string, commands *CommandMap, usageFn UsageFunction, w io.Writer) (CommandExecutor, error) {
	// Missing command
	if len(args) < 2 {
		usageFn(args, commands, w)
		return nil, fmt.Errorf("no command specified")
	}

	// Call the command's registered function
	for command, cmdDetails := range *commands {
		if command == args[1] {
			return cmdDetails.Handler(args[1:], w)
		}
	}

	// Error - Unrecongised command
	usageFn(args, commands, w)
	return nil, fmt.Errorf("command %v not recognised", args[1])
}
