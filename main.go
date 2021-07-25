package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/rconway/eogo/process"
	"github.com/rconway/eogo/records"
	"github.com/rconway/eogo/utils"
)

func usage(args []string, commands *utils.CommandMap, w io.Writer) {
	fmt.Fprintf(w, "Usage:\n%v <command> <options>\n", filepath.Base(args[0]))
	fmt.Fprintln(w, "  <command> is one of...")
	for k, v := range *commands {
		fmt.Fprintf(w, "    %v: %v\n", k, v.Description)
	}
	fmt.Fprintln(w, "  <options> varies for each command")
}

// Receives the full program command-line
func processCmdline(args []string, w io.Writer) (utils.CommandExecutor, error) {
	commands := utils.CommandMap{
		"records": utils.CommandDetails{Handler: records.RecordsCmdlineHandler, Description: "tooling for OGC API Records"},
		"process": utils.CommandDetails{Handler: process.ProcessCmdlineHandler, Description: "tooling for OGC API Processes"},
	}
	if len(args) < 2 {
		usage(args, &commands, w)
		return nil, fmt.Errorf("no command specified")
	}
	return utils.HandleCmdline(args, &commands, usage, w)
}

func main() {
	outWriter := os.Stdout

	cmdExe, err := processCmdline(os.Args, outWriter)
	if err == nil {
		if cmdExe != nil {
			err = cmdExe.Execute()
		} else {
			err = fmt.Errorf("the command executor is nil")
		}
	}

	if err != nil {
		fmt.Fprintf(outWriter, "ERROR: %v\n", err)
	}
}
