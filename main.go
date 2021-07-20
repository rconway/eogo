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

func processCmdline(args []string, w io.Writer) {
	commands := utils.CommandMap{
		"records": utils.CommandDetails{Handler: records.RecordsCmdlineHandler, Description: "tooling for OGC API Records"},
		"process": utils.CommandDetails{Handler: process.ProcessCmdlineHandler, Description: "tooling for OGC API Processes"},
	}
	err := utils.HandleCmdline(args, &commands, usage, w)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
	}
}

func main() {
	processCmdline(os.Args, os.Stdout)
}
